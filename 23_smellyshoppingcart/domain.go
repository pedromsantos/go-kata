//nolint:all // This is intentionally legacy/smelly code for characterization-testing practice - do not fix
package smellyshoppingcart

import "fmt"

// Product is a catalogue item.
type Product struct {
	Code  string
	Name  string
	Price float64
}

func NewProduct(code, name string, price float64) Product {
	return Product{Code: code, Name: name, Price: price}
}

func (p Product) Equals(other Product) bool {
	return p.Code == other.Code
}

// LineItem pairs a product with a quantity in a cart.
type LineItem struct {
	Product  Product
	Quantity int
}

// Cart is the shopping cart aggregate.
type Cart struct {
	ID              string
	CustomerName    string
	items           []LineItem
	promotionEngine *PromotionEngine
}

func NewCart(id, customerName string) *Cart {
	return &Cart{
		ID:              id,
		CustomerName:    customerName,
		promotionEngine: &PromotionEngine{},
	}
}

func (c *Cart) AddProduct(product Product, quantity int) {
	for i := range c.items {
		if c.items[i].Product.Equals(product) {
			c.items[i].Quantity += quantity
			return
		}
	}
	c.items = append(c.items, LineItem{Product: product, Quantity: quantity})
}

func (c *Cart) LineItems() []LineItem {
	return c.items
}

func (c *Cart) CalculateSubtotal() float64 {
	return c.promotionEngine.Apply(c.items)
}

// PromotionApplier is the interface CartSummaryNotifier depends on.
type PromotionApplier interface {
	Apply(items []LineItem) float64
}

var promotionEngineTimesApplied = 0

// PromotionEngine prices a cart's line items, applying promotional rules.
type PromotionEngine struct{}

func (e *PromotionEngine) Apply(items []LineItem) float64 {
	promotionEngineTimesApplied++

	total := 0.0
	for _, item := range items {
		total += e.priceFor(item)
	}
	return total
}

func PromotionEngineGetTimesApplied() int {
	return promotionEngineTimesApplied
}

func (e *PromotionEngine) priceFor(item LineItem) float64 {
	twoForOneCodes := []string{"VOUCHER"}
	bulkDiscountCode := "TSHIRT"
	bulkDiscountThreshold := 3
	bulkDiscountPrice := 19.0

	for _, code := range twoForOneCodes {
		if item.Product.Code == code {
			payableUnits := (item.Quantity + 1) / 2
			return float64(payableUnits) * item.Product.Price
		}
	}

	if item.Product.Code == bulkDiscountCode && item.Quantity >= bulkDiscountThreshold {
		return float64(item.Quantity) * bulkDiscountPrice
	}

	return float64(item.Quantity) * item.Product.Price
}

// Clock is the port CheckoutCart uses to obtain the confirmation timestamp.
type Clock interface {
	Now() string
}

// NotificationPort is the port used to notify a customer.
type NotificationPort interface {
	Send(to, message string)
}

// ShoppingCartRepository is the port used to persist and load carts.
type ShoppingCartRepository interface {
	Save(cart *Cart) error
	FindByID(id string) (*Cart, error)
}

// CartSummaryNotifier prices a cart and notifies the customer of the total.
type CartSummaryNotifier struct {
	promotionEngine PromotionApplier
	notifications   NotificationPort
}

func NewCartSummaryNotifier(promotionEngine PromotionApplier, notifications NotificationPort) *CartSummaryNotifier {
	return &CartSummaryNotifier{promotionEngine: promotionEngine, notifications: notifications}
}

func (n *CartSummaryNotifier) NotifyTotal(customerEmail string, items []LineItem) float64 {
	total := n.promotionEngine.Apply(items)
	n.notifications.Send(customerEmail, fmt.Sprintf("Cart total: %.2f€", total))
	return total
}
