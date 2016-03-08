package sgf

import (
	"io/ioutil"
	"log"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func readSGF(path string) string {
	s, e := ioutil.ReadFile(path)
	if e != nil {
		log.Fatal("can't read file: ", e.Error())
	}
	return string(s)
}

var _ = Describe("SGF", func() {
	var (
		sgf        = readSGF("./testdata/1.sgf")
		sgfTree, _ = SGFTreeParse(sgf)
	)

	Context("With sgf", func() {
		It("should be a 19x19", func() {
			Expect(sgfTree.Size()).To(Equal(19))
		})
		It("should have a title", func() {
			Expect(sgfTree.Title()).To(Equal("第25期竜星戦本戦Eブロック6回戦"))
		})
		It("should have a komi", func() {
			Expect(sgfTree.Komi()).To(Equal(6.5))
		})
		It("should have a result", func() {
			Expect(sgfTree.Result()).To(Equal("B+R"))
		})
		It("should have a datetime", func() {
			Expect(sgfTree.DateTime()).To(Equal("2015/11/09"))
		})
		It("should have a player black", func() {
			Expect(sgfTree.PlayerBlack()).To(Equal("坂井秀至"))
		})
		It("should have a player white", func() {
			Expect(sgfTree.PlayerWhite()).To(Equal("山田規喜"))
		})
		It("should have a place", func() {
			Expect(sgfTree.Place()).To(Equal("東京囲碁将棋チャンネルスタジオ"))
		})
	})
})
