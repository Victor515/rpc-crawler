package main

import (
	"crawler/engine"
	"crawler/scheduler"
	"crawler/zhenai/parser"
	"rpc-crawler/persist/client"
	"fmt"
	"rpc-crawler/config"
	crawler_config "crawler/config"
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
		Parser: engine.NewFuncParser(parser.ParseCityList, crawler_config.ParseCityList),
	})

	//e.Run(engine.Request{
	//	Url: "http://www.zhenai.com/zhenghun/aba",
	//	ParserFunc:parser.ParseCity,
	//})

}