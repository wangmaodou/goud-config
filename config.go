package main

import (
	"./zk"
	"log"
	"os"
	"bufio"
	"strings"
	"sync"
	"errors"
)

const (
	NODE_NAME = "/goud/config"
)

var (
	conf *Rconfig
)

type Rconfig struct {
}

func GetRemoteConfig(hosts []string, service string, path string) (*Rconfig,error) {
	sync.Once{}.Do(func() {
		newRemoteConfig(hosts,service,path)
	})
	if conf==nil {
		return nil,errors.New("Something wrong happened.")
	}
	return conf,nil
}

//连接zookeeper，并注册/goud/config/service目录
//读取默认配置文件，读到map中，并将其注册到zookeeper的service目录下
//为每个目录添加监听
func newRemoteConfig(hosts []string, service string, path string) {
	client, err := zoke.GetInstanceClient(hosts)
	checkError(err)
	createServiceNode(client, service)
	param := readDefaultConfig(path)
	commitConfig(param)

	conf = &Rconfig{}
}

//get
func (c *Rconfig) GetValue(key string) string {
	return ""
}

//init config node on zookeeper for current service.
func createServiceNode(client *zoke.ZkClient, service string) {
	client.CreateNode(NODE_NAME+"/"+service, "0")
}

//read default configs form path.
func readDefaultConfig(path string) map[string]string {
	file, err := os.OpenFile(path, os.O_RDONLY, 0755)
	checkError(err)
	result := getProperties(file)
	return result
}

//commit the default config to zookeeper.
func commitConfig(param map[string]string) {

}

//get configs from java properties file.
func getProperties(file *os.File) map[string]string {
	reader := bufio.NewReader(file)
	result := make(map[string]string)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		s := string(line)
		s = strings.Replace(s, " ", "", -1)
		if len(s) == 0 || strings.Index(s, "#") == 0 {
			continue
		}
		kv := strings.Split(s, "=")
		if len(kv) != 2 {
			continue
		}
		result[kv[0]] = kv[1]
	}
	return result
}

//error check
func checkError(err error) {
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
