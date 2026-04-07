//nolint:all // This is intentionally legacy code for refactoring practice - do not fix linting issues
package gildedrose

// Item represents an item in the Gilded Rose inventory.
type Item struct {
	Name    string
	SellIn  int
	Quality int
}

// GildedRose manages the inventory.
type GildedRose struct {
	Items []*Item
}

// NewGildedRose creates a new GildedRose instance.
func NewGildedRose(items []*Item) *GildedRose {
	return &GildedRose{Items: items}
}

// UpdateQuality updates the quality of all items.
// This is the legacy code to be tested and refactored.
func (g *GildedRose) UpdateQuality() []*Item {
	for i := 0; i < len(g.Items); i++ {
		if g.Items[i].Name != "Aged Brie" && g.Items[i].Name != "Backstage passes to a TAFKAL80ETC concert" {
			if g.Items[i].Quality > 0 {
				if g.Items[i].Name != "Sulfuras, Hand of Ragnaros" {
					g.Items[i].Quality = g.Items[i].Quality - 1
				}
			}
		} else {
			if g.Items[i].Quality < 50 {
				g.Items[i].Quality = g.Items[i].Quality + 1
				if g.Items[i].Name == "Backstage passes to a TAFKAL80ETC concert" {
					if g.Items[i].SellIn < 11 {
						if g.Items[i].Quality < 50 {
							g.Items[i].Quality = g.Items[i].Quality + 1
						}
					}
					if g.Items[i].SellIn < 6 {
						if g.Items[i].Quality < 50 {
							g.Items[i].Quality = g.Items[i].Quality + 1
						}
					}
				}
			}
		}
		if g.Items[i].Name != "Sulfuras, Hand of Ragnaros" {
			g.Items[i].SellIn = g.Items[i].SellIn - 1
		}
		if g.Items[i].SellIn < 0 {
			if g.Items[i].Name != "Aged Brie" {
				if g.Items[i].Name != "Backstage passes to a TAFKAL80ETC concert" {
					if g.Items[i].Quality > 0 {
						if g.Items[i].Name != "Sulfuras, Hand of Ragnaros" {
							g.Items[i].Quality = g.Items[i].Quality - 1
						}
					}
				} else {
					g.Items[i].Quality = g.Items[i].Quality - g.Items[i].Quality
				}
			} else {
				if g.Items[i].Quality < 50 {
					g.Items[i].Quality = g.Items[i].Quality + 1
				}
			}
		}
	}
	return g.Items
}
