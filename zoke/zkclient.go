package zoke

import (
	"errors"
	"fmt"
	"github.com/go-zookeeper/zk"
	"log"
	"sync"
	"time"
)

const (
	_ = iota
	EventNodeCreated
	EventNodeDeleted
	EventNodeDataChanged
	EventNodeChildrenChanged
)

var (
	methods = map[string]CallBack{}
	client  *ZkClient
)

type ZkClient struct {
	conn *zk.Conn
}

type CallBack interface {
	OnDataChange(path string, client *ZkClient)
	OnChildNodeChange(path string, client *ZkClient)
	//OnNodeCreate(path string)
	//OnNodeDelete(path string)
}

func newClient(hosts []string) {
	cli := ZkClient{}
	option := zk.WithEventCallback(listen)
	conn, _, err := zk.Connect(hosts, time.Second*5, option)
	checkError(err)
	cli.conn = conn
	client = &cli
}

func GetInstanceClient(hosts []string) (*ZkClient, error) {
	once := sync.Once{}
	once.Do(func() {
		newClient(hosts)
	})
	if client == nil {
		return nil, errors.New("Failed to connect zookeeper,please check your hosts.")
	} else {
		return client, nil
	}
}

/**
 * Init the config on the remote zookeeper.
 */
func (z *ZkClient) InitServiceConfig(path string, config map[string]string) {

}

func (z *ZkClient) Close() {
	z.conn.Close()
}

func (z *ZkClient) CreateNode(path string, data string) error {
	_, err := z.conn.Create(path, []byte(data), zk.FlagEphemeral, zk.WorldACL(zk.PermAll))
	return err
}

func (z *ZkClient) SetNodeData(path string, data string) error {
	_, err := z.conn.Set(path, []byte(data), zk.FlagEphemeral)
	return err
}

func (z *ZkClient) GetNodeData(path string) (string, error) {
	date, _, err := z.conn.Get(path)
	return string(date), err
}

func (z *ZkClient) GetChildNodes(path string) ([]string, error) {
	childNodes, _, err := z.conn.Children(path)
	return childNodes, err
}

func (z *ZkClient) AddChildrenWatch(path string, call CallBack) {
	methods[path] = call
}

func (z *ZkClient) AddDataWatch(path string, call CallBack) {
	methods[path] = call
}

func listen(event zk.Event) {
	fmt.Println(event)
	switch {
	case event.Type == zk.EventNodeDataChanged:
		fun, ok := methods[event.Path]
		if ok {
			fun.OnDataChange(event.Path, client)
			go func() {
				client.conn.GetW(event.Path)
			}()
		}
	case event.Type == zk.EventNodeChildrenChanged:
		fun, ok := methods[event.Path]
		if ok {
			fun.OnChildNodeChange(event.Path, client)
			go func() {
				client.conn.GetW(event.Path)
			}()
		} /*
			case event.Type==zoke.EventNodeCreated:
				fun, ok :=methods[event.Path]
				if ok{
					fun.OnNodeCreate(event.Path)
					go func(){
						client.conn.ExistsW(event.Path)
					}()
				}
			case event.Type==zoke.EventNodeDeleted:
				fun, ok :=methods[event.Path]
				if ok{
					fun.OnNodeDelete(event.Path)
				}*/

	default:

	}
}

// error check
func checkError(err error) {
	if err != nil {
		log.Println(err)
	}
}
