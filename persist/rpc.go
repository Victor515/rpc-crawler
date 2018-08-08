package persist

import (
	"crawler/engine"
	"gopkg.in/olivere/elastic.v5"
	"crawler/persist"
	"log"
)

type ItemSaverService struct {
	Clinet *elastic.Client
	Index  string
}

func (s *ItemSaverService) Save(item engine.Item, result *string) error{
	log.Printf("Item %v saved", item)
	err := persist.Save(s.Clinet, item, s.Index)

	if err == nil{
		*result = "ok"
	}else{
		log.Printf("error occured when saving item %v: %v", item, err)
	}

	return err
}
