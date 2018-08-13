package main

import (
	"rpc-crawler/rpcsupport"
	"fmt"
	"rpc-crawler/worker"
	"log"
	"flag"
)

var port = flag.Int("port", 0, "the port to listen on")

func main() {
	flag.Parse()
	if *port == 0{
		fmt.Println("Must specify a port")
		return
	}
	log.Fatal(rpcsupport.ServeRpc(fmt.Sprintf(":%d",
		*port), worker.CrawlerService{}))
}
