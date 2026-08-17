package smellyshoppingcart

import "testing"

// fakeShoppingCartRepository is an in-test fake for ShoppingCartRepository,
// scoped to this test file only (not the shared InMemoryShoppingCartRepository).
type fakeShoppingCartRepository struct {
	carts map[string]*Cart
}

func newFakeShoppingCartRepository() *fakeShoppingCartRepository {
	return &fakeShoppingCartRepository{carts: map[string]*Cart{}}
}

func (r *fakeShoppingCartRepository) seed(cart *Cart) {
	r.carts[cart.ID] = cart
}

func (r *fakeShoppingCartRepository) Save(cart *Cart) error {
	r.carts[cart.ID] = cart
	return nil
}

func (r *fakeShoppingCartRepository) FindByID(id string) (*Cart, error) {
	cart, ok := r.carts[id]
	if !ok {
		return nil, nil
	}
	return cart, nil
}

// fixedClock is a Clock double that always returns the same instant.
type fixedClock struct {
	instant string
}

func (c fixedClock) Now() string {
	return c.instant
}

// spyNotifier records every notification sent through it.
type spyNotifier struct {
	sentTo      []string
	sentMessage []string
}

func (n *spyNotifier) Send(to, message string) {
	n.sentTo = append(n.sentTo, to)
	n.sentMessage = append(n.sentMessage, message)
}

func (n *spyNotifier) calledTimes() int {
	return len(n.sentTo)
}

var mug = NewProduct("MUG", "Coffee Mug", 7.5)
var voucher = NewProduct("VOUCHER", "Voucher", 5.0)

const fixedConfirmedAt = "2024-01-01T00:00:00.000Z"

// fixedRandomSource returns 0.5 -> math.Floor(0.5 * 1_000_000) = 500000
func fixedRandomSource() float64 {
	return 0.5
}

func newCheckoutCartAcceptanceFixture() (*fakeShoppingCartRepository, *spyNotifier, *CheckoutCart) {
	repository := newFakeShoppingCartRepository()
	notifier := &spyNotifier{}
	useCase := NewCheckoutCart(repository, notifier, fixedClock{instant: fixedConfirmedAt}, fixedRandomSource)
	return repository, notifier, useCase
}

func TestCheckoutCartAcceptance(t *testing.T) {
	t.Run("confirms checkout and returns a receipt when the cart has no discounts", func(t *testing.T) {
		repository, _, useCase := newCheckoutCartAcceptanceFixture()
		cart := NewCart("cart-1", "Ada Lovelace")
		cart.AddProduct(mug, 1)
		repository.seed(cart)

		receipt, err := useCase.Execute("cart-1", "ada@example.com")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if receipt.CartID != "cart-1" {
			t.Errorf("expected cart-1, got %s", receipt.CartID)
		}
		if receipt.Total != 7.5 {
			t.Errorf("expected 7.5, got %v", receipt.Total)
		}
		if receipt.ConfirmationCode != "ORD-500000" {
			t.Errorf("expected ORD-500000, got %s", receipt.ConfirmationCode)
		}
		if receipt.ConfirmedAt != fixedConfirmedAt {
			t.Errorf("expected %s, got %s", fixedConfirmedAt, receipt.ConfirmedAt)
		}
	})

	t.Run("notifies the customer of the confirmed total when checkout succeeds", func(t *testing.T) {
		repository, notifier, useCase := newCheckoutCartAcceptanceFixture()
		cart := NewCart("cart-2", "Ada Lovelace")
		cart.AddProduct(mug, 1)
		repository.seed(cart)

		receipt, err := useCase.Execute("cart-2", "ada@example.com")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if notifier.calledTimes() != 1 {
			t.Fatalf("expected 1 notification, got %d", notifier.calledTimes())
		}
		expectedMessage := "Order confirmed: " + receipt.ConfirmationCode + ", total 7.50€"
		if notifier.sentTo[0] != "ada@example.com" || notifier.sentMessage[0] != expectedMessage {
			t.Errorf("expected notification to ada@example.com with %q, got %q to %q", expectedMessage, notifier.sentMessage[0], notifier.sentTo[0])
		}
	})

	t.Run("computes the confirmed total using real promotion rules when the cart qualifies for a two-for-one discount", func(t *testing.T) {
		repository, notifier, useCase := newCheckoutCartAcceptanceFixture()
		cart := NewCart("cart-3", "Grace Hopper")
		cart.AddProduct(voucher, 3) // two-for-one: 2 payable units * 5.0€ = 10.0€
		repository.seed(cart)

		receipt, err := useCase.Execute("cart-3", "grace@example.com")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if receipt.Total != 10.0 {
			t.Errorf("expected 10.0, got %v", receipt.Total)
		}
		expectedMessage := "Order confirmed: " + receipt.ConfirmationCode + ", total 10.00€"
		if notifier.sentMessage[len(notifier.sentMessage)-1] != expectedMessage {
			t.Errorf("expected %q, got %q", expectedMessage, notifier.sentMessage[len(notifier.sentMessage)-1])
		}
	})

	t.Run("rejects checkout when the cart does not exist", func(t *testing.T) {
		_, _, useCase := newCheckoutCartAcceptanceFixture()

		_, err := useCase.Execute("missing-cart", "nobody@example.com")

		if err == nil {
			t.Fatalf("expected an error")
		}
		if err.Error() != "cart missing-cart not found" {
			t.Errorf("expected 'cart missing-cart not found', got %v", err)
		}
	})

	t.Run("does not notify the customer when the cart does not exist", func(t *testing.T) {
		_, notifier, useCase := newCheckoutCartAcceptanceFixture()

		if _, err := useCase.Execute("missing-cart", "nobody@example.com"); err == nil {
			t.Fatalf("expected an error")
		}

		if notifier.calledTimes() != 0 {
			t.Errorf("expected no notifications, got %d", notifier.calledTimes())
		}
	})
}
