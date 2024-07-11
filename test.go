package main

import (
	"fmt"
	"github.com/go-zookeeper/zk"
	"time"
)

var hosts = []string{"127.0.0.1:2181"}

var path1 = "/maodou"

var flags int32 = zk.FlagEphemeral
var data1 = []byte("hello,this is a zoke go test demo!!!")
var acls = zk.WorldACL(zk.PermAll)
var conn *zk.Conn

func init() {
}

func main3() {
	option := zk.WithEventCallback(callback)
	var err error
	conn, _, err = zk.Connect(hosts, time.Second*5, option)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	time.Sleep(time.Second * 10000000)
	/*_, _, _, err = conn.ExistsW(path1)
	if err != nil {
		fmt.Println(err)
		return
	}

	create(conn, path1, data1)
	_, _, _, err = conn.ChildrenW(path1)
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		create(conn, path1+"/yu", data1)
		time.Sleep(time.Second * 2)
		conn.Delete(path1+"/yu",0)
		time.Sleep(time.Second* 2)
	}

	_, _, _, err = conn.ExistsW(path1)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, _, _, err = conn.ExistsW(path1+"yuuu")
	if err != nil {
		fmt.Println(err)
		return
	}
	//delete(conn, path1)
	conn.Delete(path1,0)*/
}

func callback(event zk.Event) {
	fmt.Println("*******************")
	fmt.Println("path:", event.Path)
	fmt.Println("type:", event.Type.String())
	fmt.Println("state:", event.State.String())
	fmt.Println("-------------------")
	go func() {
		_, _, _, err := conn.ChildrenW(path1)
		if err != nil {
			fmt.Println(err)
			return
		}
	}()
}

func create(conn *zk.Conn, path string, data []byte) {
	_, err_create := conn.Create(path, data, flags, acls)
	if err_create != nil {
		fmt.Println(err_create)
		return
	}
}
