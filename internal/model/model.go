package model

type ZenquoteResponse struct {
	Content string `json:"q"`
	Author  string `json:"a"`
}

type QuotableResponse struct {
	Content string `json:"content"`
	Author  string `json:"author"`
}

var ZENQUOTES_URL = "https://zenquotes.io/api/random"
var QUOTABLE_URL = "https://api.quotable.io/random"
