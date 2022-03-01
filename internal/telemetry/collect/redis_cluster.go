package collect

import (
	"errors"
	"fmt"
	"sync/atomic"
	"time"

	redis "github.com/go-redis/redis/v7"
	"github.com/marmotedu/iam/pkg/log"
	uuid "github.com/satori/go.uuid"
)

var (
	switchReids atomic.Value
	dataRedis   atomic.Value
	redisUp     atomic.Value
)

type RedisCluster struct {
	Addr     string
	IsSwitch bool
}

func (r *RedisCluster) singleton() redis.UniversalClient {
	return singleton(r.IsSwitch)
}
func (r *RedisCluster) up() bool {
	if v := redisUp.Load(); v != nil {
		return v.(bool)
	}

	return false
}

func NewRedisClusterPool(addr string) redis.UniversalClient {
	log.Debug("Creating new Redis connection pool")
	var client redis.UniversalClient
	opts := &redis.Options{
		Addr: addr,
	}
	client = redis.NewClient(opts)
	return client
}

func singleton(IsSwitch bool) redis.UniversalClient {
	if IsSwitch {
		v := switchReids.Load()
		if v != nil {
			return v.(redis.UniversalClient)
		}

		return nil
	}
	if v := dataRedis.Load(); v != nil {
		return v.(redis.UniversalClient)
	}

	return nil
}
func connectSingleton(isSwitch bool, addr string) bool {
	if singleton(isSwitch) == nil {
		log.Debug("Connecting to redis cluster")
		if isSwitch {
			switchReids.Store(NewRedisClusterPool(addr))

			return true
		}
		dataRedis.Store(NewRedisClusterPool(addr))

		return true
	}

	return true
}
func clusterConnectionIsOpen(cluster RedisCluster) bool {
	c := singleton(cluster.IsSwitch)
	testKey := "redis-test-" + uuid.Must(uuid.NewV4()).String()
	if err := c.Set(testKey, "test", time.Second).Err(); err != nil {
		log.Warnf("Error trying to set test key: %s", err.Error())

		return false
	}
	if _, err := c.Get(testKey).Result(); err != nil {
		log.Warnf("Error trying to get test key: %s", err.Error())

		return false
	}

	return true
}

func ConnectToRedis() {
	tick := time.NewTicker(time.Second)
	defer tick.Stop()
	c := []RedisCluster{
		{Addr: ":6379", IsSwitch: true}, {Addr: ":6380"},
	}
	var ok bool
	for _, v := range c {
		if !connectSingleton(v.IsSwitch, v.Addr) {
			break
		}
		if !clusterConnectionIsOpen(v) {
			break
		}
		ok = true
	}
	redisUp.Store(ok)
again:
	for {
		select {
		case <-tick.C:
			for _, v := range c {
				if !connectSingleton(v.IsSwitch, v.Addr) { //只是防止空指针
					redisUp.Store(false)
					fmt.Println("connectSingleton")
					goto again
				}
				if !clusterConnectionIsOpen(v) { //可以重连
					fmt.Println("clusterConnectionIsOpen")
					redisUp.Store(false)

					break
				}
			}
		}
	}
}

var ErrRedisIsDown = errors.New("storage: Redis is either down or ws not configured")

func (r *RedisCluster) StartSubHandler(channel string, callback func(interface{})) error {
	//	fmt.Println("start receive msg")
	if !r.up() {
		return ErrRedisIsDown
	}
	client := r.singleton()
	if client == nil {
		return errors.New("redis connection failed")
	}

	//订阅频道
	pubsub := client.Subscribe(channel)
	defer pubsub.Close()

	if _, err := pubsub.Receive(); err != nil {
		log.Errorf("Error while receiving pubsub message: %s", err.Error())

		return err
	}
	log.Info("start receive msg")
	for msg := range pubsub.Channel() {
		callback(msg)
	}

	return nil
}
