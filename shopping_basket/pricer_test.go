package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDiscount(t *testing.T) {
	assert := assert.New(t)

	product := Product("product aaa")

	b := NewBasket()
	b.AddProduct(product, 1)

	catalogue := Catalogue{
		prices: map[Product]Cost{
			product: 100,
		},
	}

	offers := Offers{
		discounts: map[Product]int{
			product: 10,
		},
	}

	pricer := NewBasketPricer(catalogue, offers)
	bill, err := pricer.GetPrice(b)
	assert.NoError(err)
	assert.NotNil(bill)
	assert.Equal(Cost(90), bill.GetTotal())
	assert.Equal(Cost(100), bill.GetSubtotal())
	assert.Equal(Cost(10), bill.GetDiscount())
}

func TestGetFree(t *testing.T) {
	assert := assert.New(t)

	product := Product("product aaa")

	catalogue := Catalogue{
		prices: map[Product]Cost{
			product: 100,
		},
	}

	offers := Offers{
		howManyToGetFree: map[Product]int{
			product: 2,
		},
	}

	b := NewBasket()
	b.AddProduct(product, 1)
	pricer := NewBasketPricer(catalogue, offers)
	bill, err := pricer.GetPrice(b)
	assert.NoError(err)
	assert.NotNil(bill)
	assert.Equal(Cost(100), bill.GetTotal())
	assert.Equal(Cost(100), bill.GetSubtotal())
	assert.Equal(Cost(0), bill.GetDiscount())

	b.AddProduct(product, 1)
	bill, err = pricer.GetPrice(b)
	assert.NoError(err)
	assert.NotNil(bill)
	assert.Equal(Cost(200), bill.GetTotal())
	assert.Equal(Cost(200), bill.GetSubtotal())
	assert.Equal(Cost(0), bill.GetDiscount())

	b.AddProduct(product, 1)
	bill, err = pricer.GetPrice(b)
	assert.NoError(err)
	assert.NotNil(bill)
	assert.Equal(Cost(200), bill.GetTotal())
	assert.Equal(Cost(300), bill.GetSubtotal())
	assert.Equal(Cost(100), bill.GetDiscount())

	b.AddProduct(product, 27)
	bill, err = pricer.GetPrice(b)
	assert.NoError(err)
	assert.NotNil(bill)
	assert.Equal(Cost(2000), bill.GetTotal())
	assert.Equal(Cost(3000), bill.GetSubtotal())
	assert.Equal(Cost(1000), bill.GetDiscount())

	b.AddProduct(product, 1)
	bill, err = pricer.GetPrice(b)
	assert.NoError(err)
	assert.NotNil(bill)
	assert.Equal(Cost(2100), bill.GetTotal())
	assert.Equal(Cost(3100), bill.GetSubtotal())
	assert.Equal(Cost(1000), bill.GetDiscount())
}

func TestBothOffer(t *testing.T) {
	assert := assert.New(t)

	product := Product("product aaa")

	catalogue := Catalogue{
		prices: map[Product]Cost{
			product: 100,
		},
	}

	offers := Offers{
		howManyToGetFree: map[Product]int{
			product: 2,
		},
		discounts: map[Product]int{
			product: 35,
		},
	}

	b := NewBasket()
	b.AddProduct(product, 3)

	//take 35%
	pricer := NewBasketPricer(catalogue, offers)
	bill, err := pricer.GetPrice(b)
	assert.NoError(err)
	assert.NotNil(bill)
	assert.Equal(Cost(195), bill.GetTotal())
	assert.Equal(Cost(300), bill.GetSubtotal())
	assert.Equal(Cost(105), bill.GetDiscount())

	offers.discounts[product] = 30
	//take "buy 2, get 1" free
	bill, err = pricer.GetPrice(b)
	assert.NoError(err)
	assert.NotNil(bill)
	assert.Equal(Cost(200), bill.GetTotal())
	assert.Equal(Cost(300), bill.GetSubtotal())
	assert.Equal(Cost(100), bill.GetDiscount())

	b.AddProduct(product, 1)
	//take "buy 2, get 1" free, and 30% for 4th product
	bill, err = pricer.GetPrice(b)
	assert.NoError(err)
	assert.NotNil(bill)
	assert.Equal(Cost(270), bill.GetTotal())
	assert.Equal(Cost(400), bill.GetSubtotal())
	assert.Equal(Cost(130), bill.GetDiscount())
}

