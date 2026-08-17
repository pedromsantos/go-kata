//nolint:all // Intentionally smelly tests - this is the kata, do not fix
package smellyshoppingcart

import (
	"io"
	"os"
	"strings"
	"testing"
	"time"
)

const defaultCustomerEmail = "customer@example.com"

func TestInMemoryShoppingCartRepositorySmells(t *testing.T) {
	t.Run("test2", func(t *testing.T) {
		repository := &InMemoryShoppingCartRepository{}
		cart := NewCart("cart-1", "Ada Lovelace")
		cart.AddProduct(NewProduct("MUG", "Coffee Mug", 7.5), 1)

		if err := repository.Save(cart); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	t.Run("finds the cart saved earlier", func(t *testing.T) {
		repository := &InMemoryShoppingCartRepository{}
		found, err := repository.FindByID("cart-1")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if found == nil {
			t.Errorf("expected to find the cart saved by the previous test")
		}
	})

	t.Run("saves and re-finds and mutates and re-saves and counts items and checks the customer name", func(t *testing.T) {
		repository := &InMemoryShoppingCartRepository{}
		cart := NewCart("cart-2", "Grace Hopper")
		cart.AddProduct(NewProduct("VOUCHER", "Voucher", 5.0), 1)
		if err := repository.Save(cart); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		firstFind, err := repository.FindByID("cart-2")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		firstFind.AddProduct(NewProduct("TSHIRT", "T-Shirt", 20.0), 1)
		if err := repository.Save(firstFind); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		secondFind, err := repository.FindByID("cart-2")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if secondFind == nil {
			t.Fatalf("expected to find cart-2")
		}
		if secondFind.ID != "cart-2" {
			t.Errorf("expected id cart-2, got %s", secondFind.ID)
		}
		if secondFind.CustomerName != "Grace Hopper" {
			t.Errorf("expected customer name Grace Hopper, got %s", secondFind.CustomerName)
		}
		if len(secondFind.LineItems()) != 2 {
			t.Errorf("expected 2 line items, got %d", len(secondFind.LineItems()))
		}
		notFound, err := repository.FindByID("does-not-exist")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if notFound != nil {
			t.Errorf("expected nil for an unknown cart id")
		}
	})

	t.Run("slowly waits for the in-memory store to be ready", func(t *testing.T) {
		time.Sleep(50 * time.Millisecond)
		repository := &InMemoryShoppingCartRepository{}
		cart := NewCart("cart-3", "Margaret Hamilton")
		if err := repository.Save(cart); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		found, err := repository.FindByID("cart-3")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if found == nil {
			t.Errorf("expected to find cart-3")
		}
	})

	t.Run("saves a cart built by bypassing the constructor instead of the real construction path", func(t *testing.T) {
		doubledCart := &Cart{ID: "cart-4", CustomerName: "Katherine Johnson"}
		repository := &InMemoryShoppingCartRepository{}

		if err := repository.Save(doubledCart); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		found, err := repository.FindByID("cart-4")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if found != doubledCart {
			t.Errorf("expected the same instance back")
		}
	})
}

func TestEmailNotificationGateway(t *testing.T) {
	t.Run("sends an order confirmation email", func(t *testing.T) {
		originalStdout := os.Stdout
		r, w, err := os.Pipe()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		os.Stdout = w

		gateway := NewEmailNotificationGateway()
		gateway.Send(defaultCustomerEmail, "Order confirmed: ORD-1")

		w.Close()
		os.Stdout = originalStdout
		output, err := io.ReadAll(r)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if !strings.Contains(string(output), defaultCustomerEmail) {
			t.Errorf("expected output to contain %s, got %s", defaultCustomerEmail, output)
		}
	})
}

func aCartWithProducts(id string) *Cart {
	cart := NewCart(id, "Ada Lovelace")
	cart.AddProduct(NewProduct("MUG", "Coffee Mug", 7.5), 2)
	cart.AddProduct(NewProduct("VOUCHER", "Gift Voucher", 5), 1)
	return cart
}

func TestInMemoryShoppingCartRepositoryIntegration(t *testing.T) {
	setup := func() *InMemoryShoppingCartRepository {
		ClearInMemoryShoppingCartRepository()
		return &InMemoryShoppingCartRepository{}
	}

	t.Run("FindsCart_WhenSavedThroughRepository", func(t *testing.T) {
		repository := setup()
		defer ClearInMemoryShoppingCartRepository()
		cart := aCartWithProducts("repository-integration-cart-1")

		if err := repository.Save(cart); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		found, err := repository.FindByID(cart.ID)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if found != cart {
			t.Errorf("expected found cart to be the same instance as the saved cart")
		}
	})

	t.Run("ReturnsNull_WhenCartIdIsUnknown", func(t *testing.T) {
		repository := setup()
		defer ClearInMemoryShoppingCartRepository()

		found, err := repository.FindByID("unknown-cart")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if found != nil {
			t.Errorf("expected nil, got %v", found)
		}
	})
}
