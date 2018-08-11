package main

import (
	"crawler/engine"
	"crawler/scheduler"
	"crawler/zhenai/parser"
	"rpc-crawler/persist/client"
	"fmt"
	"rpc-crawler/config"
)

func main() {
	// use distributed itemsaver
	itemChan, err := client.ItemSaver(fmt.Sprintf(":%d", config.ItemSaverPort))
	if err != nil{
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueueScheduler{},
		WorkerCount: 100,
		ItemChan:    itemChan,
	}
	//e := engine.SimpleEngine{}

	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})

	//e.Run(engine.Request{
	//	Url: "http://www.zhenai.com/zhenghun/aba",
	//	ParserFunc:parser.ParseCity,
	//})

}