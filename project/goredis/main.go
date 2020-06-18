package main


import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"os"
	"time"
)

var redisConn *redis.Pool

func GetRedisConn() *redis.Pool {
	return redisConn
}

func InitRedisPoll() {
	redisConn = &redis.Pool{
		MaxIdle:     100,
		MaxActive:   10,
		IdleTimeout: 3000,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", "127.0.0.1:6379")
			if err != nil {
				return nil, err
			}
			// 验证密码
			// if conf.RedisConf.Password != "" {
			// 	if _, err := c.Do("AUTH", conf.RedisConf.Password); err != nil {
			// 		_ = c.Close()
			// 		return nil, err
			// 	}
			// }
			// 选择数据库
			if _, err := c.Do("SELECT", 3); err != nil {
				_ = c.Close()
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}
	return
}

//
func HSet(key string, field string, data string) error {
	conn := GetRedisConn().Get()
	defer conn.Close()
	
	_, err := conn.Do("HSET", key, field, data)
	if err != nil {
		return err
	}
	
	return nil
}

func HGet(key string, field string) (string, error) {
	conn := GetRedisConn().Get()
	defer conn.Close()
	
	reply, err := redis.String(conn.Do("HGET", key, field))
	if err != nil && err != redis.ErrNil {
		return "", err
	}
	if err == redis.ErrNil {
		return "", nil
	}
	return reply, nil
}

func HDel(key,field string) error {
	conn := GetRedisConn().Get()
	defer conn.Close()
	_,err := conn.Do("HDEL",key,field)
	return err
}




func main() {
	InitRedisPoll()
	taskName := "app检查任务"
	taskID := "1"
	entryID := "2"
	err := HSet(taskName,taskID,entryID)
	if err != nil{
		fmt.Println("添加reids失败",err.Error())
		os.Exit(1)
	}
	
	
	
	rdate,_ := HGet(taskName,taskID)
	fmt.Println("获取的数据是：",rdate)
	
	
	// er := HDel(taskName,taskID)
	// if er != nil{
	// 	fmt.Println("key 删除失败")
	// }
	// rdate2,_ := HGet(taskName,taskID)
	// if len(rdate2) <=0 {
	// 	fmt.Println("任务不存在")
	// }
	
	


}




