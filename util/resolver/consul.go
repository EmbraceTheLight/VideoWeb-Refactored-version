package resolver

// consul 相关封装逻辑
import (
	"context"
	"fmt"
	consulAPI "github.com/hashicorp/consul/api"
	"math/rand"
	"sync"
	"time"
)

// ConsulResolver 封装consul相关操作，单例模式
type ConsulResolver struct {
	client *consulAPI.Client
}

var (
	resolverInstance *ConsulResolver
	once             sync.Once
)

func init() {
	GetConsulResolver()
}

// GetConsulResolver 获取consul resolver实例
func GetConsulResolver() *ConsulResolver {
	once.Do(func() {
		config := consulAPI.DefaultConfig() // 默认是127.0.0.1:8500
		client, err := consulAPI.NewClient(config)
		if err != nil {
			panic(fmt.Sprintf("Failed to create Consul client: %v", err))
		}
		resolverInstance = &ConsulResolver{client: client}
	})
	return resolverInstance
}

// GetServiceAddr 获取指定服务名的地址列表
func GetServiceAddr(ctx context.Context, serviceName string) ([]string, error) {
	cr := GetConsulResolver()
	queryOptions := &consulAPI.QueryOptions{}
	services, _, err := cr.client.Health().Service(serviceName, "", true, queryOptions.WithContext(ctx)) // 只拿健康节点
	if err != nil {
		return nil, err
	}
	var addrs []string
	for _, service := range services {
		ip := service.Service.Address
		if ip == "" {
			ip = service.Node.Address
		}
		addrs = append(addrs, fmt.Sprintf("%s:%d", ip, service.Service.Port))
	}
	return addrs, nil
}

// GetRandomAddr 取得指定服务名的地址列表中的一个随机地址
func GetRandomAddr(ctx context.Context, serviceName string) (string, error) {
	addrs, err := GetServiceAddr(ctx, serviceName)
	if err != nil {
		return "", err
	}
	if len(addrs) == 0 {
		return "", fmt.Errorf("No address found for service %s", serviceName)
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	selected := addrs[r.Intn(len(addrs))]
	return selected, nil
}
