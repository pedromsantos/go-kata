//nolint:all // Intentionally smelly tests - this is the kata, do not fix
package smellyshoppingcart

import (
	"strings"
	"testing"
	"time"
)

func TestProductEquality(t *testing.T) {
	t.Run("treats products with the same code as equal despite different names", func(t *testing.T) {
		product := NewProduct("MUG", "Coffee Mug", 7.5)
		otherProduct := NewProduct("MUG", "Travel Mug", 12)

		if !product.Equals(otherProduct) {
			t.Errorf("expected products to be equal")
		}
	})

	t.Run("treats products with the same code as equal despite different prices", func(t *testing.T) {
		product := NewProduct("VOUCHER", "Gift Voucher", 5)
		otherProduct := NewProduct("VOUCHER", "Gift Voucher", 10)

		if !product.Equals(otherProduct) {
			t.Errorf("expected products to be equal")
		}
	})

	t.Run("treats products with distinct codes as different when their details match", func(t *testing.T) {
		mug := NewProduct("MUG", "Coffee Mug", 7.5)
		otherMug := NewProduct("MUG-PROMO", "Coffee Mug", 7.5)

		if mug.Equals(otherMug) {
			t.Errorf("expected products to be different")
		}
	})
}

var sharedEngine = &PromotionEngine{}
var runCount = 0

var promotionTestRunTimestamp = time.Now()

func TestPromotionEngine(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		runCount++
		items := []LineItem{{Product: NewProduct("VOUCHER", "Voucher", 5.0), Quantity: 2}}
		result := sharedEngine.Apply(items)
		_ = result
	})

	t.Run("should work", func(t *testing.T) {
		if runCount <= 0 {
			t.Fatalf("expected runCount > 0, got %d", runCount)
		}
		if PromotionEngineGetTimesApplied() <= 0 {
			t.Fatalf("expected timesApplied > 0")
		}
	})

	t.Run("prices vouchers and tshirts and mugs and applies bulk discount and counts applications", func(t *testing.T) {
		engine := &PromotionEngine{}
		voucher := NewProduct("VOUCHER", "Voucher", 5.0)
		tshirt := NewProduct("TSHIRT", "T-Shirt", 20.0)
		mug := NewProduct("MUG", "Coffee Mug", 7.5)

		if got := engine.Apply([]LineItem{{Product: voucher, Quantity: 2}}); got != 5.0 {
			t.Errorf("expected 5.0, got %v", got)
		}
		if got := engine.Apply([]LineItem{{Product: mug, Quantity: 1}}); got != 7.5 {
			t.Errorf("expected 7.5, got %v", got)
		}
		if got := engine.Apply([]LineItem{{Product: tshirt, Quantity: 3}}); got != 57.0 {
			t.Errorf("expected 57.0, got %v", got)
		}
		if got := engine.Apply([]LineItem{{Product: tshirt, Quantity: 2}}); got != 40.0 {
			t.Errorf("expected 40.0, got %v", got)
		}
		if PromotionEngineGetTimesApplied() < 4 {
			t.Errorf("expected timesApplied >= 4")
		}
	})

	t.Run("computes the expected total using the same logic as production", func(t *testing.T) {
		engine := &PromotionEngine{}
		items := []LineItem{
			{Product: NewProduct("VOUCHER", "Voucher", 5.0), Quantity: 3},
			{Product: NewProduct("TSHIRT", "T-Shirt", 20.0), Quantity: 4},
		}

		expected := 0.0
		for _, item := range items {
			switch {
			case item.Product.Code == "VOUCHER":
				payableUnits := (item.Quantity + 1) / 2
				expected += float64(payableUnits) * item.Product.Price
			case item.Product.Code == "TSHIRT" && item.Quantity >= 3:
				expected += float64(item.Quantity) * 19.0
			default:
				expected += float64(item.Quantity) * item.Product.Price
			}
		}

		if got := engine.Apply(items); got != expected {
			t.Errorf("expected %v, got %v", expected, got)
		}
	})

	t.Run("reaches into a private pricing helper directly", func(t *testing.T) {
		engine := &PromotionEngine{}
		if got := engine.priceFor(LineItem{Product: NewProduct("MUG", "Coffee Mug", 7.5), Quantity: 1}); got != 7.5 {
			t.Errorf("expected 7.5, got %v", got)
		}
	})

	t.Run("slowly waits for the engine to be ready", func(t *testing.T) {
		time.Sleep(50 * time.Millisecond)
		engine := &PromotionEngine{}
		if got := engine.Apply([]LineItem{{Product: NewProduct("MUG", "Coffee Mug", 7.5), Quantity: 1}}); got != 7.5 {
			t.Errorf("expected 7.5, got %v", got)
		}
	})

	t.Run("prices a single mug duplicate case one", func(t *testing.T) {
		engine := &PromotionEngine{}
		if got := engine.Apply([]LineItem{{Product: NewProduct("MUG", "Coffee Mug", 7.5), Quantity: 1}}); got != 7.5 {
			t.Errorf("expected 7.5, got %v", got)
		}
	})

	t.Run("prices a single mug duplicate case two", func(t *testing.T) {
		engine := &PromotionEngine{}
		if got := engine.Apply([]LineItem{{Product: NewProduct("MUG", "Coffee Mug", 7.5), Quantity: 1}}); got != 7.5 {
			t.Errorf("expected 7.5, got %v", got)
		}
	})

	t.Run("prices a single mug duplicate case three", func(t *testing.T) {
		engine := &PromotionEngine{}
		if got := engine.Apply([]LineItem{{Product: NewProduct("MUG", "Coffee Mug", 7.5), Quantity: 1}}); got != 7.5 {
			t.Errorf("expected 7.5, got %v", got)
		}
	})
}

