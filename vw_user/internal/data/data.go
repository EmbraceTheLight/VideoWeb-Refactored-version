package data

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
	utilCtx "util/context"
	"vw_user/internal/conf"
	"vw_user/internal/data/dal/query"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	NewMySQL,
	NewRedisClusterClient,
	NewMongo,
	NewUserRepo,
	NewUserInfoRepo,
	NewCaptRepo,
	NewFileRepo,
)

// Data .
type Data struct {
	mysql *gorm.DB
	redis *redis.ClusterClient
	mongo *mongo.Client
}

// NewData .
func NewData(
	mysql *gorm.DB,
	redisCluster *redis.ClusterClient,
	mongo *mongo.Client,
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
	return db
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
func NewMongo(c *conf.Data) *mongo.Client {
	ctx, cancel := context.WithTimeout(utilCtx.NewBaseContext(), time.Duration(c.Mongo.ConnTimeout.Seconds)*time.Second)
	defer cancel()

	mongoCli, err := mongo.Connect(
		ctx,
		options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s", c.Mongo.Host, c.Mongo.Port)).
			SetMaxPoolSize(uint64(c.Mongo.MaxOpen)),
	)
	if err != nil {
		panic(err)
	}
	if err := mongoCli.Ping(ctx, nil); err != nil {
		panic(err)
	}
	return mongoCli
}
