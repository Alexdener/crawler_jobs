package parser

import (
	"crawler_jobs/engine"
)

type FuncParser struct {
	parser engine.ParserFunc // 对应解析函数
	name   string            // 函数名
}

func NewFuncParser(f engine.ParserFunc, name string) FuncParser {
	return FuncParser{
		parser: f,
		name:   name,
	}
}

func (f FuncParser) Parse(bytes []byte, url string) engine.ParseResult {
	return f.parser(bytes, url)
}
