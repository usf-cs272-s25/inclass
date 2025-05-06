package main

import (
	mcp_golang "github.com/metoro-io/mcp-golang"
	"github.com/metoro-io/mcp-golang/transport/stdio"
)

type StockArgs struct {
	Submitter string `json:"submitter" jsonschema:"required,description=The name of the thing calling this tool (openai, google, claude, etc)"`
	Company   string `json:"content" jsonschema:"required,description=The name of the company whose stock price the user wants to know"`
}

func main() {
	done := make(chan struct{})

	server := mcp_golang.NewServer(stdio.NewStdioServerTransport())
	err := server.RegisterTool("get_stock_price", "Gets the stock price for the requested company", func(arguments StockArgs) (*mcp_golang.ToolResponse, error) {
		text := mcp_golang.NewTextContent(("$250 per share"))
		resp := mcp_golang.NewToolResponse(text)
		return resp, nil
	})
	if err != nil {
		panic(err)
	}

	err = server.Serve()
	if err != nil {
		panic(err)
	}

	<-done
}
