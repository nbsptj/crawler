package main

import (
	"crawler/config"
	"crawler/dongchedi"
	"crawler/engine"
)

func main() {
	e := &engine.SimpleEngine{}

	e.Run(engine.Request{
		Url: "https://www.dongchedi.com",
		Parser: engine.NewFuncParser(
			dongchedi.ParseModel,
			config.PARSE_MODEL_FUNC_NAME,
		),
	})
}
