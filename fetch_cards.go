package main

import (
	"fmt"
	"log"
	"os"

	trello "github.com/VojtechVitek/go-trello"
)

var (
	appKey  = os.Getenv("TRELLO_APP_KEY")
	token   = os.Getenv("TRELLO_TOKEN")
	boardId = os.Getenv("TRELLO_BOARD_ID")
)

func main() {

	trello, err := trello.NewAuthClient(appKey, &token)
	if err != nil {
		log.Fatal(err)
	}

	lists, err := allLists(trello)
	if err != nil {
		log.Fatal(err)
	}

	for _, list := range lists {
		log.Printf("Name: %s Id: %s ", list.Name, list.Id)
	}

	var listId string
	fmt.Println("Copy and paste Id of the List")
	fmt.Scan(&listId)

	list, err := trello.List(listId)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Cards in list %s: \n\n", listId)

	cards, _ := list.Cards()
	for _, card := range cards {
		fmt.Println(card.Url)
	}
}

func allLists(trello *trello.Client) ([]trello.List, error) {
	board, err := trello.Board(boardId)
	if err != nil {
		return nil, err
	}
	lists, err := board.Lists()
	if err != nil {
		return nil, err
	}
	return lists, nil
}
