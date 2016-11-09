package resource_test

import (
	"encoding/json"

	"github.com/tscolari/concourse-datadog-event-resource/resource"
	datadog "github.com/zorkian/go-datadog-api"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
)

var _ = Describe("Parser", func() {
	Describe("ParseInput", func() {
		var (
			input  resource.Input
			buffer *gbytes.Buffer
		)

		JustBeforeEach(func() {
			buffer = gbytes.NewBuffer()
			Expect(json.NewEncoder(buffer).Encode(input)).To(Succeed())
		})

		BeforeEach(func() {
			input = resource.Input{
				Source: resource.Source{
					ApplicationKey: "app-key",
					ApiKey:         "api-key",
					Priority:       "normal",
					TitlePrefix:    "My Event: ",
				},
				Params: resource.Params{
					Event: datadog.Event{
						Title: "hello",
						Text:  "my event",
					},
				},
			}
		})

		It("correctly parses the source into a input object", func() {
			input, err := resource.ParseInput(buffer)
			Expect(err).ToNot(HaveOccurred())

			Expect(input.Source.ApiKey).To(Equal("api-key"))
			Expect(input.Source.ApplicationKey).To(Equal("app-key"))
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

			_, err := resource.ParseInput(buffer)
			Expect(err).To(MatchError(ContainSubstring("parsing input:")))
		})
	})
})
