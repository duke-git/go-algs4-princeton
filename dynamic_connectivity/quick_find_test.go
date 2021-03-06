package dynamicconnectivity

import (
	"go-algs4-princeton/internal"
	"testing"
)

func TestUnion_Connected(t *testing.T) {
	assert := internal.NewAssert(t, "TestUnion_Connected")
	qf := NewQuickFind(10)
	qf.Union(3, 4)
	qf.Union(7, 8)

	qf.PrintData()

	assert.Equal(true, qf.Connected(3, 4))
	assert.Equal(true, qf.Connected(7, 8))
	assert.Equal(false, qf.Connected(3, 7))
}

func TestQuickUnion_Connected(t *testing.T) {
	assert := internal.NewAssert(t, "TestQuickUnion_Connected")
	qf := NewQuickFind(10)
	qf.QuickUnion(3, 4)
	qf.QuickUnion(7, 8)

	qf.PrintData()

	assert.Equal(true, qf.QuickConnected(3, 4))
	assert.Equal(true, qf.QuickConnected(7, 8))
	assert.Equal(false, qf.QuickConnected(3, 7))
}

func TestWeightedQuickUnion(t *testing.T) {
	assert := internal.NewAssert(t, "TestWeightedQuickUnion")
	qf := NewQuickFind(10)
	qf.WeightedQuickUnion(3, 4)
	qf.WeightedQuickUnion(7, 8)

	qf.PrintData()

	assert.Equal(true, qf.QuickConnected(3, 4))
	assert.Equal(true, qf.QuickConnected(7, 8))
	assert.Equal(false, qf.QuickConnected(3, 7))
}

func Test_WeightedQuickUnion_WithPathCompression(t *testing.T) {
	assert := internal.NewAssert(t, "Test_WeightedQuickUnion_WithPathCompression")
	qf := NewQuickFind(10)
	qf.WeightedQuickUnionWithPathCompression(3, 4)
	qf.WeightedQuickUnionWithPathCompression(7, 8)

	qf.PrintData()

	assert.Equal(true, qf.QuickConnected(3, 4))
	assert.Equal(true, qf.QuickConnected(7, 8))
	assert.Equal(false, qf.QuickConnected(3, 7))
}
