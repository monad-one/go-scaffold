package spec

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	uut "github.com/orbs-network/go-experiment/services/statestorage"
	"github.com/orbs-network/go-experiment/types/services/statestorage"
)

var _ = Describe("Read Write", func() {

	var (
		service uut.Service
		stop chan error
	)

	BeforeEach(func() {
		stop = make(chan error, 10)
		service = uut.NewService()
		service.Start(&stop)
	})

	AfterEach(func() {
		service.Stop()
	})

	It("should read zero from unknown keys", func() {
		out, err := service.ReadKey(&statestorage.ReadKeyInput{Key: "unknown111"})
		Expect(err).ToNot(HaveOccurred())
		Expect(out.Value).To(BeEquivalentTo(0))
		out, err = service.ReadKey(&statestorage.ReadKeyInput{Key: "unknown222"})
		Expect(err).ToNot(HaveOccurred())
		Expect(out.Value).To(BeEquivalentTo(0))
	})

	It("should read the updated value after writing it", func() {
		_, err := service.WriteKey(&statestorage.WriteKeyInput{Key: "address111", Value: 2003})
		Expect(err).ToNot(HaveOccurred())
		out, err := service.ReadKey(&statestorage.ReadKeyInput{Key: "address111"})
		Expect(err).ToNot(HaveOccurred())
		Expect(out.Value).To(BeEquivalentTo(2003))
		_, err = service.WriteKey(&statestorage.WriteKeyInput{Key: "address111", Value: 867})
		Expect(err).ToNot(HaveOccurred())
		out, err = service.ReadKey(&statestorage.ReadKeyInput{Key: "address111"})
		Expect(err).ToNot(HaveOccurred())
		Expect(out.Value).To(BeEquivalentTo(867))
	})

})