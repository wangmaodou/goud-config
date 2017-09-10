package zookeeper

import (
	"github.com/samuel/go-zookeeper/zk"
	"time"
	"fmt"
)

const (
	_                        = iota
	EventNodeCreated
	EventNodeDeleted
	EventNodeDataChanged
	EventNodeChildrenChanged
)

var (
	methods = map[string]CallBack{}
)

type zkClient struct {
	conn *zk.Conn
}

type CallBack interface {
	OnDataChange(path string)
	OnChildNodeChange(path string)
}

func Connect(hosts []string) (*zkClient, error) {
	client := zkClient{}
	option:=zk.WithEventCallback(listen)
	conn, _, err := zk.Connect(hosts, time.Second*5,option)
	if err != nil {
		return nil, err
	}
	client.conn = conn
	return &client, nil
}

func (z *zkClient) Close() {
	z.conn.Close()
}

func (z *zkClient) CreateNode(path string, data string) error {
	_, err := z.conn.Create(path, []byte(data), zk.FlagEphemeral, zk.WorldACL(zk.PermAll))
	return err
}

func (z *zkClient) Set(path string, data string) error {
	_, err := z.conn.Set(path, []byte(data), zk.FlagEphemeral)
	return err
}

func (z *zkClient) GetNodeData(path string) (string, error) {
	date, _, err := z.conn.Get(path)
	return string(date), err
}

func (z *zkClient) GetChildNodes(path string) ([]string, error) {
	childNodes, _, err := z.conn.Children(path)
	return childNodes, err
}

func (z *zkClient) AddWatch(path string, fun CallBack) {
	methods[path] = fun
}

func listen(event zk.Event) {
	fmt.Println(event)
	switch {
	case event.Type == zk.EventNodeDataChanged:
		fun, ok := methods[event.Path]
		if ok {
			fun.OnDataChange(event.Path)
		}
	case event.Type==zk.EventNodeChildrenChanged:
		fun, ok :=methods[event.Path]
		if ok{
			fun.OnChildNodeChange(event.Path)
		}
	default:


	}
}


