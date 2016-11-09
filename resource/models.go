package resource

import datadog "github.com/zorkian/go-datadog-api"

type Source struct {
	Priority       string   `json:"priority"`
	DatadogUrl     string   `json:"datadog_url"`
	ApiKey         string   `json:"api_key"`
	ApplicationKey string   `json:"application_key"`
	TitlePrefix    string   `json:"title_prefix"`
	Sources        []string `json:"sources"`
	Tags           []string `json:"tags"`
}

type Params struct {
	Event    datadog.Event `json:"version"`
	Sources  []string      `json:"sources"`
	Tags     []string      `json:"tags"`
	Priority string        `json:"priority"`
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
