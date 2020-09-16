package merkle_test

import (
	"github.com/muxxer/iota.go/consts"
	. "github.com/muxxer/iota.go/merkle"
	"github.com/muxxer/iota.go/trinary"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var (
	hashToSign = consts.NullHashTrytes
)

var _ = Describe("Signing", func() {
	var tree *MerkleTree

	BeforeSuite(func() {
		var err error
		tree, err = CreateMerkleTree(seed, securityLevel, depth, MerkleCreateOptions{Parallelism: 1})
		Expect(err).ToNot(HaveOccurred())
	})

	DescribeTable("sign and validate",
		func(leafIndex uint32) {
			path, err := tree.AuditPath(leafIndex)
			Expect(err).ToNot(HaveOccurred())
			fragments, err := SignatureFragments(seed, leafIndex, securityLevel, hashToSign)
			Expect(err).ToNot(HaveOccurred())

			valid, err := ValidateSignatureFragments(tree.Root, leafIndex, path, fragments, hashToSign)
			Expect(err).ToNot(HaveOccurred())
			Expect(valid).To(BeTrue())
		},
		Entry("leafIndex: 0", uint32(0)),
		Entry("leafIndex: 1", uint32(1)),
		Entry("leafIndex: 2", uint32(2)),
		Entry("max leafIndex", uint32(1<<depth-1)),
	)

	Context("MerkleRoot()", func() {

		It("valid audit path", func() {
			path := make([]trinary.Trytes, 32)
			for i := range path {
				path[i] = consts.NullHashTrytes
			}
			root, err := MerkleRoot(consts.NullHashTrytes, 1<<32-1, path)
			Expect(err).ToNot(HaveOccurred())
			Expect(root).To(Equal("MDKGSWENCCKHKNSHEZUX9LCCDKDJJR9BXLXXKRVMUGBLOVESSLRKWOPOE9UUZZOTOIOVMTCKQLTDQITPD"))
		})

		It("audit path too short", func() {
			path := make([]trinary.Trytes, depth)
			_, err := MerkleRoot(tree.Root, 1<<depth, path)
			Expect(err).To(Equal(ErrInvalidAuditPathLength))
		})

		It("audit path invalid tryte lengths", func() {
			path := []trinary.Trytes{""}
			_, err := MerkleRoot(tree.Root, 0, path)
			Expect(err).To(HaveOccurred())
		})
	})

})
