package main

import (
	"crawler_distributed/rpcsupport"
	"crawler_distributed/persist"
	"gopkg.in/olivere/elastic.v5"
	"log"
)

func main() {
	log.Fatal(rpcsupport.ServeRpc("dating_profile", ":1234"))
}

func ServeRpc(index, host string) error{
	// start up a new client
	client, err := elastic.NewClient(
		elastic.SetSniff(false),
	)
	if err != nil{
		return err
	}

	return rpcsupport.ServeRpc(host, &persist.ItemSaverService{
		Clinet: client,
		Index: index,
	})
}
