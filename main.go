package main

import (
	"./zk"
	"./web"
	"fmt"
)

type Impl struct{
	zoke.CallBack
}

func (c Impl) OnDataChange(path string,client *zoke.ZkClient){
	fmt.Println("onDataChange======>",path)
	fmt.Println(client.GetNodeData(path))
}

func (c Impl) OnChildNodeChange(path string,client *zoke.ZkClient){
	fmt.Println("OnChildNodeChange======>",path)
	fmt.Println(client.GetChildNodes(path))

}

func main() {
	/*var ch chan string
	ch=make(chan string,10)
	client, err := zk.Connect([]string{"127.0.0.1:2181"})
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
	}*/
	web.Start()

}
