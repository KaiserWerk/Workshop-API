package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Review struct {
	Id        uint32 `json:"id"`
	ProductId uint32 `json:"product_id"`
	Rating    uint8  `json:"rating"`
	Text      string `json:"text"`
}

var (
	reviewMut sync.Mutex
	reviewId  uint32 = 100
	reviews          = map[uint32]Review{
		1: Review{
			Id:        1,
			ProductId: 1,
			Rating:    3,
			Text:      "Gelungenes Produkt zum fairen Preis, könnte allerdings länger haltbar sein.",
		},
	}
)

func getNextReviewId() uint32 {
	return atomic.AddUint32(&reviewId, 1)
}

func GetAllReviews(productId uint32) []Review {
	reviewMut.Lock()
	defer reviewMut.Unlock()

	revs := make([]Review, len(reviews))

	i := 0
	for _, v := range reviews {
		if v.ProductId == productId {
			revs[i] = v
			i++
		}
	}

	return revs
}

func GetReview(id uint32) (Review, error) {
	reviewMut.Lock()
	defer reviewMut.Unlock()

	r, ok := reviews[id]
	if !ok {
		return Review{}, fmt.Errorf("could not find Review for Id")
	}

	return r, nil
}

func AddReview(productId uint32, r Review) Review {
	reviewMut.Lock()
	defer reviewMut.Unlock()

	r.Id = getNextReviewId()
	r.ProductId = productId
	reviews[r.Id] = r
	return r
}

func EditReview(r Review) error {
	reviewMut.Lock()
	defer reviewMut.Unlock()

	original, ok := reviews[r.Id]
	if !ok {
		return fmt.Errorf("could not find entry for this Id")
	}
	r.ProductId = original.ProductId // product ID is not changeable
	reviews[r.Id] = r

	return nil
}
