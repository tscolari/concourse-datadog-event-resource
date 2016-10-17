package params

import "strings"

type Input struct {
	Source  Source  `json:"source"`
	Params  Params  `json:"params"`
	Version Version `json:"version"`
}

func (i *Input) Sources() string {
	sources := ""
	if len(i.Source.Sources) != 0 {
		sources = strings.Join(i.Source.Sources, ",")
	}

	if len(i.Params.Sources) != 0 {
		sources = strings.Join(i.Params.Sources, ",")
	}

	return sources
}

func (i *Input) Tags() string {
	tags := ""

	if len(i.Source.Tags) != 0 {
		tags = strings.Join(i.Source.Tags, ",")
	}

	if len(i.Params.Tags) != 0 {
		tags = strings.Join(i.Params.Tags, ",")
	}
	return tags
}
