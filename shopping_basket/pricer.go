package main

type basketPricer struct {
	catalogue Catalogue
	offers    Offers
}

func NewBasketPricer(catalogue Catalogue, offers Offers) *basketPricer {
	return &basketPricer{
		catalogue: catalogue,
		offers:    offers,
	}
}

type Bill struct {
	subtotal Cost
	discount Cost
	total    Cost
}

func (p *basketPricer) GetPrice(b *Basket) (*Bill, error) {
	products := b.GetAll()
	var subtotal Cost
	var discount Cost

	for product, amount := range products {
		price, err := p.catalogue.GetPrice(product)
		if err != nil {
			return nil, err
		}
		subtotal += price * Cost(amount)

		// TODO: both offers
		sale := p.offers.GetDiscount(product)
		if sale > 0 && sale < 100 {
			discount += Cost(sale) * price * Cost(amount) / Cost(100)
		}

		getFree := p.offers.HowManyToGetFree(product)
		if getFree > 0 {
			howMany := amount / (getFree + 1)
			discount += Cost(howMany) * price
		}
	}

	return &Bill{
		subtotal: subtotal,
		discount: discount,
		total:    subtotal - discount,
	}, nil
}
