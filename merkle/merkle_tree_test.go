package merkle_test

import (
	"github.com/muxxer/iota.go/consts"
	"github.com/muxxer/iota.go/kerl"
	. "github.com/muxxer/iota.go/merkle"
	"github.com/muxxer/iota.go/signing"
	"github.com/muxxer/iota.go/signing/key"
	"github.com/muxxer/iota.go/trinary"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const (
	seed          = "ABCDEFGHIJKLMNOPQRSTUVWXYZ9ABCDEFGHIJKLMNOPQRSTUVWXYZ9ABCDEFGHIJKLMNOPQRSTUVWXYZ9"
	securityLevel = 2
	depth         = 7
	expectedRoot  = "WFQ9CDKRMAKHKUEPPRGF9GWXCLLHRUGYNYDNWFQI9QDPLCQISLASULACYZLXGG9GGFGNRRSXEFSDHTBLW"
)

var _ = Describe("Merkle", func() {

	Context("CreateMerkleTree()", func() {
		// Using Depth 7
		merkleTree, err := CreateMerkleTree(seed, securityLevel, depth)

		It("does not throw", func() {
			Expect(err).To(BeNil())

		})

		It("creates a correctly-sized tree", func() {
			Expect(merkleTree.Layers[7].Level).To(Equal(7))
			Expect(merkleTree.Layers[3].Level).To(Equal(3))
			Expect(len(merkleTree.Layers)).To(Equal(8))
		})

		It("does not use Kerl KDF", func() {
			Expect(merkleTree.Root).NotTo(Equal("VERHESGRVSUWWZJNCKMQREASXZOIW9BBYGHV9QCLVCIGJYZOEIODSIHRCBZAFNNAJSTSC9LRHKKBLJPDB"))
		})

		It("leaves are computed using Shake KDF", func() {
			leavesCount := 1 << uint(depth)
			leaves := merkleTree.Layers[depth].Hashes
			for index := 0; index < leavesCount; index++ {
				subSeedTrits, _ := signing.Subseed(seed, uint64(index), kerl.NewKerl())
				keyTrits, _ := key.Shake(subSeedTrits, consts.SecurityLevel(securityLevel))
				digestsTrits, _ := signing.Digests(keyTrits, kerl.NewKerl())
				addressTrits, _ := signing.Address(digestsTrits, kerl.NewKerl())
				address, _ := trinary.TritsToTrytes(addressTrits)
				Expect(leaves[index]).To(Equal(address))
			}
		})

		It("each node is the hash of the corresponding two children using Kerl sponge", func() {
			layers := merkleTree.Layers
			for d := 1; d <= depth; d++ {
				for pair := 0; pair < 1<<uint(d); pair += 2 {
					sponge := kerl.NewKerl()
					sponge.MustAbsorbTrytes(layers[d].Hashes[pair])
					sponge.MustAbsorbTrytes(layers[d].Hashes[pair+1])
					Expect(layers[d-1].Hashes[pair/2]).To(Equal(sponge.MustSqueezeTrytes(consts.HashTrinarySize)))
				}
			}
		})

		It("match root", func() {
			Expect(merkleTree.Layers[0].Hashes[0]).To(Equal(merkleTree.Root))
			Expect(merkleTree.Root).To(Equal(expectedRoot))
		})

	})
})
