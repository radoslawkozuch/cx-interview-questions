package main

import (
	"errors"
)

type Product string
type Cost float32

var ProductNotFoundError = errors.New("Product Not Found")

type Catalogue struct {
	prices map[Product]Cost
}

func (c *Catalogue) GetPrice(p Product) (Cost, error) {
	if price, ok := c.prices[p]; ok {
		return price, nil
	}
	return 0, ProductNotFoundError
}

type Offers struct {
	discounts        map[Product]int
	howManyToGetFree map[Product]int
}

func (o *Offers) GetDiscount(p Product) int {
	return o.discounts[p]
}

func (o *Offers) HowManyToGetFree(p Product) int {
	return o.howManyToGetFree[p]
}

type Basket struct {
	products map[Product]int
}

func NewBasket() *Basket {
	return &Basket{
		products: make(map[Product]int),
	}
}

func (b *Basket) AddProduct(p Product, amount int) {
	b.products[p] += amount
	if b.products[p] <= 0 {
		delete(b.products, p)
	}
}

func (b *Basket) GetAmount(p Product) int {
	return b.products[p]
}

func (b *Basket) GetAll() map[Product]int {
	return b.products
}

func (b *Basket) ClearBasket() {
	b.products = make(map[Product]int)
}
