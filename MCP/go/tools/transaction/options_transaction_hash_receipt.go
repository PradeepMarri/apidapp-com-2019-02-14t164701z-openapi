package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/apidapp/mcp-server/config"
	"github.com/apidapp/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func Options_transaction_hash_receiptHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		hashVal, ok := args["hash"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: hash"), nil
		}
		hash, ok := hashVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: hash"), nil
		}
		url := fmt.Sprintf("%s/transaction/%s/receipt", cfg.BaseURL, hash)
		req, err := http.NewRequest("OPTIONS", url, nil)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication based on auth type
		// Fallback to single auth parameter
		if cfg.APIKey != "" {
			req.Header.Set("X-Api-Key", cfg.APIKey)
		}
		req.Header.Set("Accept", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Request failed", err), nil
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to read response body", err), nil
		}

		if resp.StatusCode >= 400 {
			return mcp.NewToolResultError(fmt.Sprintf("API error: %s", body)), nil
		}
		// Use properly typed response
		var result models.Empty
		if err := json.Unmarshal(body, &result); err != nil {
			// Fallback to raw text if unmarshaling fails
			return mcp.NewToolResultText(string(body)), nil
		}

		prettyJSON, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to format JSON", err), nil
		}

		return mcp.NewToolResultText(string(prettyJSON)), nil
	}
}

func CreateOptions_transaction_hash_receiptTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("options_transaction_hash_receipt",
		mcp.WithDescription(""),
		mcp.WithString("hash", mcp.Required(), mcp.Description("Automatically added")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Options_transaction_hash_receiptHandler(cfg),
	}
}
