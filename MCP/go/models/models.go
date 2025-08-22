package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// StationList represents the StationList schema from the OpenAPI specification
type StationList struct {
	Total int64 `json:"total,omitempty"`
	Items []Station `json:"items,omitempty"`
}

// PodcastSearchParams represents the PodcastSearchParams schema from the OpenAPI specification
type PodcastSearchParams struct {
	Page int `json:"page,omitempty"`
	Query string `json:"query,omitempty"`
	Filters PodcastFilters `json:"filters,omitempty"` // Filters for podcast search
	Hitsperpage int `json:"hitsPerPage,omitempty"`
}

// Station represents the Station schema from the OpenAPI specification
type Station struct {
	Key string `json:"key,omitempty"`
	Stream string `json:"stream,omitempty"`
	Website string `json:"website,omitempty"`
	Description string `json:"description,omitempty"`
	Logo string `json:"logo,omitempty"`
	Country string `json:"country,omitempty"`
	Genre string `json:"genre,omitempty"`
	Language string `json:"language,omitempty"`
	Name string `json:"name,omitempty"`
	City string `json:"city,omitempty"`
}

// StationSearchResults represents the StationSearchResults schema from the OpenAPI specification
type StationSearchResults struct {
	Hits []Station `json:"hits,omitempty"`
	Total int64 `json:"total,omitempty"`
}

// PodcastCategory represents the PodcastCategory schema from the OpenAPI specification
type PodcastCategory struct {
	Id string `json:"id,omitempty"`
	Parent string `json:"parent,omitempty"`
	Text string `json:"text,omitempty"`
}

// PodcastEpisodeList represents the PodcastEpisodeList schema from the OpenAPI specification
type PodcastEpisodeList struct {
	Items []PodcastEpisode `json:"items,omitempty"`
	Total int64 `json:"total,omitempty"`
}

// PodcastSearchResults represents the PodcastSearchResults schema from the OpenAPI specification
type PodcastSearchResults struct {
	Hits []Podcast `json:"hits,omitempty"`
	Total int64 `json:"total,omitempty"`
}

// StationGenre represents the StationGenre schema from the OpenAPI specification
type StationGenre struct {
	Name string `json:"name,omitempty"`
}

// PodcastEpisode represents the PodcastEpisode schema from the OpenAPI specification
type PodcastEpisode struct {
	Season int64 `json:"season,omitempty"`
	Tags []string `json:"tags,omitempty"`
	Author string `json:"author,omitempty"`
	Explicit bool `json:"explicit,omitempty"`
	Fileurl string `json:"fileUrl,omitempty"`
	Summary string `json:"summary"`
	Image string `json:"image,omitempty"`
	Publishdate string `json:"publishDate"`
	Key string `json:"key,omitempty"`
	Link string `json:"link,omitempty"`
	Episodetype string `json:"episodeType,omitempty"`
	Description string `json:"description"`
	Title string `json:"title"`
	Duration int64 `json:"duration,omitempty"`
	Episode int64 `json:"episode,omitempty"`
	Subtitle string `json:"subtitle,omitempty"`
	Size int64 `json:"size,omitempty"`
	Block bool `json:"block,omitempty"`
}

// PodcastFilters represents the PodcastFilters schema from the OpenAPI specification
type PodcastFilters struct {
	Language []string `json:"language,omitempty"`
	Podcasttype string `json:"podcastType,omitempty"`
	Category []string `json:"category,omitempty"`
	Country []string `json:"country,omitempty"`
}

// StationFilters represents the StationFilters schema from the OpenAPI specification
type StationFilters struct {
	Country []string `json:"country,omitempty"`
	Genre []string `json:"genre,omitempty"`
	Language []string `json:"language,omitempty"`
}

// StationSearchParams represents the StationSearchParams schema from the OpenAPI specification
type StationSearchParams struct {
	Filters StationFilters `json:"filters,omitempty"` // Filters for station search
	Hitsperpage int `json:"hitsPerPage,omitempty"`
	Page int `json:"page,omitempty"`
	Query string `json:"query,omitempty"`
}

// Language represents the Language schema from the OpenAPI specification
type Language struct {
	Name string `json:"name,omitempty"`
}

// Country represents the Country schema from the OpenAPI specification
type Country struct {
	Name string `json:"name,omitempty"`
	Iso string `json:"iso,omitempty"`
	Iso3 string `json:"iso3,omitempty"`
}

// Podcast represents the Podcast schema from the OpenAPI specification
type Podcast struct {
	Description string `json:"description"`
	Owneremail string `json:"ownerEmail,omitempty"`
	Subtitle string `json:"subtitle,omitempty"`
	Showtype string `json:"showType,omitempty"`
	Country string `json:"country,omitempty"`
	Title string `json:"title"`
	Keywords []string `json:"keywords,omitempty"`
	Link string `json:"link,omitempty"`
	Copyright string `json:"copyright,omitempty"`
	Explicit bool `json:"explicit,omitempty"`
	Image string `json:"image,omitempty"`
	Author string `json:"author,omitempty"`
	Categories []string `json:"categories"`
	Block bool `json:"block,omitempty"`
	Key string `json:"key,omitempty"`
	Ownername string `json:"ownerName,omitempty"`
	Language string `json:"language"`
	Summary string `json:"summary"`
}
