package main

import (
	"fmt"
	"github.com/wangmaodou/goud-config/server"
	"github.com/wangmaodou/goud-config/zoke"
)

type Impl struct {
	zoke.CallBack
}

func (c Impl) OnDataChange(path string, client *zoke.ZkClient) {
	fmt.Println("onDataChange======>", path)
	fmt.Println(client.GetNodeData(path))
}

func (c Impl) OnChildNodeChange(path string, client *zoke.ZkClient) {
	fmt.Println("OnChildNodeChange======>", path)
	fmt.Println(client.GetChildNodes(path))

}

func main() {
	server.Start()
}
