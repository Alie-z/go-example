package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

func main() {
	GlobalClient := redis.NewClient(
		&redis.Options{
			Addr:         "127.0.0.1:6379",
			DialTimeout:  10 * time.Second,
			ReadTimeout:  30 * time.Second,
			WriteTimeout: 30 * time.Second,
			Password:     "111",
			PoolSize:     10,
			DB:           0,
		},
	)
	err := GlobalClient.Ping().Err()
	if nil != err {
		panic(err)
	}
	//redis乐观锁支持，可以通过watch监听一些Key, 如果这些key的值没有被其他人改变的话，才可以提交事务。
	// 定义一个回调函数，用于处理事务逻辑
	fn := func(tx *redis.Tx) error {
		// 先查询下当前watch监听的key的值
		v, err := tx.Get("pipe_test").Result()
		if err != nil && err != redis.Nil {
			return err
		}
		// 这里可以处理业务
		fmt.Println(v)
		// 如果key的值没有改变的话，Pipelined函数才会调用成功
		_, err = tx.Pipelined(func(pipe redis.Pipeliner) error {
			// 在这里给key设置最新值
			pipe.Set("pipe_test", "new value 1111111111111", 0)
			return nil
		})
		return err
	}
	// 使用Watch监听一些Key, 同时绑定一个回调函数fn, 监听Key后的逻辑写在fn这个回调函数里面
	// 如果想监听多个key，可以这么写：client.Watch(fn, "key1", "key2", "key3")
	err = GlobalClient.Watch(fn, "pipe_test")
	// ctx := context.Background()
	// getValue(ctx, GlobalClient)
	// hSet(GlobalClient)
	lSet(GlobalClient)
	if nil != err {
		panic(err)
	}
}
func getValue(ctx context.Context, client *redis.Client) {
	get := client.Get("name1")
	fmt.Println(">>>", get.Val(), get.Err())

	val, err := client.Get("name1").Result()
	switch {
	case err == redis.Nil:
		fmt.Println("key不存在")
	case err != nil:
		fmt.Println("错误", err)
	case val == "":
		fmt.Println("值是空字符串")
	}
}

func hSet(client *redis.Client) {
	// 保存到redis
	err := client.HMSet("user:1001", map[string]interface{}{
		"field1": "value1",
		"field2": "value2",
		"field3": "value3",
	})
	fmt.Println(">>> err", err)
	// if nil != err {
	// 	panic(err)
	// }
	// // 获取设置的字段值
	values, e := client.HMGet("user:1001", "field1", "field2", "field3").Result()
	fmt.Println(">>> values", values, e)
	if e == nil {
		fmt.Printf("field1: %s, field2: %s, field3: %s\n", values[0], values[1], values[2])
	}

}

func lSet(client *redis.Client) {
	// 保存到redis
	err := client.LPush("user:1002", "value1", "value2", "value3")
	fmt.Println(">>> err", err)

	values, e := client.LRange("user:1002", 0, -1).Result()
	fmt.Println(">>> values", values, e)
	if e == nil {
		for _, v := range values {
			fmt.Printf("%s\n", v)
		}
	}
}
