package smellyshoppingcart

import "testing"

const cartMotherDefaultCartID = "cart-1"
const cartMotherDefaultCustomerName = "Ada Lovelace"

type cartMother struct{}

var CartMother = cartMother{}

func (cartMother) create() *Cart {
	return NewCart(cartMotherDefaultCartID, cartMotherDefaultCustomerName)
}

func (m cartMother) EmptyCart() *Cart {
	return m.create()
}

func (m cartMother) VoucherCart(quantity int) *Cart {
	cart := m.create()
	cart.AddProduct(NewProduct("VOUCHER", "Voucher", 5), quantity)
	return cart
}

func (m cartMother) TShirtCart(quantity int) *Cart {
	cart := m.create()
	cart.AddProduct(NewProduct("TSHIRT", "T-Shirt", 20), quantity)
	return cart
}

func TestCartMother(t *testing.T) {
	t.Run("creates a valid cart with stable defaults", func(t *testing.T) {
		cart := CartMother.create()

		if cart.ID != "cart-1" {
			t.Errorf("expected id cart-1, got %s", cart.ID)
		}
		if cart.CustomerName != "Ada Lovelace" {
			t.Errorf("expected customer name Ada Lovelace, got %s", cart.CustomerName)
		}
		if len(cart.LineItems()) != 0 {
			t.Errorf("expected 0 line items, got %d", len(cart.LineItems()))
		}
	})

	t.Run("uses named scenarios to create valid carts with controlled quantities", func(t *testing.T) {
		emptyCart := CartMother.EmptyCart()
		voucherCart := CartMother.VoucherCart(3)
		tShirtCart := CartMother.TShirtCart(4)

		if len(emptyCart.LineItems()) != 0 {
			t.Errorf("expected 0 line items, got %d", len(emptyCart.LineItems()))
		}
		if len(voucherCart.LineItems()) != 1 {
			t.Fatalf("expected 1 line item, got %d", len(voucherCart.LineItems()))
		}
		if voucherCart.LineItems()[0].Product.Code != "VOUCHER" {
			t.Errorf("expected VOUCHER, got %s", voucherCart.LineItems()[0].Product.Code)
		}
		if voucherCart.LineItems()[0].Quantity != 3 {
			t.Errorf("expected quantity 3, got %d", voucherCart.LineItems()[0].Quantity)
		}
		if len(tShirtCart.LineItems()) != 1 {
			t.Fatalf("expected 1 line item, got %d", len(tShirtCart.LineItems()))
		}
		if tShirtCart.LineItems()[0].Product.Code != "TSHIRT" {
			t.Errorf("expected TSHIRT, got %s", tShirtCart.LineItems()[0].Product.Code)
		}
		if tShirtCart.LineItems()[0].Quantity != 4 {
			t.Errorf("expected quantity 4, got %d", tShirtCart.LineItems()[0].Quantity)
		}
	})
}
