package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type (
	Card struct {
		face  string
		value int
	}

	Deck struct {
		cards map[string][]Card
	}
)

func (d *Deck) add(card Card) {
	d.cards[card.face] = append(d.cards[card.face], card)
}

func (d *Deck) String() string {
	result := ""
	for k, cards := range d.cards {
		total := 0
		for _, card := range cards {
			total += card.value
		}
		result += fmt.Sprintf("%v = %v\n", k, total)
	}
	return result
}

func generateNumbers(action chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	num := 4
	for i := 3; num <= 200; i += 2 {
		num += i
		action <- num
	}
	close(action)
}

func generateCards(cardCh chan Card, wg *sync.WaitGroup, action chan int) {
	defer wg.Done()
	for num := range action {
		card := Card{
			value: rand.Intn(13) + 1, // Card values between 1 and 13
		}
		if num%2 == 0 {
			card.face = "heart"
		} else {
			card.face = "club"
		}
		cardCh <- card
	}
	close(cardCh)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	deck := &Deck{cards: make(map[string][]Card)}
	var wg sync.WaitGroup

	cardCh := make(chan Card, 10)
	actionCh := make(chan int, 10)

	wg.Add(2)
	go generateNumbers(actionCh, &wg)
	go generateCards(cardCh, &wg, actionCh)

	go func() {
		wg.Wait()
		//close(cardCh)
	}()

	for card := range cardCh {
		deck.add(card)
	}

	fmt.Println(deck.String())
}
