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

	return &Bill{
		subtotal: subtotal,
		discount: discount,
		total:    subtotal - discount,
	}, nil
}
