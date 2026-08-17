//nolint:all // This is intentionally legacy/smelly code for characterization-testing practice - do not fix
package smellyshoppingcart

import "fmt"

// EmailNotificationGateway sends notifications by "emailing" (printing) them.
type EmailNotificationGateway struct {
	fromAddress string
}

func NewEmailNotificationGateway() *EmailNotificationGateway {
	return &EmailNotificationGateway{fromAddress: "orders@shop.example.com"}
}

func (g *EmailNotificationGateway) Send(to, message string) {
	fmt.Printf("[EMAIL %s -> %s] %s\n", g.fromAddress, to, message)
}

var inMemoryCarts = map[string]*Cart{}

// InMemoryShoppingCartRepository stores carts in a package-level map shared
// across every instance.
type InMemoryShoppingCartRepository struct{}

func (r *InMemoryShoppingCartRepository) Save(cart *Cart) error {
	inMemoryCarts[cart.ID] = cart
	return nil
}

func (r *InMemoryShoppingCartRepository) FindByID(id string) (*Cart, error) {
	cart, ok := inMemoryCarts[id]
	if !ok {
		return nil, nil
	}
	return cart, nil
}

func ClearInMemoryShoppingCartRepository() {
	inMemoryCarts = map[string]*Cart{}
}
