package main

import "errors"

var IncorrectBasket = errors.New("IncorrectBasket")

type BasketPricer interface {
	GetPrice(b *Basket) (Bill, error)
}

type Bill interface {
	GetSubtotal() Cost
	GetDiscount() Cost
	GetTotal() Cost
}

func NewBasketPricer(catalogue Catalogue, offers Offers) BasketPricer {
	return &basketPricer{
		catalogue: catalogue,
		offers:    offers,
	}
}

type basketPricer struct {
	catalogue Catalogue
	offers    Offers
}

type bill struct {
	subtotal Cost
	discount Cost
	total    Cost
}

func (b *bill) GetSubtotal() Cost {
	return b.subtotal
}

func (b *bill) GetDiscount() Cost {
	return b.discount
}

func (b *bill) GetTotal() Cost {
	return b.total
}

func (p *basketPricer) GetPrice(b *Basket) (Bill, error) {
	products := b.GetAll()
	var subtotal Cost
	var discount Cost

	for product, amount := range products {
		price, err := p.catalogue.GetPrice(product)
		if err != nil {
			return nil, err
		}

		if price < 0 || amount < 0 {
			return nil, IncorrectBasket
		}

		subtotal += price * Cost(amount)

		sale := p.offers.GetDiscount(product)
		getFree := p.offers.HowManyToGetFree(product)

		// both offers - it is needed to select better
		if sale > 0 && getFree > 0 {
			if sale*(getFree+1) > 100 {
				getFree = 0
			}
		}

		if getFree > 0 {
			howMany := amount / (getFree + 1)
			discount += Cost(howMany) * price
			amount = amount % (getFree + 1)
		}

		if sale > 0 && sale < 100 {
			discount += Cost(sale) * price * Cost(amount) / Cost(100)
		}
	}

	return &bill{
		subtotal: subtotal,
		discount: discount,
		total:    subtotal - discount,
	}, nil
}
