package client

import (
	"crawler/engine"
	"rpc-crawler/rpcsupport"
	"fmt"
	"rpc-crawler/config"
	"rpc-crawler/worker"
)

func CreateProcessor() (engine.Processor, error){
	client, err := rpcsupport.NewClient(fmt.Sprintf(":%d", config.WorkerPort0))
	if err != nil{
		return nil, err
	}
	return func(request engine.Request) (engine.ParserResult, error) {
		// serialize request
		sReq := worker.SerializeRequest(request)
		var sRes worker.ParserResult

		// call rpc
		err := client.Call(config.CrawlerServiceRpc, sReq, &sRes)
		if err != nil{
			return engine.ParserResult{}, err
		}
		// return deserialized result
		return worker.DeserializeResult(sRes), nil
	}, nil
}
