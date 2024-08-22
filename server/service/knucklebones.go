package service

import (
	"errors"
	"fmt"
	"math/rand"
)

type player struct {
	*Board
	id         string
	gameId     string
	lastAction Action
}

type Action struct {
	Action ACTION
	Value  int
}

type ACTION string

const (
	ROLL  ACTION = "ROLL"
	PLACE ACTION = "PLACE"
	SCORE ACTION = "SCORE"
)

type Knucklebones struct {
	Players []*player
	turn    int
	max     int
	min     int
}

func CreateKnucklebonesGame() *Knucklebones {

	players := []*player{
		&player{
			Board:      CreateBoard(),
			id:         "1",
			gameId:     "",
			lastAction: Action{},
		},
		&player{
			Board:      CreateBoard(),
			id:         "",
			gameId:     "",
			lastAction: Action{},
		},
	}

	return &Knucklebones{
		Players: players,
		turn:    0,
		max:     6,
		min:     1,
	}
}

func (k *Knucklebones) PlayerAction(playerId string, action *Action) any {
	if playerId != k.Players[k.turn].id {
		return errors.New("invalid turn order")
	}

	switch action.Action {
	case ROLL:
		return k.RollDie(k.Players[k.turn])
		break
	case PLACE:
		k.PlaceValue(k.Players[k.turn], action)
		break
	case SCORE:
		return k.GetScore(FindPlayer(k.Players, playerId))
		break
	}

	return nil
}

func FindPlayer(p []*player, id string) *player {
	for _, pl := range p {
		if pl.id == id {
			return pl
		}
	}

	return nil
}

func (k *Knucklebones) RollDie(p *player) *Action {
	p.lastAction = Action{
		Action: ROLL,
		Value:  rand.Intn(k.max - k.min),
	}

	return &p.lastAction
}

func (k *Knucklebones) PlaceValue(p *player, action *Action) error {
	if p.lastAction.Action != ROLL {
		return errors.New("last action must be a roll")
	}
	err := p.Board.AddToColumn(action.Value, p.lastAction.Value)
	if err == nil {
		k.NextTurn()
	}
	return err
}

func (k *Knucklebones) NextTurn() bool {
	k.turn = k.turn % len(k.Players)
	return true
}

func (k *Knucklebones) GetScore(p *player) int {
	fmt.Println(p.GetValue())
	return p.GetValue()
}
