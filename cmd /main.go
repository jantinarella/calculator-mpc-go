package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func main() {
	fmt.Println("Hello, World!")
	server := server.NewMCPServer("Calculator", "1.0.0")

	add := mcp.NewTool("add",
		mcp.WithDescription("Add two numbers"),
		mcp.WithNumber("a",
			mcp.Required(),
			mcp.Description("First number"),
		),
		mcp.WithNumber("b",
			mcp.Required(),
			mcp.Description("Second number"),
		),
	)

	server.AddTool(add, addHandler)

}

func addHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	a, ok := request.Params.Arguments["a"].(float64)
	if !ok {
		return nil, errors.New("first number must be a number")
	}

	b, ok := request.Params.Arguments["b"].(float64)
	if !ok {
		return nil, errors.New("second number must be a number")
	}

	result := a + b
	return mcp.NewToolResultText(fmt.Sprintf("%.2f", result)), nil
}
