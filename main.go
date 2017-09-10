package main

import (
	"./zk"
	"time"
	"fmt"
)

type Impl struct{
	zookeeper.CallBack
}

func (c Impl) OnDataChange(path string){
	fmt.Println("onDataChange======>",path)
}

func (c Impl) OnChildNodeChange(path string){
	fmt.Println("onChildNodeChange======>",path)
}

func main() {
	var ch chan string
	ch=make(chan string,10)
	client, err := zookeeper.Connect([]string{"127.0.0.1:2181"})
	if err != nil {
		//fmt.Println(err)
		return
	}
	defer client.Close()
	client.CreateNode("/test/one","900")
	fmt.Println(client.GetNodeData("/test"))
	client.AddWatch("/test", Impl{})
	client.Set("/test/one","00")
	client.Set("/test","00")
	client.CreateNode("/test/node","maodou")
	go func() {
		for {
			ch<-"0"
			time.Sleep(time.Second*1)
		}
	}()
	for i:=0;i<50;i++{
		<-ch
	}

}
