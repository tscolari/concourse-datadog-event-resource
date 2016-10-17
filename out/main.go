package main

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/tscolari/concourse-datadog-event-resource/params"
	datadog "github.com/zorkian/go-datadog-api"
)

func main() {
	input, err := params.ParseInput(os.Stdin)
	must(err)

	client := datadog.NewClient(input.Source.ApiKey, input.Source.ApplicationKey)

	event := &input.Params.Event
	if _, err := os.Stat(event.Title); err == nil {
		title, err := ioutil.ReadFile(event.Title)
		must(err)

		event.Title = string(title)
	}

	if _, err := os.Stat(event.Text); err == nil {
		text, err := ioutil.ReadFile(event.Text)
		must(err)

		event.Text = string(text)
	}

	event, err = client.PostEvent(event)
	must(err)

	response := params.InOutResponse{
		Version:  params.Version{Id: event.Id},
		Metadata: params.EventToMetadata(*event),
	}

	err = json.NewEncoder(os.Stdout).Encode(response)
	must(err)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
