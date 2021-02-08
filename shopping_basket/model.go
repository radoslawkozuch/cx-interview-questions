package main

import (
	"errors"
	"math"
)

type Product string
type Cost float32

func (c *Cost) Round() {
	*c = Cost(math.Round(float64(*c*100)) / 100)
}

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

type SpecialOffer struct {
	applicableProducts []Product
	requiredAmount     int
}

type Offers struct {
	discounts        map[Product]int
	howManyToGetFree map[Product]int
	specialOffers    []SpecialOffer
}

func (o *Offers) GetDiscount(p Product) int {
	return o.discounts[p]
}

func (o *Offers) HowManyToGetFree(p Product) int {
	return o.howManyToGetFree[p]
}

func (o *Offers) GetSpecialOffers() []SpecialOffer {
	return o.specialOffers
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
