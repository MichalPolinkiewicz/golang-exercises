package main

import (
	"fmt"
	"math/rand"
	"time"
)

var colors = []string{"spades", "heart", "diamonds", "clubs"}
var values = []string{"Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King", "Ace"}

type card struct {
	color string
	value string
}

type deck struct {
	cards []card
}

func (c *card) String() string {
	return fmt.Sprintf("%s of %s", c.value, c.color)
}

func (d *deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	cards := d.cards

	for i := range cards {
		newPosition := r.Intn(len(cards) - 1)
		cards[i], cards[newPosition] = cards[newPosition], cards[i]
	}
}

func (d *deck) sort() {

}

func New() deck{
	var deck deck
	for _, c := range colors {
		for _,v := range values {
			deck.cards = append(deck.cards, card{c, v})
		}
	}
	return deck
}

func (d *deck) exclude(f func(s *string) bool) {
	cards := d.cards
	for i, r := range cards {
		if f(&r.value) {
			copy(cards[i:], cards[i+1:])
			cards[len(cards)-1] = card{}
			cards = cards[:len(cards)-1]
		}
	}
	d.cards = cards
}

func main() {

}
