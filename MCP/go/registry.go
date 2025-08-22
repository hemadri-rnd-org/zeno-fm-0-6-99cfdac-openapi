package main

import (
	"github.com/aggregators-api-service/mcp-server/config"
	"github.com/aggregators-api-service/mcp-server/models"
	tools_api_v2 "github.com/aggregators-api-service/mcp-server/tools/api_v2"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_api_v2.CreateDeletepodcastTool(cfg),
		tools_api_v2.CreateGetpodcastTool(cfg),
		tools_api_v2.CreateGetpartneraggregatorstationsTool(cfg),
		tools_api_v2.CreateGetpodcastcategoriesTool(cfg),
		tools_api_v2.CreateGetpodcastepisodesTool(cfg),
		tools_api_v2.CreateDeletepodcast_1Tool(cfg),
		tools_api_v2.CreateGetpodcastepisodeTool(cfg),
		tools_api_v2.CreateGetstationgenresTool(cfg),
		tools_api_v2.CreateSearchstationsTool(cfg),
		tools_api_v2.CreateGetpodcastlanguagesTool(cfg),
		tools_api_v2.CreateSearchpodcastsTool(cfg),
		tools_api_v2.CreateGetstationcountriesTool(cfg),
		tools_api_v2.CreateGetstationlanguagesTool(cfg),
		tools_api_v2.CreateGetpodcastcountriesTool(cfg),
	}
}
