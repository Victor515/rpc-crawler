package main

import (
	"rpc-crawler/rpcsupport"
	"rpc-crawler/persist"
	"gopkg.in/olivere/elastic.v5"
	"log"
	"rpc-crawler/config"
	"fmt"
)

func main() {
	log.Fatal(ServeRpc(config.ElasticIndex, fmt.Sprintf(":%d", config.ItemSaverPort)))
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
