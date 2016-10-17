package params

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"time"

	datadog "github.com/zorkian/go-datadog-api"
)

func ParseInput(inputStream io.Reader) (Input, error) {
	var input Input
	if err := json.NewDecoder(inputStream).Decode(&input); err != nil {
		return input, fmt.Errorf("parsing input: %s", err.Error())
	}
	return input, nil
}

func EventToMetadata(event datadog.Event) Metadata {
	eventTime := time.Unix(int64(event.Time), 0)

	return Metadata{
		MetadataField{Name: "Title", Value: event.Title},
		MetadataField{Name: "Text", Value: event.Text},
		MetadataField{Name: "Time", Value: eventTime.Format(time.RFC3339)},
		MetadataField{Name: "Priority", Value: event.Priority},
		MetadataField{Name: "Alert Type", Value: event.AlertType},
		MetadataField{Name: "Host", Value: event.Host},
		MetadataField{Name: "Source Type", Value: event.SourceType},
		MetadataField{Name: "Tags", Value: strings.Join(event.Tags, ", ")},
		MetadataField{Name: "Resource", Value: event.Resource},
	}
}
