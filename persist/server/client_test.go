package main

import (
	"testing"
	"rpc-crawler/rpcsupport"
	"crawler/engine"
	"crawler/model"
	"time"
	"rpc-crawler/config"
)

func TestItemSaver(t *testing.T){
	const host = ":1234"
	// start ItemSaver Server
	go ServeRpc("test1", host)

	// give the server time to start
	time.Sleep(time.Second)

	// start client
	client, err := rpcsupport.NewClient(host)
	if err != nil{
		panic(err)
	}

	// save
	expected := engine.Item{
		Url:     "http://album.zhenai.com/u/108906739",
		Type:    "zhenai",
		Id:      "108906739",
		Payload: model.Profile{
			Age: 34,
			Height: 162,
			Weight: 57,
			Income: "3001-5000元",
			Gender: "女",
			Name: "安静的雪",
			Xinzuo: "牡羊座",
			Occupation: "人事/行政",
			Marriage: "离异",
			House: "已购房",
			Hukou: "山东菏泽",
			Education: "大学本科",
			Car: "未购车",
		},
	}

	result := ""
	err = client.Call(config.ItemSaverRpc, expected, &result)

	if err != nil || result != "ok"{
		t.Errorf("Result: %s; Error: %v", result, err)
	}
}
