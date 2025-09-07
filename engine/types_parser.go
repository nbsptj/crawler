package engine

import "crawler/config"

type FuncParser struct {
	parser ParserFunc
	name   string
}

func (f *FuncParser) Parse(r *Request) *Result {
	return f.parser(r)
}

func (f *FuncParser) Serialize() (string, map[string]any) {
	return f.name, nil
}

func NewFuncParser(parser ParserFunc, name string) *FuncParser {
	return &FuncParser{
		parser, name,
	}
}

type NilParser struct{}

func (NilParser) Parse(r *Request) *Result {
	return &Result{}
}

func (NilParser) Serialize() (string, map[string]any) {
	return config.NIL_PARSER, nil
}
