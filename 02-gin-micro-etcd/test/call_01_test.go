package test

import (
	"github.com/micro/go-micro/v2/client/selector"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

func callAPI(addr, path, method string) (string, error) {
	req, _ := http.NewRequest(method, "http://"+addr+path, nil)
	client := http.DefaultClient
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	buf, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(buf), nil
}

// 获取node信息，手动调用
func Test_call_01(t *testing.T) {
	etcdReg := etcd.NewRegistry(
		registry.Addrs("127.0.0.1:2379"))

	getService, err := etcdReg.GetService("ProdSrv")
	if err != nil {
		log.Fatalf("get service failed, err:%v\n", err)
		return
	}

	next := selector.RoundRobin(getService)

	node, err := next()
	if err != nil {
		log.Fatalln(err)
		return
	}

	res, err := callAPI(node.Address, "/v1/prods", "GET")
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Println(res) // [{"ProdID":100,"ProdName":"Prod100"},{"ProdID":101,"ProdName":"Prod101"}]

}