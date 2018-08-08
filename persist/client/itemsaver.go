package client

import (
	"crawler/engine"
	"log"
	"crawler_distributed/rpcsupport"
)

func ItemSaver(host string) (chan engine.Item, error){
	client, err := rpcsupport.NewClient(host)

	if err != nil{
		return nil, err
	}

	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for{
			item := <- out
			log.Printf("Item Saver got Item%d: %v", itemCount, item)
			itemCount++

			// call rpc to save
			result := ""
			err := client.Call("ItemSaverService.Save", item, &result)

			if err != nil{
				log.Printf("Itemsaver error when saving item %v: %v", item, err)
			}
		}
	}()

	return out, nil
}
