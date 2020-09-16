package integration_test

import (
	. "github.com/muxxer/iota.go/api"
	. "github.com/muxxer/iota.go/api/integration/samples"
	. "github.com/muxxer/iota.go/consts"
	. "github.com/muxxer/iota.go/trinary"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pkg/errors"
)

var _ = Describe("StoreAndBroadcast()", func() {
	api, err := ComposeAPI(HTTPClientSettings{}, nil)
	if err != nil {
		panic(err)
	}

	Context("call", func() {
		It("resolves to correct response", func() {
			bundleTrytes, err := api.StoreAndBroadcast(BundleTrytes)
			Expect(err).ToNot(HaveOccurred())
			Expect(bundleTrytes).To(Equal(BundleTrytes))
		})
	})

	Context("invalid input", func() {
		It("returns an error for invalid trytes", func() {
			_, err := api.StoreAndBroadcast([]Trytes{"balalaika"})
			Expect(errors.Cause(err)).To(Equal(ErrInvalidAttachedTrytes))
		})
	})

})
