package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/aggregators-api-service/mcp-server/config"
	"github.com/aggregators-api-service/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func GetpodcastepisodesHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		podcastKeyVal, ok := args["podcastKey"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: podcastKey"), nil
		}
		podcastKey, ok := podcastKeyVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: podcastKey"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["limit"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("limit=%v", val))
		}
		if val, ok := args["offset"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("offset=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/api/v2/podcasts/%s/episodes%s", cfg.BaseURL, podcastKey, queryString)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication based on auth type
		// Fallback to single auth parameter
		if cfg.APIKey != "" {
			req.Header.Set("x-zeno-api-key", cfg.APIKey)
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
		var result models.PodcastEpisodeList
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

func CreateGetpodcastepisodesTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_api_v2_podcasts_podcastKey_episodes",
		mcp.WithDescription("Get podcast episodes"),
		mcp.WithString("podcastKey", mcp.Required(), mcp.Description("")),
		mcp.WithString("limit", mcp.Description("")),
		mcp.WithString("offset", mcp.Description("")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    GetpodcastepisodesHandler(cfg),
	}
}
