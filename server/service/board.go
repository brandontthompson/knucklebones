package service

import "errors"

type Board struct {
	spaces *[3]GameBuffer
	value  int
}

func CreateBoard() *Board {
	var buf = [3]GameBuffer{}

	for i := 0; i < len(buf); i++ {
		buf[i] = *CreateGameBuffer(3)
	}

	return &Board{spaces: &buf}
}

func (b *Board) AddToColumn(column, value int) error {
	if column > len(b.spaces) || column < 0 {
		return errors.New("column index is out of bounds")
	}
	b.spaces[column].AddValue(value)
	return nil
}

func (b *Board) RemoveFromColumn(column, value int) error {
	if column > len(b.spaces) || column < 0 {
		return errors.New("column index is out of bounds")
	}
	b.spaces[column].RemoveValue(value)
	return nil
}

func (b *Board) GetValue() int {
	for _, buffer := range b.spaces {
		b.value += buffer.GetValue()
	}
	return b.value
}
