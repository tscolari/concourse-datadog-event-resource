package params_test

import (
	"github.com/tscolari/concourse-datadog-event-resource/params"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Input", func() {
	var input params.Input

	Describe("Sources", func() {
		Context("when only `Source` is supplied", func() {
			BeforeEach(func() {
				input = params.Input{
					Source: params.Source{
						Sources: []string{
							"ci",
							"dev",
						},
					},
				}
			})

			It("returns the sources as a string", func() {
				Expect(input.Sources()).To(Equal("ci,dev"))
			})
		})

		Context("when only `Params` is supplied", func() {
			BeforeEach(func() {
				input = params.Input{
					Params: params.Params{
						Sources: []string{
							"ci",
							"dev",
						},
					},
				}
			})

			It("returns the sources as a string", func() {
				Expect(input.Sources()).To(Equal("ci,dev"))
			})
		})
	})

	Context("when both Source and Params is supplied", func() {
		BeforeEach(func() {
			input = params.Input{
				Params: params.Params{
					Sources: []string{
						"ci",
						"dev",
					},
				},
				Source: params.Source{
					Sources: []string{
						"ci",
						"production",
					},
				},
			}
		})

		It("ensures Params takes precedence", func() {
			Expect(input.Sources()).To(Equal("ci,dev"))
		})
	})

	Describe("Tags", func() {
		Context("when only `Source` is supplied", func() {
			BeforeEach(func() {
				input = params.Input{
					Source: params.Source{
						Tags: []string{
							"ci",
							"dev",
						},
					},
				}
			})

			It("returns the tags as a string", func() {
				Expect(input.Tags()).To(Equal("ci,dev"))
			})
		})

		Context("when only `Params` is supplied", func() {
			BeforeEach(func() {
				input = params.Input{
					Params: params.Params{
						Tags: []string{
							"ci",
							"dev",
						},
					},
				}
			})

			It("returns the tags as a string", func() {
				Expect(input.Tags()).To(Equal("ci,dev"))
			})
		})
	})

	Context("when both Source and Params is supplied", func() {
		BeforeEach(func() {
			input = params.Input{
				Params: params.Params{
					Tags: []string{
						"ci",
						"dev",
					},
				},
				Source: params.Source{
					Tags: []string{
						"ci",
						"production",
					},
				},
			}
		})

		It("ensures Params takes precedence", func() {
			Expect(input.Tags()).To(Equal("ci,dev"))
		})
	})
})
