//nolint:all // This is intentionally legacy/smelly code for characterization-testing practice - do not fix
package smellyshoppingcart

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

var productCatalog = map[string]Product{
	"VOUCHER": NewProduct("VOUCHER", "Voucher", 5.0),
	"TSHIRT":  NewProduct("TSHIRT", "T-Shirt", 20.0),
	"MUG":     NewProduct("MUG", "Coffee Mug", 7.5),
}

func findProduct(code string) (Product, error) {
	product, ok := productCatalog[code]
	if !ok {
		return Product{}, fmt.Errorf("unknown product code %s", code)
	}
	return product, nil
}

// AddProductToCart adds a catalogue product to an existing cart.
type AddProductToCart struct {
	repository ShoppingCartRepository
}

func NewAddProductToCart(repository ShoppingCartRepository) *AddProductToCart {
	return &AddProductToCart{repository: repository}
}

func (uc *AddProductToCart) Execute(cartID, productCode string, quantity int) error {
	cart, err := uc.repository.FindByID(cartID)
	if err != nil {
		return err
	}
	if cart == nil {
		return fmt.Errorf("cart %s not found", cartID)
	}

	product, err := findProduct(productCode)
	if err != nil {
		return err
	}

	cart.AddProduct(product, quantity)
	return uc.repository.Save(cart)
}

func orderClockNow() string {
	return time.Now().Format(time.RFC3339)
}

// OrderClock is the default production Clock.
type OrderClock struct{}

func (c *OrderClock) Now() string {
	return orderClockNow()
}

// Receipt is the outcome of a successful checkout.
type Receipt struct {
	CartID           string
	Total            float64
	ConfirmationCode string
	ConfirmedAt      string
}

// CheckoutCart prices a cart, confirms the order, and notifies the customer.
type CheckoutCart struct {
	repository   ShoppingCartRepository
	notifier     NotificationPort
	clock        Clock
	randomSource func() float64
}

func NewCheckoutCart(repository ShoppingCartRepository, notifier NotificationPort, clock Clock, randomSource func() float64) *CheckoutCart {
	return &CheckoutCart{repository: repository, notifier: notifier, clock: clock, randomSource: randomSource}
}

// NewDefaultCheckoutCart wires production defaults, mirroring the TS
// constructor's default parameter values.
func NewDefaultCheckoutCart(repository ShoppingCartRepository) *CheckoutCart {
	return NewCheckoutCart(repository, NewEmailNotificationGateway(), &OrderClock{}, rand.Float64)
}

func (uc *CheckoutCart) Execute(cartID, customerEmail string) (Receipt, error) {
	cart, err := uc.repository.FindByID(cartID)
	if err != nil {
		return Receipt{}, err
	}
	if cart == nil {
		return Receipt{}, fmt.Errorf("cart %s not found", cartID)
	}

	total := cart.CalculateSubtotal()
	confirmationCode := fmt.Sprintf("ORD-%d", int(math.Floor(uc.randomSource()*1_000_000)))
	confirmedAt := uc.clock.Now()

	uc.notifier.Send(customerEmail, fmt.Sprintf("Order confirmed: %s, total %.2f€", confirmationCode, total))

	return Receipt{CartID: cartID, Total: total, ConfirmationCode: confirmationCode, ConfirmedAt: confirmedAt}, nil
}
