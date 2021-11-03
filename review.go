package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Review struct {
	Id        uint32
	ProductId uint32
	Rating    uint8
	Text      string
}

var (
	reviewMut sync.Mutex
	reviewId  uint32 = 100
	reviews          = map[uint32]Review{}
)

func getNextReviewId() uint32 {
	return atomic.AddUint32(&reviewId, 1)
}

func GetAllReviews() []Review {
	reviewMut.Lock()
	defer reviewMut.Unlock()

	revs := make([]Review, 0, len(reviews))

	i := 0
	for _, v := range reviews {
		revs[i] = v
		i++
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
