package main

import (
	"crawler/engine"
	"crawler/scheduler"
	"crawler/zhenai/parser"
	itemsaver "rpc-crawler/persist/client"
	"fmt"
	crawler_config "crawler/config"
	worker "rpc-crawler/worker/client"
	"net/rpc"
	"rpc-crawler/rpcsupport"
	"log"
	"flag"
	"strings"
)

var(
	itemSaverHost = flag.String("itemsaver_hosts", "", "itemsaver host")
	workerHosts = flag.String("worker_hosts", "", "worker hosts(comma separated)")
)

func main() {
	flag.Parse()
	// use distributed itemsaver
	itemChan, err := itemsaver.ItemSaver(fmt.Sprintf("%s", *itemSaverHost))
	if err != nil{
		panic(err)
	}

	pool := createClientPool(strings.Split(*workerHosts, ","))

	processor := worker.CreateProcessor(pool)



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

func createClientPool(hosts []string) chan *rpc.Client{
	var clients []*rpc.Client
	for _,h := range hosts{
		client, err := rpcsupport.NewClient(h)
		if err == nil{
			clients = append(clients, client)
			log.Printf("Connected to %s", h)
		}else{
			log.Printf("Error connected to %s: %v", h, err)
		}
	}

	out := make(chan *rpc.Client)
	go func() {
		for{
			for _, client := range  clients{
				out <- client
			}
		}
	}()
	return out
}