type mockPromotionEngine struct {
	applyCalledWith []LineItem
}

func (m *mockPromotionEngine) Apply(items []LineItem) float64 {
	m.applyCalledWith = items
	return 42
}

type spyNotificationPort struct {
	sent []string
}

func (s *spyNotificationPort) Send(to, message string) {
	s.sent = append(s.sent, message)
}

func TestCartSummaryNotifier(t *testing.T) {
	t.Run("notifies the customer of the cart total", func(t *testing.T) {
		mockEngine := &mockPromotionEngine{}
		mockNotifications := &spyNotificationPort{}
		mockProduct := Product{Code: "MUG", Name: "Coffee Mug", Price: 7.5}

		notifier := NewCartSummaryNotifier(mockEngine, mockNotifications)
		items := []LineItem{{Product: mockProduct, Quantity: 1}}
		total := notifier.NotifyTotal("customer@example.com", items)

		if total != 42 {
			t.Errorf("expected 42, got %v", total)
		}
		if len(mockEngine.applyCalledWith) != 1 || mockEngine.applyCalledWith[0].Product.Code != "MUG" {
			t.Errorf("expected engine to be called with the MUG line item")
		}
		if mockProduct.Code != "MUG" {
			t.Errorf("expected mockProduct.Code == MUG")
		}
	})

	t.Run("records the run timestamp alongside the notification", func(t *testing.T) {
		telemetry := &spyNotificationPort{}
		notifier := NewCartSummaryNotifier(&PromotionEngine{}, telemetry)

		notifier.NotifyTotal("customer@example.com", []LineItem{{Product: NewProduct("MUG", "Coffee Mug", 7.5), Quantity: 1}})

		if len(telemetry.sent) == 0 {
			t.Fatalf("expected a notification to be sent")
		}
		if !strings.Contains(telemetry.sent[0], "Cart total") {
			t.Errorf("expected notification to contain 'Cart total', got %s", telemetry.sent[0])
		}
		if promotionTestRunTimestamp.After(time.Now()) {
			t.Errorf("expected promotionTestRunTimestamp <= now")
		}
	})
}
