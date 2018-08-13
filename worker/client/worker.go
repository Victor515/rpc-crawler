package client

import (
	"crawler/engine"
	"rpc-crawler/config"
	"rpc-crawler/worker"
	"net/rpc"
)

func CreateProcessor(clientChan chan *rpc.Client) (engine.Processor){
	return func(request engine.Request) (engine.ParserResult, error) {
		// serialize request
		sReq := worker.SerializeRequest(request)
		var sRes worker.ParserResult

		c := <-clientChan
		// call rpc
		err := c.Call(config.CrawlerServiceRpc, sReq, &sRes)
		if err != nil{
			return engine.ParserResult{}, err
		}
		// return deserialized result
		return worker.DeserializeResult(sRes), nil
	}
}
