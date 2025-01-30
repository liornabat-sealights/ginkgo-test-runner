package main

import (
	"crypto/tls"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/liornabat-sealights/ginkgo-test-runner/lib/types"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"time"
)

var _ = BeforeSuite(func() {
	time.Sleep(2 * time.Second) // Wait for server to start
})

var _ = Describe("Calculator Service", func() {

	Describe("ResultResponse Type", func() {
		Context("NewResultResponse", func() {
			It("should create a new empty ResultResponse", func() {
				response := types.NewResultResponse()
				Expect(response).To(Equal(&types.ResultResponse{}))
			})
		})

		Context("SetResult", func() {
			It("should set the result value correctly", func() {
				r := &types.ResultResponse{
					ValueA: 1,
					ValueB: 2,
					Result: 3,
				}
				result := r.SetResult(4)
				Expect(result).To(Equal(&types.ResultResponse{
					ValueA: 1,
					ValueB: 2,
					Result: 4,
				}))
			})
		})

		Context("SetValues", func() {
			It("should set the input values correctly", func() {
				r := &types.ResultResponse{
					ValueA: 1,
					ValueB: 2,
					Result: 3,
				}
				result := r.SetValues(4, 0)
				Expect(result).To(Equal(&types.ResultResponse{
					ValueA: 4,
					ValueB: 0,
					Result: 3,
				}))
			})
		})
	})

	Describe("Server API Endpoints", func() {
		Context("Addition Endpoint", func() {
			It("should correctly handle valid addition requests", func() {
				resp, err := call("/add", "1", "2")
				Expect(err).NotTo(HaveOccurred())
				Expect(resp.Result).To(Equal(float64(3)))
			})

			It("should handle invalid input gracefully", func() {
				_, err := call("/add", "bad", "2")
				Expect(err).To(HaveOccurred())

				_, err = call("/add", "1", "bad")
				Expect(err).To(HaveOccurred())

				_, err = call("/add", "1", "")
				Expect(err).To(HaveOccurred())
			})
		})

		Context("Subtraction Endpoint", func() {
			It("should correctly handle valid subtraction requests", func() {
				resp, err := call("/sub", "1", "2")
				Expect(err).NotTo(HaveOccurred())
				Expect(resp.Result).To(Equal(float64(-1)))
			})

			It("should handle invalid input gracefully", func() {
				_, err := call("/sub", "bad", "2")
				Expect(err).To(HaveOccurred())
			})
		})

		Context("Multiplication Endpoint", func() {
			It("should correctly handle valid multiplication requests", func() {
				resp, err := call("/mul", "1", "2")
				Expect(err).NotTo(HaveOccurred())
				Expect(resp.Result).To(Equal(float64(2)))
			})

			It("should handle invalid input gracefully", func() {
				_, err := call("/mul", "bad", "2")
				Expect(err).To(HaveOccurred())
			})
		})

		Context("Division Endpoint", func() {
			It("should correctly handle valid division requests", func() {
				resp, err := call("/div", "1", "2")
				Expect(err).NotTo(HaveOccurred())
				Expect(resp.Result).To(Equal(float64(0.5)))
			})

			It("should handle invalid input gracefully", func() {
				_, err := call("/div", "bad", "2")
				Expect(err).To(HaveOccurred())
			})
		})
	})
})
