package zoke

import "fmt"

var (
	Reactor=Impl{}
)

type Impl struct{
	CallBack
}

func (c Impl) OnDataChange(path string,client *ZkClient){
	fmt.Println("onDataChange======>",path)
	fmt.Println(client.GetNodeData(path))
}

func (c Impl) OnChildNodeChange(path string,client *ZkClient){
	fmt.Println("OnChildNodeChange======>",path)
	fmt.Println(client.GetChildNodes(path))

}

