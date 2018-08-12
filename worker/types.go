package worker

import (
	"crawler/engine"
	"crawler/config"
	"crawler/zhenai/parser"
	"errors"
	"fmt"
	"log"
)

type SerializedParser struct {
	Name string // function name
	Args interface{}
}

// request wrapper for rpc call
type Request struct {
	Url string
	Parser SerializedParser
}

type ParserResult struct {
	Items []engine.Item
	Requests []Request
}

func SerializeRequest(r engine.Request) Request{
	name, args := r.Parser.Serialize()
	return Request{
		Url:r.Url,
		Parser:SerializedParser{
			Name:name,
			Args:args,
		},
	}
}

func SerializeResult(r engine.ParserResult) ParserResult{
	result := ParserResult{
		Items:r.Items,
	}
	for _, req := range r.Requests{
		result.Requests = append(result.Requests, SerializeRequest(req))
	}
	return result
}

func DeserializeRequest(r Request) (engine.Request ,error){
	parser, err := deserializeParser(r.Parser)
	if err != nil{
		return engine.Request{}, err
	}else{
		return engine.Request{
			Url:r.Url,
			Parser:parser,
		}, nil
	}
}

func DeserializeResult(r ParserResult) engine.ParserResult{
	result := engine.ParserResult{
		Items: r.Items,
	}

	for _, req := range r.Requests{
		request, err := DeserializeRequest(req)
		if err != nil{
			log.Printf("error deserializing the request: %v", err)
			continue
		}
		result.Requests = append(result.Requests, request)
	}
	return result
}

func deserializeParser(p SerializedParser) (engine.Parser, error){
	switch p.Name {
	case config.ParseCityList:
		return engine.NewFuncParser(
			parser.ParseCityList,
			config.ParseCityList,
		), nil
	case config.ParseCity:
		return engine.NewFuncParser(
			parser.ParseCity,
			config.ParseCity,
		), nil
	case config.NilParser:
		return engine.NilParser{}, nil
	case config.ParseProfile:
		if userName, ok := p.Args.(string); ok{
			return parser.NewProfileParser(userName), nil
		}else{
			return nil, fmt.Errorf("invalid args: %v", p.Args)
		}
	default:
		return nil, errors.New("Cannot deserialize parser: " + p.Name)
	}
}
