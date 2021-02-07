package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBasket(t *testing.T) {

	assert := assert.New(t)

	product := Product("product aaa")

	b := NewBasket()
	b.AddProduct(product, 1)
	assert.Equal(1, b.GetAmount(product))

	b.AddProduct(product, 2)
	assert.Equal(3, b.GetAmount(product))

	b.AddProduct(product, -1)
	assert.Equal(2, b.GetAmount(product))

	b.AddProduct(product, -5)
	assert.Equal(0, b.GetAmount(product))

	b.AddProduct(product, -5)
	assert.Equal(0, b.GetAmount(product))
}
