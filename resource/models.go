package resource

import datadog "github.com/zorkian/go-datadog-api"

type Source struct {
	ApiKey         string   `json:"api_key"`
	ApplicationKey string   `json:"application_key"`
	Priority       string   `json:"priority"`
	Sources        []string `json:"sources"`
	Tags           []string `json:"tags"`
	TitlePrefix    string   `json:"title_prefix"`
}

type Params struct {
	Event    datadog.Event `json:"version"`
	Priority string        `json:"priority"`
	Sources  []string      `json:"sources"`
	Tags     []string      `json:"tags"`
}

type InOutResponse struct {
	Version  Version  `json:"version"`
	Metadata Metadata `json:"metadata"`
}

type Version struct {
	Id int `json:"id"`
}

type CheckResponse []Version
type Metadata []MetadataField

type MetadataField struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
