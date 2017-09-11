package main

import (
	"./zk"
	"time"
	"fmt"
)

type Impl struct{
	zookeeper.CallBack
}

func (c Impl) OnDataChange(path string,client *zookeeper.ZkClient){
	fmt.Println("onDataChange======>",path)
	fmt.Println(client.GetNodeData(path))
}

func (c Impl) OnChildNodeChange(path string,client *zookeeper.ZkClient){
	fmt.Println("OnChildNodeChange======>",path)
	fmt.Println(client.GetChildNodes(path))

}

func main() {
	var ch chan string
	ch=make(chan string,10)
	client, err := zookeeper.Connect([]string{"127.0.0.1:2181"})
	if err != nil {
		fmt.Println(err)
		return
	}
	defer client.Close()
	impl:=&Impl{}
	client.AddChildrenWatch("/maodou",impl)
	go func() {
		for {
			ch<-"0"
			time.Sleep(time.Second*1)
		}
	}()
	for i:=0;i<500;i++{
		<-ch
	}

}
