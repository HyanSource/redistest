package main

import (
	"fmt"
	"strconv"

	"github.com/garyburd/redigo/redis"
)

func main() {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379") //連線

	if err != nil {
		fmt.Println("redis dial err:", err) //顯示訊息
	}

	defer conn.Close()

	//RedisSetHash(conn, rand.Intn(100), "AS")
	RedisSetHash(conn, 1, "AS")

	RedisGetHash(conn, 1)

	fmt.Println("OK！")

}

//redis寫入Hash的方法
func RedisSetHash(con redis.Conn, id int, name string) {

	//會複寫的方法
	//_, err := con.Do("HSET", "account", strconv.Itoa(id), name)
	//不會複寫的方法
	_, err := con.Do("HSETNX", "account", strconv.Itoa(id), name)

	if err != nil {
		fmt.Println("Hset err:", err)
		return
	}

	fmt.Println("success")
}

//redis讀取Hash的方法
func RedisGetHash(con redis.Conn, id int) {
	Read, err := redis.String(con.Do("HGET", "account", strconv.Itoa(id)))

	if err != nil {
		fmt.Println("Hget err:", err)
		return
	}

	fmt.Println("Name:", Read)

	fmt.Println("success")
}
