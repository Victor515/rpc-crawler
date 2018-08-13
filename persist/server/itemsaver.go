package main

import (
	"rpc-crawler/rpcsupport"
	"rpc-crawler/persist"
	"gopkg.in/olivere/elastic.v5"
	"log"
	"rpc-crawler/config"
	"fmt"
	"flag"
)

var port = flag.Int("port", 0, "the port to listen on")

func main() {
	flag.Parse()
	if *port == 0{
		fmt.Println("Must specify a port")
		return
	}
	log.Fatal(ServeRpc(config.ElasticIndex, fmt.Sprintf(":%d", *port)))
}

func ServeRpc(index, host string) error{
	// start up a new client
	client, err := elastic.NewClient(
		elastic.SetSniff(false),
		elastic.SetURL(config.ElasticUrl),
	)
	if err != nil{
		return err
	}


	return rpcsupport.ServeRpc(host, &persist.ItemSaverService{
		Clinet: client,
		Index: index,
	})
}
