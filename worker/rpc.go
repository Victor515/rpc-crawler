package worker

import "crawler/engine"

type CrawlerService struct {}

func (s CrawlerService) Process(req Request, result * ParserResult) error{
	engineReq, err := DeserializeRequest(req)
	if err != nil{
		return err
	}
	
	// start worker
	engineResult, err := engine.Worker(engineReq)
	if err != nil{
		return err
	}
	*result = SerializeResult(engineResult)
	return nil
}
