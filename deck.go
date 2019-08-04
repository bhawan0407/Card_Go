package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

// create a new type 'deck'
// which is an array of strings
type deck []string

func newDeck() deck {

	cards := deck{}

	cardSuits := []string{"Clubs", "Spades", "Diamonds", "Hearts"}
	cardValues := []string{"Ace", "1", "2", "3", "4", "5", "6", "7", "8", "9", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards

}

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func (d deck) deal(handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

func newDeckFromFile(fileName string) deck {
	bs, err := ioutil.ReadFile(fileName)

	if err != nil {
		fmt.Println("Error: ", err)
		return deck{}
	}

	deckString := string(bs)

	return deck(strings.Split(deckString, ","))
}

func (d deck) shuffle() {

	source := rand.NewSource(time.Now().UnixNano())

	r := rand.New(source)

	for i := range d {
		tmp := r.Intn(len(d))
		d[tmp], d[i] = d[i], d[tmp]
	}
}
