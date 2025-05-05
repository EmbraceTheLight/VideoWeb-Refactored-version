package data

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/registry"
	consulAPI "github.com/hashicorp/consul/api"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/bson"
	"os"
	"time"
	utilCtx "util/context"
	mysql2 "util/dbutil/mysqlutil"
	"vw_comment/internal/data/dal/model"
	"vw_comment/internal/data/dal/query"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"vw_comment/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	// Data
	NewData,
	NewTransaction,
	NewMySQL,
	NewRedisClusterClient,
	NewMongo,

	// Repo
	NewVideoCommentRepo,

	// Client
	NewDiscovery,
	NewRegistrar,
)

type keyType struct{}

// transactionKey is a context key for gorm transactionKey
type transactionKey keyType

// Data .
type Data struct {
	mysql *gorm.DB
	redis *redis.ClusterClient
	mongo *MongoDB
}

// NewTransaction return a util.Transaction interface.
func NewTransaction(d *Data) mysql2.Transaction {
	return d
}

// NewData .
func NewData(
	mysql *gorm.DB,
	redisCluster *redis.ClusterClient,
	mongo *MongoDB,
	logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{
		mysql: mysql,
		redis: redisCluster,
		mongo: mongo,
	}, cleanup, nil
}

func NewMySQL(c *conf.Data) *gorm.DB {
	host, port, username, password, dbname := c.Mysql.Host, c.Mysql.Port, c.Mysql.User, c.Mysql.Password, c.Mysql.Database
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic(err)
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(int(c.Mysql.MaxIdle))
	sqlDB.SetMaxOpenConns(int(c.Mysql.MaxOpen))
	if err := sqlDB.Ping(); err != nil {
		panic(err)
	}
	query.SetDefault(db)
	initOrderMap()
	return db
}

func NewDiscovery(conf *conf.Registry) registry.Discovery {
	c := consulAPI.DefaultConfig()
	c.Address = conf.Consul.Address
	c.Scheme = conf.Consul.Scheme
	cli, err := consulAPI.NewClient(c)
	if err != nil {
		panic(err)
	}
	r := consul.New(cli, consul.WithHealthCheck(false))
	return r
}

func NewRegistrar(conf *conf.Registry) registry.Registrar {
	c := consulAPI.DefaultConfig()
	c.Address = conf.Consul.Address
	c.Scheme = conf.Consul.Scheme
	cli, err := consulAPI.NewClient(c)
	if err != nil {
		panic(err)
	}
	r := consul.New(cli, consul.WithHealthCheck(false))
	return r
}

func NewRedisClusterClient(c *conf.Data) *redis.ClusterClient {
	var address []string
	ipAddress := c.RedisCluster.Host
	for _, port := range c.RedisCluster.Port {
		address = append(address, fmt.Sprintf("%s:%s", ipAddress, port))
	}
	redisCluster := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:        address,
		Password:     c.RedisCluster.Password,
		PoolSize:     int(c.RedisCluster.PoolSize),
		MinIdleConns: int(c.RedisCluster.MinIdleConns),
		MaxRetries:   int(c.RedisCluster.MaxRetries),
		DialTimeout:  c.RedisCluster.DialTimeout.AsDuration(),
		ReadTimeout:  c.RedisCluster.ReadTimeout.AsDuration(),
		WriteTimeout: c.RedisCluster.WriteTimeout.AsDuration(),
		PoolTimeout:  c.RedisCluster.PoolTimeout.AsDuration(),
	})
	err := redisCluster.ForEachShard(context.Background(), func(ctx context.Context, shard *redis.Client) error {
		return shard.Ping(ctx).Err()
	})
	if err != nil {
		panic(err)
	}
	return redisCluster
}

