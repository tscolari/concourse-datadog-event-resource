package params_test

import (
	"encoding/json"
	"time"

	"github.com/tscolari/concourse-datadog-event-resource/params"
	datadog "github.com/zorkian/go-datadog-api"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
)

var _ = Describe("Parser", func() {
	Describe("ParseInput", func() {
		var (
			input  params.Input
			buffer *gbytes.Buffer
		)

		JustBeforeEach(func() {
			buffer = gbytes.NewBuffer()
			Expect(json.NewEncoder(buffer).Encode(input)).To(Succeed())
		})

		BeforeEach(func() {
			input = params.Input{
				Source: params.Source{
					ApplicationKey: "app-key",
					ApiKey:         "api-key",
					DatadogUrl:     "http://endpoint",
					Priority:       "normal",
					TitlePrefix:    "My Event: ",
				},
				Params: params.Params{
					Event: datadog.Event{
						Title: "hello",
						Text:  "my event",
					},
				},
			}
		})

		It("correctly parses the source into a input object", func() {
			input, err := params.ParseInput(buffer)
			Expect(err).ToNot(HaveOccurred())

			Expect(input.Source.ApiKey).To(Equal("api-key"))
			Expect(input.Source.ApplicationKey).To(Equal("app-key"))
			Expect(input.Source.DatadogUrl).To(Equal("http://endpoint"))
			Expect(input.Source.Priority).To(Equal("normal"))
			Expect(input.Source.TitlePrefix).To(Equal("My Event: "))

			Expect(input.Params.Event.Title).To(Equal("hello"))
			Expect(input.Params.Event.Text).To(Equal("my event"))
		})
	})

	Context("when json parsing fails", func() {
		It("returns an error", func() {
			buffer := gbytes.NewBuffer()
			buffer.Write([]byte("not json!"))

			_, err := params.ParseInput(buffer)
			Expect(err).To(MatchError(ContainSubstring("parsing input:")))
		})
	})

	Describe("EventToMetadata", func() {
		It("correctly maps the event properties to a Metadata object", func() {
			timeNow := time.Now()
			event := datadog.Event{
				Title:      "hello",
				Text:       "hello 2",
				Time:       int(timeNow.Unix()),
				Tags:       []string{"hello", "bye"},
				Priority:   "Urgent",
				AlertType:  "Important",
				Host:       "my-server.com",
				SourceType: "my-source",
				Resource:   "resource",
			}

			metadata := params.EventToMetadata(event)
			Expect(metadata).To(ContainElement(params.MetadataField{Name: "Title", Value: event.Title}))
			Expect(metadata).To(ContainElement(params.MetadataField{Name: "Text", Value: event.Text}))
			Expect(metadata).To(ContainElement(params.MetadataField{Name: "Time", Value: timeNow.Format(time.RFC3339)}))
			Expect(metadata).To(ContainElement(params.MetadataField{Name: "Priority", Value: event.Priority}))
			Expect(metadata).To(ContainElement(params.MetadataField{Name: "Alert Type", Value: event.AlertType}))
			Expect(metadata).To(ContainElement(params.MetadataField{Name: "Host", Value: event.Host}))
			Expect(metadata).To(ContainElement(params.MetadataField{Name: "Source Type", Value: event.SourceType}))
			Expect(metadata).To(ContainElement(params.MetadataField{Name: "Tags", Value: "hello, bye"}))
			Expect(metadata).To(ContainElement(params.MetadataField{Name: "Resource", Value: event.Resource}))
		})
	})
})
