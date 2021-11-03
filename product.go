package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Product struct {
	Id    uint32  `json:"id,omitempty"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

var (
	productId  uint32 = 100
	productMut sync.Mutex
	products   = map[uint32]Product{
		1: Product{
			Id:    1,
			Name:  "Kreisel",
			Price: 6.34,
		},
		2: Product{
			Id:    2,
			Name:  "Eisenbahn",
			Price: 299.99,
		},
		3: Product{
			Id:    3,
			Name:  "Toaster",
			Price: 19.49,
		},
		4: Product{
			Id:    4,
			Name:  "Swag",
			Price: 13.37,
		},
		5: Product{
			Id:    5,
			Name:  "Grow Tent",
			Price: 34.88,
		},
	}
)

func getNextProductId() uint32 {
	return atomic.AddUint32(&productId, 1)
}

func getAllProducts() []Product {
	productMut.Lock()
	defer productMut.Unlock()

	prods := make([]Product, len(products))

	i := 0
	for _, v := range products {
		prods[i] = v
		i++
	}

	return prods
}

func getProduct(id uint32) (Product, error) {
	productMut.Lock()
	defer productMut.Unlock()

	prod, ok := products[id]
	if !ok {
		return Product{}, fmt.Errorf("could not find product")
	}

	return prod, nil
}

func addProduct(p Product) Product {
	productMut.Lock()
	defer productMut.Unlock()

	p.Id = getNextProductId()
	products[p.Id] = p

	return p
}

func editProduct(p Product) error {
	productMut.Lock()
	defer productMut.Unlock()

	_, ok := products[p.Id]
	if !ok {
		return fmt.Errorf("could not find product for this Id")
	}
	products[p.Id] = p

	return nil
}

func removeProduct(id uint32) error {
	productMut.Lock()
	defer productMut.Unlock()

	_, ok := products[id]
	if !ok {
		return fmt.Errorf("could not find product for this Id")
	}
	delete(products, id)

	return nil
}
