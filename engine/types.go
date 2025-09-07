package engine

type ParserFunc func(*Request) *Result

type Parser interface {
	Parse(*Request) *Result
	Serialize() (string, map[string]any)
}

type Request struct {
	Url     string
	Data    []byte         // got by fetcher from Url
	Payload map[string]any // Parser may use Payload
	Parser  Parser
}

type Result struct {
	Requests []Request
	Items    []any
}
