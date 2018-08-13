package main

import (
	"rpc-crawler/rpcsupport"
	"fmt"
	"rpc-crawler/config"
	"rpc-crawler/worker"
	"log"
)

func main() {
	log.Fatal(rpcsupport.ServeRpc(fmt.Sprintf(":%d",
		config.WorkerPort0), worker.CrawlerService{}))
}
