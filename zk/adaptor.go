package zoke

import "fmt"

var (
	Reactor=new(Impl)
)

type Impl struct{
	CallBack
}

func (i Impl) OnDataChange(path string,client *ZkClient){
	fmt.Println("onDataChange======>",path)
	fmt.Println(client.GetNodeData(path))
}

func (i Impl) OnChildNodeChange(path string,client *ZkClient){
	fmt.Println("OnChildNodeChange======>",path)
	fmt.Println(client.GetChildNodes(path))

}