func TestEmptyBasket(t *testing.T) {
	assert := assert.New(t)

	b := NewBasket()
	catalogue := Catalogue{}
	offers := Offers{}

	pricer := NewBasketPricer(catalogue, offers)
	bill, err := pricer.GetPrice(b)
	assert.NoError(err)
	assert.NotNil(bill)
	assert.Equal(Cost(0), bill.GetTotal())
	assert.Equal(Cost(0), bill.GetSubtotal())
	assert.Equal(Cost(0), bill.GetDiscount())
}

func TestIncorrectBasket(t *testing.T) {
	assert := assert.New(t)

	product := Product("product aaa")

	b := &Basket{
		products: map[Product]int{
			product: -1,
		},
	}

	catalogue := Catalogue{
		prices: map[Product]Cost{
			product: 100,
		},
	}

	offers := Offers{
		discounts: map[Product]int{
			product: 10,
		},
	}

	pricer := NewBasketPricer(catalogue, offers)
	bill, err := pricer.GetPrice(b)
	assert.Error(err)
	assert.Nil(bill)

	b.products[product] = 1
	catalogue.prices[product] = -100
	pricer = NewBasketPricer(catalogue, offers)
	bill, err = pricer.GetPrice(b)
	assert.Error(err)
	assert.Nil(bill)
}

func TestUnknownProduct(t *testing.T) {
	assert := assert.New(t)

	product := Product("unknown")

	b := NewBasket()
	b.AddProduct(product, 1)

	catalogue := Catalogue{}
	offers := Offers{}

	pricer := NewBasketPricer(catalogue, offers)
	bill, err := pricer.GetPrice(b)
	assert.Error(err)
	assert.Nil(bill)
}

func TestFromReadme(t *testing.T) {
	assert := assert.New(t)

	beans := Product("Baked Beans")
	biscuits := Product("Biscuits")
	sardines := Product("Sardines")
	shampooSmall := Product("Shampoo (Small)")
	shampooMedium := Product("Shampoo (Medium)")
	shampooLarge := Product("Shampoo (Large)")

	catalogue := Catalogue{
		prices: map[Product]Cost{
			beans:         0.99,
			biscuits:      1.20,
			sardines:      1.89,
			shampooSmall:  2,
			shampooMedium: 2.50,
			shampooLarge:  3.50,
		},
	}

	offers := Offers{
		howManyToGetFree: map[Product]int{
			beans: 2,
		},
		discounts: map[Product]int{
			sardines: 25,
		},
	}

	basket1 := NewBasket()
	basket1.AddProduct(beans, 4)
	basket1.AddProduct(biscuits, 1)

	pricer := NewBasketPricer(catalogue, offers)
	bill, err := pricer.GetPrice(basket1)
	assert.NoError(err)
	assert.NotNil(bill)
	assert.Equal(Cost(4.17), bill.GetTotal())
	assert.Equal(Cost(5.16), bill.GetSubtotal())
	assert.Equal(Cost(0.99), bill.GetDiscount())

	basket2 := NewBasket()
	basket2.AddProduct(beans, 2)
	basket2.AddProduct(biscuits, 1)
	basket2.AddProduct(sardines, 2)

	bill, err = pricer.GetPrice(basket2)
	assert.NoError(err)
	assert.NotNil(bill)
	assert.Equal(Cost(6.01), bill.GetTotal())
	assert.Equal(Cost(6.96), bill.GetSubtotal())
	assert.Equal(Cost(0.95), bill.GetDiscount())

}
