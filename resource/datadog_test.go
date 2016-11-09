package resource_test

import (
	"errors"
	"time"

	"github.com/tscolari/concourse-datadog-event-resource/resource"
	"github.com/tscolari/concourse-datadog-event-resource/resource/resourcefakes"
	datadog "github.com/zorkian/go-datadog-api"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Datadog", func() {
	var (
		datadogClient *resourcefakes.FakeDatadogClient
		ddResource    *resource.Datadog
		input         resource.Input
	)

	BeforeEach(func() {
		datadogClient = new(resourcefakes.FakeDatadogClient)
		ddResource = resource.NewDatadog(datadogClient)
		input = resource.Input{
			Source: resource.Source{
				Priority: "high",
				Sources:  []string{"source_source1", "source_source2"},
				Tags:     []string{"source_tag1", "source_tag2"},
			},
			Params: resource.Params{
				Priority: "low",
				Sources:  []string{"params_source1", "params_source2"},
				Tags:     []string{"params_tag1", "params_tag2"},
			},
		}
	})

	Describe("Check", func() {
		It("asks for the events in the right time interval", func() {
			now := int(time.Now().Unix())
			_, err := ddResource.Check(resource.Input{})
			Expect(err).NotTo(HaveOccurred())
			Expect(datadogClient.GetEventsCallCount()).To(Equal(1))
			start, end, _, _, _ := datadogClient.GetEventsArgsForCall(0)

			Expect(end).To(BeNumerically("~", now, 100))
			Expect(end - start).To(Equal(1296000))
		})

		It("asks for the events with correct priority", func() {
			_, err := ddResource.Check(input)
			Expect(err).NotTo(HaveOccurred())

			Expect(datadogClient.GetEventsCallCount()).To(Equal(1))
			_, _, priority, _, _ := datadogClient.GetEventsArgsForCall(0)
			Expect(priority).To(Equal("low"))
		})

		It("asks for the events with correct sources", func() {
			_, err := ddResource.Check(input)
			Expect(err).NotTo(HaveOccurred())

			Expect(datadogClient.GetEventsCallCount()).To(Equal(1))
			_, _, _, sources, _ := datadogClient.GetEventsArgsForCall(0)
			Expect(sources).To(Equal("params_source1,params_source2"))
		})

		It("asks for the events with correct tags", func() {
			_, err := ddResource.Check(input)
			Expect(err).NotTo(HaveOccurred())

			Expect(datadogClient.GetEventsCallCount()).To(Equal(1))
			_, _, _, _, tags := datadogClient.GetEventsArgsForCall(0)
			Expect(tags).To(Equal("params_tag1,params_tag2"))
		})

		It("returns the correct check response", func() {
			datadogClient.GetEventsReturns([]datadog.Event{
				datadog.Event{Id: 1234},
				datadog.Event{Id: 5678},
			}, nil)

			events, err := ddResource.Check(input)
			Expect(err).NotTo(HaveOccurred())
			Expect(events).To(ConsistOf(resource.CheckResponse{
				resource.Version{Id: 1234},
				resource.Version{Id: 5678},
			}))
		})

		Context("when the client returns an error", func() {
			It("returns the error", func() {
				datadogClient.GetEventsReturns([]datadog.Event{}, errors.New("failed badly"))
				_, err := ddResource.Check(resource.Input{})

				Expect(err).To(MatchError(ContainSubstring("failed badly")))
			})
		})
	})

	Describe("In", func() {
		var (
			event   datadog.Event
			timeNow time.Time
		)

		BeforeEach(func() {
			timeNow = time.Now()

			event = datadog.Event{
				Id:         12345,
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

			datadogClient.GetEventReturns(&event, nil)
		})

		It("fetches the event with the correct ID", func() {
			input := resource.Input{Version: resource.Version{Id: 12345}}
			_, _, err := ddResource.In(input)
			Expect(err).NotTo(HaveOccurred())

			Expect(datadogClient.GetEventCallCount()).To(Equal(1))
			id := datadogClient.GetEventArgsForCall(0)

			Expect(id).To(Equal(12345))
		})

		It("returns the correct response", func() {
			response, _, err := ddResource.In(input)
			Expect(err).NotTo(HaveOccurred())

			Expect(response.Version.Id).To(Equal(12345))
			Expect(response.Metadata).To(ConsistOf(resource.Metadata{
				resource.MetadataField{Name: "Title", Value: event.Title},
				resource.MetadataField{Name: "Text", Value: event.Text},
				resource.MetadataField{Name: "Time", Value: timeNow.Format(time.RFC3339)},
				resource.MetadataField{Name: "Priority", Value: event.Priority},
				resource.MetadataField{Name: "Alert Type", Value: event.AlertType},
				resource.MetadataField{Name: "Host", Value: event.Host},
				resource.MetadataField{Name: "Source Type", Value: event.SourceType},
				resource.MetadataField{Name: "Tags", Value: "hello, bye"},
				resource.MetadataField{Name: "Resource", Value: event.Resource},
			}))
		})

		It("returns the fetched event", func() {
			_, returnedEvent, err := ddResource.In(input)
			Expect(err).NotTo(HaveOccurred())
			Expect(returnedEvent).To(Equal(event))
		})

		Context("when the client returns an error", func() {
			It("returns an error", func() {
				datadogClient.GetEventReturns(nil, errors.New("failed"))

				_, _, err := ddResource.In(input)
				Expect(err).To(MatchError(ContainSubstring("failed")))
			})
		})
	})

	Describe("Out", func() {
		var (
			timeNow time.Time
			event   datadog.Event
		)

		BeforeEach(func() {
			timeNow = time.Now()
			input.Params.Event = datadog.Event{
				Title:      "Hello",
				Text:       "World",
				Priority:   "not-used",
				Tags:       []string{"not", "used"},
				Host:       "my-test.com",
				AlertType:  "Some type",
				SourceType: "Another type",
				Time:       int(timeNow.Unix()),
				Resource:   "that one",
			}

			event = datadog.Event{
				Id:         12345,
				Title:      "returned title",
				Text:       "returned text",
				Priority:   "returned priority",
				Tags:       []string{"returned", "tags"},
				Host:       "return.com",
				AlertType:  "return type",
				SourceType: "return type",
				Time:       int(timeNow.Unix()),
				Resource:   "return one",
			}

			datadogClient.PostEventReturns(&event, nil)
		})

		It("sends the correct event to the client", func() {
			_, _, err := ddResource.Out(input)
			Expect(err).NotTo(HaveOccurred())

			Expect(datadogClient.PostEventCallCount()).To(Equal(1))
			event := datadogClient.PostEventArgsForCall(0)

			Expect(event.Title).To(Equal(input.Params.Event.Title))
			Expect(event.Text).To(Equal(input.Params.Event.Text))
			Expect(event.Priority).To(Equal(input.Params.Priority))
			Expect(event.Tags).To(Equal(input.Params.Tags))
			Expect(event.Resource).To(Equal(input.Params.Event.Resource))
			Expect(event.Url).To(Equal(input.Params.Event.Url))
			Expect(event.Host).To(Equal(input.Params.Event.Host))
			Expect(event.AlertType).To(Equal(input.Params.Event.AlertType))
			Expect(event.Aggregation).To(Equal(input.Params.Event.Aggregation))
			Expect(event.Time).To(Equal(input.Params.Event.Time))
		})

		It("returns the correct response", func() {
			response, _, err := ddResource.Out(input)
			Expect(err).NotTo(HaveOccurred())

			Expect(response.Version.Id).To(Equal(12345))
			Expect(response.Metadata).To(ConsistOf(resource.Metadata{
				resource.MetadataField{Name: "Title", Value: "returned title"},
				resource.MetadataField{Name: "Text", Value: "returned text"},
				resource.MetadataField{Name: "Time", Value: timeNow.Format(time.RFC3339)},
				resource.MetadataField{Name: "Priority", Value: "returned priority"},
				resource.MetadataField{Name: "Alert Type", Value: "return type"},
				resource.MetadataField{Name: "Host", Value: "return.com"},
				resource.MetadataField{Name: "Source Type", Value: "return type"},
				resource.MetadataField{Name: "Tags", Value: "returned, tags"},
				resource.MetadataField{Name: "Resource", Value: "return one"},
			}))
		})

		It("returns the correct event", func() {
			_, returnedEvent, err := ddResource.Out(input)
			Expect(err).NotTo(HaveOccurred())
			Expect(returnedEvent).To(Equal(event))
		})

		Context("when the client returns an error", func() {
			It("returns an error", func() {
				datadogClient.PostEventReturns(nil, errors.New("failed"))

				_, _, err := ddResource.Out(input)
				Expect(err).To(MatchError(ContainSubstring("failed")))
			})
		})
	})
})
