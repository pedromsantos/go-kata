package shoppingcart

// Cart - implement this following TDD
type Cart struct{}

// CartRepository interface for persistence
type CartRepository interface {
	Save(cart *Cart) error
	FindByID(id string) (*Cart, error)
}
