package service

type player struct {
	*Board
	id     string
	gameId string
}

type Knucklebones struct {
	Player1 *player
	Player2 *player
	turn    *player
}

func CreateKnucklebonesGame() *Knucklebones {

	gb1 := CreateBoard()
	gb2 := CreateBoard()

	return &Knucklebones{
		Player1: &player{
			Board:  gb1,
			id:     "1",
			gameId: "",
		},
		Player2: &player{
			Board:  gb2,
			id:     "2",
			gameId: "",
		},
		turn: nil,
	}
}

func (k *Knucklebones) PlayerAction() error {
	return nil
}
