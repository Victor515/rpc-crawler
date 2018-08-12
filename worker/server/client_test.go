package main

import (
	"testing"
	"rpc-crawler/rpcsupport"
	"rpc-crawler/worker"
	"time"
	"crawler/config"
	rpcconfig "rpc-crawler/config"
	"fmt"
)

func TestCrawlerService(t *testing.T){
	const host = ":9000"
	go rpcsupport.ServeRpc(host, worker.CrawlerService{})
	time.Sleep(time.Second)

	client, err := rpcsupport.NewClient(host)
	if err != nil{
		panic(err)
	}

	req := worker.Request{
		Url:"http://album.zhenai.com/u/108906739",
		Parser: worker.SerializedParser{
			Name: config.ParseProfile,
			Args: "安静的雪",
		},
	}

	var result worker.ParserResult
	err = client.Call(rpcconfig.CrawlerServiceRpc, req, &result)
	if err != nil{
		t.Error(err)
	}else{
		fmt.Println(result)
	}
}
