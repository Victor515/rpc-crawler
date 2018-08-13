package main

import (
	"crawler/engine"
	"crawler/scheduler"
	"crawler/zhenai/parser"
	itemsaver "rpc-crawler/persist/client"
	"fmt"
	"rpc-crawler/config"
	crawler_config "crawler/config"
	worker "rpc-crawler/worker/client"
)

func main() {
	// use distributed itemsaver
	itemChan, err := itemsaver.ItemSaver(fmt.Sprintf(":%d", config.ItemSaverPort))
	if err != nil{
		panic(err)
	}

	processor, err := worker.CreateProcessor()
	if err != nil{
		panic(err)
	}


	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueueScheduler{},
		WorkerCount: 100,
		ItemChan:    itemChan,
		RequestProcessor: processor,
	}
	//e := engine.SimpleEngine{}

	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		Parser: engine.NewFuncParser(parser.ParseCityList, crawler_config.ParseCityList),
	})

	//e.Run(engine.Request{
	//	Url: "http://www.zhenai.com/zhenghun/aba",
	//	ParserFunc:parser.ParseCity,
	//})

}