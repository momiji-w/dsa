package btcompare

import (
	"github.com/momiji-w/dsa/utils"
	"testing"
)

func TestCompare(t *testing.T) {
	if !Compare(utils.Root, utils.Root) {
		t.Fatalf("Tree should be identical")
	}

	if Compare(utils.Root, utils.Root2) {
		t.Fatalf("Tree not should be identical")
	}
}