func NewMongo(c *conf.Data) *MongoDB {
	ctx, cancel := context.WithTimeout(utilCtx.NewBaseContext(), time.Duration(c.Mongo.ConnTimeout.Seconds)*time.Second)
	defer cancel()

	// 1. Get and wrap mongo client
	mongoCli, err := mongo.Connect(
		ctx,
		options.Client().
			SetConnectTimeout(c.Mongo.ConnTimeout.AsDuration()).
			SetAppName("vw_video").
			SetMaxPoolSize(uint64(c.Mongo.MaxOpen)).
			SetRetryReads(true).
			SetRetryWrites(true).
			ApplyURI(fmt.Sprintf("mongodb://%s:%s@%s:%s/?replicaSet=rs0", c.Mongo.Username, c.Mongo.Password, c.Mongo.Host, c.Mongo.Port)),
	)
	if err != nil {
		panic(err)
	}
	if err := mongoCli.Ping(ctx, nil); err != nil {
		panic(err)
	}

	mdb := &MongoDB{
		db:          "video_web",
		mongoClient: mongoCli,
	}
	err = setUniqueIndex(ctx, mdb, uc_status, bson.D{{"comment_id", 1}, {"user_id", 1}}, "comment_id__user_id_unique_idx")
	if err != nil {
		panic(err)
	}
	return mdb
}

func setUniqueIndex(ctx context.Context, mdb *MongoDB, collection string, keys bson.D, indexName string) error {
	idxName, err := mdb.Collection(collection).Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys:    keys,
		Options: options.Index().SetName(indexName).SetUnique(true),
	})
	logger := log.NewHelper(log.With(log.NewStdLogger(os.Stdout)))
	if err != nil {
		return err
	}
	logger.Infof("create unique index %s on collection %s success", idxName, collection)
	return nil
}

// startTx sets transactionKey to context and starts a transaction.
func startTx(ctx context.Context) (context.Context, func(err error)) {
	tx := query.Q.Begin()
	ctx = utilCtx.WithValue(ctx, transactionKey{}, tx) // set transactionKey to context
	return ctx, func(err error) { commitTx(ctx, err) }
}

// commitTx commits the transactionKey in context.
// err is the error that out of the transaction.
func commitTx(ctx context.Context, err error) {
	value := utilCtx.MustGetValue(ctx, transactionKey{})

	tx := value.(*query.QueryTx)

	// Handle panic which could occur in tilCtx.MustGetValue.
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic: %v", r) // 确保 err 被赋值，事务正确回滚
			tx.Rollback()
			panic(r) // 继续抛出 panic，防止业务逻辑被吞掉
		}
	}()

	if err != nil || tx.Error != nil {
		tx.Rollback()
		return
	}
	tx.Commit()
}

// WithTx starts and commits(or rollbacks) a transaction Automatically.
// Closure function fn is defined at Biz layer, and it contains the transaction logic you want to execute.
func (d *Data) WithTx(ctx context.Context, fn func(context.Context) error) error {
	var err error
	ctx, commit := startTx(ctx)
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("error: %v", r) // 确保 err 被赋值，事务正确回滚
			commit(err)
			panic(r) // 继续抛出 panic，防止业务逻辑被吞掉
		}
		commit(err)
	}()
	err = fn(ctx)
	return err
}

// getQuery is a helper function.
// It returns common query *query.Query or transactional query *(query.QueryTx).Query.
// With this function, methods of data layer don't need to care about if it's in transactionKey or not.
func getQuery(ctx context.Context) *query.Query {
	// if ctx has transactionKey, return transactional query
	tx, ok := utilCtx.GetValue(ctx, transactionKey{})
	if ok {
		return tx.(*query.QueryTx).Query
	}
	return query.Q
}

// addCommentModel is a helper function.
// It replaces the underlying gorm.DB's model with a video model, so that gorm-gen can use optimistic lock.
// Use the function when need to UPDATE a video model.
// See https://github.com/go-gorm/optimisticlock/issues/36 for more details.
func addCommentModel(ctx context.Context, commentId int64) (query.ICommentDo, *model.Comment, error) {
	comment := getQuery(ctx).Comment
	commentDo := comment.WithContext(ctx)
	commentModel, err := commentDo.Where(comment.CommentID.Eq(commentId)).First()
	if err != nil {
		return commentDo, nil, err
	}
	commentDo.ReplaceDB(commentDo.UnderlyingDB().Model(commentModel))
	return commentDo, commentModel, nil
}
