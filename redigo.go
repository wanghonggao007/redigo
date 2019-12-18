package main

import (
	"fmt"

	"github.com/wanghonggao007/redigo/redis"
)

func main() {
	c, err := redis.Dial("tcp", "192.168.131.134:16379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	fmt.Println(c, err)
	if _, err := c.Do("AUTH", "123"); err != nil {
		fmt.Println(err)
		c.Close()
		return
	}
	defer c.Close()
	_, err = c.Do("SET", "xing", "Wang")
	if err != nil {
		fmt.Println("redis set failed:", err)
	}

	username, err := redis.String(c.Do("GET", "xing"))
	if err != nil {
		fmt.Println("redis get failed:", err)
	} else {
		fmt.Printf("Get xing: %v \n", username)
	}
}
