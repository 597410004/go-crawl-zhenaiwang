package types

func NilParseResult([]byte) ParseResult {
	return ParseResult{}
}

type ParseResult struct {
	Requests []Request
	Items    []interface{}
}

type Request struct {
	Url       string
	ParseFunc func([]byte) ParseResult
}
