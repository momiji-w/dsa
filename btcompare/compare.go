package btcompare

import "github.com/momiji-w/dsa/utils"

func Compare(source, target *utils.BinaryNode) bool {
	if source == nil && target == nil {
		return true
	}
	if source == nil || target == nil {
		return false
	}

	if source.Value != target.Value {
		return false
	}

	return Compare(source.Left, target.Left) && Compare(source.Right, target.Right)
}
