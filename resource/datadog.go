package resource

import (
	"fmt"
	"strings"
	"time"

	datadog "github.com/zorkian/go-datadog-api"
)

type DatadogClient interface {
	GetEvents(start, end int, priority, sources, tags string) ([]datadog.Event, error)
	GetEvent(id int) (*datadog.Event, error)
	PostEvent(event *datadog.Event) (*datadog.Event, error)
}

type Datadog struct {
	client DatadogClient
}

func NewDatadog(client DatadogClient) *Datadog {
	return &Datadog{
		client: client,
	}
}

func (d *Datadog) Check(input Input) (CheckResponse, error) {
	end := time.Now()
	start := end.AddDate(0, 0, -15)

	events, err := d.client.GetEvents(
		int(start.Unix()),
		int(end.Unix()),
		input.Priority(),
		input.Sources(),
		input.Tags(),
	)

	if err != nil {
		return CheckResponse{}, fmt.Errorf("fetching events: %s", err)
	}

	checkResponse := CheckResponse{}
	titlePrefix := input.Source.TitlePrefix
	for _, event := range events {
		if strings.HasPrefix(event.Title, titlePrefix) {
			checkResponse = append(checkResponse, Version{Id: event.Id})
		}
	}

	return checkResponse, nil
}

func (d *Datadog) In(input Input) (InOutResponse, datadog.Event, error) {
	event, err := d.client.GetEvent(input.Version.Id)
	if err != nil {
		return InOutResponse{}, datadog.Event{}, err
	}

	if !strings.HasPrefix(event.Title, input.Source.TitlePrefix) {
		return InOutResponse{}, datadog.Event{}, fmt.Errorf("event `id:%d` doesn't match title prefix", event.Id)
	}

	response := InOutResponse{
		Version:  Version{Id: event.Id},
		Metadata: d.eventToMetadata(*event),
	}

	return response, *event, nil
}

func (d *Datadog) Out(input Input) (InOutResponse, datadog.Event, error) {
	event := input.Params.Event
	if event.Time == 0 {
		event.Time = int(time.Now().Unix())
	}

	if input.Params.Event.Tags == nil && len(input.Params.Event.Tags) == 0 {
		event.Tags = input.Source.Tags
	}

	if input.Params.Event.Priority == "" {
		event.Priority = input.Source.Priority
	}

	if input.Source.TitlePrefix != "" {
		event.Title = fmt.Sprintf("%s %s", input.Source.TitlePrefix, event.Title)
	}

	returnedEvent, err := d.client.PostEvent(&event)
	if err != nil {
		return InOutResponse{}, datadog.Event{}, fmt.Errorf("creating event: %s", err)
	}

	response := InOutResponse{
		Version:  Version{Id: returnedEvent.Id},
		Metadata: d.eventToMetadata(*returnedEvent),
	}

	return response, *returnedEvent, nil
}

func (d *Datadog) eventToMetadata(event datadog.Event) Metadata {
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
