package main

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/tscolari/concourse-datadog-event-resource/resource"
	datadog "github.com/zorkian/go-datadog-api"
)

func main() {
	storePath := os.Args[1]
	input, err := resource.ParseInput(os.Stdin)
	must(err)

	client := datadog.NewClient(input.Source.Auth.ApiKey, input.Source.Auth.ApplicationKey)
	ddResource := resource.NewDatadog(client)

	response, event, err := ddResource.Out(input)
	eventFile, err := os.Create(filepath.Join(storePath, "event.json"))
	must(err)

	must(json.NewEncoder(eventFile).Encode(event))
	must(json.NewEncoder(os.Stdout).Encode(response))
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
