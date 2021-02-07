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
	assert.Equal(Cost(90), bill.total)
	assert.Equal(Cost(100), bill.subtotal)
	assert.Equal(Cost(10), bill.discount)
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
	assert.Equal(Cost(100), bill.total)
	assert.Equal(Cost(100), bill.subtotal)
	assert.Equal(Cost(0), bill.discount)

	b.AddProduct(product, 1)
	bill, err = pricer.GetPrice(b)
	assert.NoError(err)
	assert.NotNil(bill)
	assert.Equal(Cost(200), bill.total)
	assert.Equal(Cost(200), bill.subtotal)
	assert.Equal(Cost(0), bill.discount)

	b.AddProduct(product, 1)
	bill, err = pricer.GetPrice(b)
	assert.NoError(err)
	assert.NotNil(bill)
	assert.Equal(Cost(200), bill.total)
	assert.Equal(Cost(300), bill.subtotal)
	assert.Equal(Cost(100), bill.discount)

	b.AddProduct(product, 27)
	bill, err = pricer.GetPrice(b)
	assert.NoError(err)
	assert.NotNil(bill)
	assert.Equal(Cost(2000), bill.total)
	assert.Equal(Cost(3000), bill.subtotal)
	assert.Equal(Cost(1000), bill.discount)

	b.AddProduct(product, 1)
	bill, err = pricer.GetPrice(b)
	assert.NoError(err)
	assert.NotNil(bill)
	assert.Equal(Cost(2100), bill.total)
	assert.Equal(Cost(3100), bill.subtotal)
	assert.Equal(Cost(1000), bill.discount)
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
