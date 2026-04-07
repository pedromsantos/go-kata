package gildedrose

import (
	"fmt"
	"testing"

	approvals "github.com/approvals/go-approval-tests"
	"github.com/stretchr/testify/assert"
	"pgregory.net/rapid"
)

// =============================================================================
// 1. CHARACTERIZATION TESTS - Document current behavior
// =============================================================================

// 1.1 Unit Tests - Start here!
func TestGildedRose(t *testing.T) {
	// Example characterization test - uncomment and modify:
	// t.Run("normal item decreases quality by 1", func(t *testing.T) {
	// 	items := []*Item{{Name: "foo", SellIn: 10, Quality: 20}}
	// 	UpdateQuality(items)
	// 	assert.Equal(t, 19, items[0].Quality)
	// })
}

// 1.2 Parameterized Tests (Table-Driven) - Expand coverage
func TestGildedRoseTableDriven(t *testing.T) {
	// Uncomment to use table-driven tests:
	// testCases := []struct {
	// 	name            string
	// 	inputName       string
	// 	inputSellIn     int
	// 	inputQuality    int
	// 	expectedSellIn  int
	// 	expectedQuality int
	// }{
	// 	{"normal item", "foo", 10, 20, 9, 19},
	// 	{"normal item at zero quality", "foo", 10, 0, 9, 0},
	// 	{"aged brie increases", "Aged Brie", 10, 20, 9, 21},
	// }
	//
	// for _, tc := range testCases {
	// 	t.Run(tc.name, func(t *testing.T) {
	// 		items := []*Item{{Name: tc.inputName, SellIn: tc.inputSellIn, Quality: tc.inputQuality}}
	// 		UpdateQuality(items)
	// 		assert.Equal(t, tc.expectedSellIn, items[0].SellIn)
	// 		assert.Equal(t, tc.expectedQuality, items[0].Quality)
	// 	})
	// }
}

// =============================================================================
// 2. GOLDEN MASTER / APPROVAL TESTS - Lock in current behavior
// =============================================================================

func TestGildedRoseApproval(t *testing.T) {
	// Uncomment to use approval tests:
	// Run once to create .approved.txt, then changes will fail until approved
	//
	// var result strings.Builder
	// names := []string{"foo", "Aged Brie", "Backstage passes to a TAFKAL80ETC concert", "Sulfuras, Hand of Ragnaros"}
	// sellIns := []int{-1, 0, 1, 11}
	// qualities := []int{0, 1, 2, 49, 50}
	//
	// for _, name := range names {
	// 	for _, sellIn := range sellIns {
	// 		for _, quality := range qualities {
	// 			items := []*Item{{Name: name, SellIn: sellIn, Quality: quality}}
	// 			UpdateQuality(items)
	// 			result.WriteString(fmt.Sprintf("%s, %d, %d => %d, %d\n",
	// 				name, sellIn, quality, items[0].SellIn, items[0].Quality))
	// 		}
	// 	}
	// }
	//
	// approvals.VerifyString(t, result.String())
	_ = approvals.VerifyString // silence unused import
	_ = fmt.Sprintf            // silence unused import
}

// =============================================================================
// 3. PROPERTY-BASED TESTS - Capture invariants
// =============================================================================

func TestGildedRoseProperties(t *testing.T) {
	// Uncomment to use property-based tests:

	// t.Run("quality is never negative", func(t *testing.T) {
	// 	rapid.Check(t, func(t *rapid.T) {
	// 		name := rapid.StringMatching(`[a-zA-Z ]+`).Draw(t, "name")
	// 		sellIn := rapid.IntRange(-10, 50).Draw(t, "sellIn")
	// 		quality := rapid.IntRange(0, 100).Draw(t, "quality")
	//
	// 		items := []*Item{{Name: name, SellIn: sellIn, Quality: quality}}
	// 		UpdateQuality(items)
	//
	// 		assert.GreaterOrEqual(t, items[0].Quality, 0, "quality should never be negative")
	// 	})
	// })

	// t.Run("quality never exceeds 50 except Sulfuras", func(t *testing.T) {
	// 	rapid.Check(t, func(t *rapid.T) {
	// 		name := rapid.SampledFrom([]string{"foo", "Aged Brie", "Backstage passes to a TAFKAL80ETC concert"}).Draw(t, "name")
	// 		sellIn := rapid.IntRange(-10, 50).Draw(t, "sellIn")
	// 		quality := rapid.IntRange(0, 50).Draw(t, "quality")
	//
	// 		items := []*Item{{Name: name, SellIn: sellIn, Quality: quality}}
	// 		UpdateQuality(items)
	//
	// 		assert.LessOrEqual(t, items[0].Quality, 50, "quality should never exceed 50")
	// 	})
	// })

	// t.Run("Sulfuras never changes", func(t *testing.T) {
	// 	rapid.Check(t, func(t *rapid.T) {
	// 		sellIn := rapid.IntRange(-10, 50).Draw(t, "sellIn")
	// 		quality := rapid.IntRange(0, 80).Draw(t, "quality")
	//
	// 		items := []*Item{{Name: "Sulfuras, Hand of Ragnaros", SellIn: sellIn, Quality: quality}}
	// 		originalSellIn := items[0].SellIn
	// 		originalQuality := items[0].Quality
	// 		UpdateQuality(items)
	//
	// 		assert.Equal(t, originalSellIn, items[0].SellIn, "Sulfuras sellIn should not change")
	// 		assert.Equal(t, originalQuality, items[0].Quality, "Sulfuras quality should not change")
	// 	})
	// })

	_ = rapid.Check  // silence unused import
	_ = assert.Equal // silence unused import
}

// =============================================================================
// Helper function matching TypeScript version
// =============================================================================

// doUpdateQuality is a helper for testing - uncomment tests to use it
func doUpdateQuality(name string, sellIn int, quality int) *Item { //nolint:unused
	items := []*Item{{Name: name, SellIn: sellIn, Quality: quality}}
	gr := NewGildedRose(items)
	gr.UpdateQuality()
	return items[0]
}

// UpdateQuality is a convenience wrapper
func UpdateQuality(items []*Item) {
	gr := NewGildedRose(items)
	gr.UpdateQuality()
}
