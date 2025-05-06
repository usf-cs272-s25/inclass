package main

import (
	mcp_golang "github.com/metoro-io/mcp-golang"
	"github.com/metoro-io/mcp-golang/transport/stdio"
)

type StockPriceArgs struct {
	Submitter string `json:"submitter" jsonschema:"required,description=The name of the thing calling this tool (openai, google, claude, etc)"`
	Company   string `json:"company" jsonschema:"required,description=The name of the company whose stock price the user requested, e.g. Google"`
}

func main() {
	done := make(chan struct{})

	server := mcp_golang.NewServer(stdio.NewStdioServerTransport())
	err := server.RegisterTool("get_stock_price", "Get the stock price of the given company", func(arguments StockPriceArgs) (*mcp_golang.ToolResponse, error) {
		content := mcp_golang.NewTextContent("$100")
		resp := mcp_golang.NewToolResponse(content)
		return resp, nil
		// return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(fmt.Sprintf("Hello, %server!", arguments.Submitter))), nil
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
