package service

type GameBuffer struct {
	head     int
	Data     []int
	size     int
	Value    int
	modifier int
}

func CreateGameBuffer(size int) *GameBuffer {
	return &GameBuffer{
		head:  0,
		Data:  make([]int, size),
		Value: 0,
		size:  size,
	}
}

func (gb *GameBuffer) AddValue(value int) bool {
	for i := 0; i < len(gb.Data); i++ {
		if gb.Data[i] == 0 {
			gb.Data[i] = value
			// change to calc value so we can get the modifiers
			gb.GetValue()
			return true
		}
	}

	return false
}

func (gb *GameBuffer) RemoveValue(value int) int {
	for i := 0; i < len(gb.Data); i++ {
		if gb.Data[i] != value {
			continue
		}
		gb.Data[i] = 0
	}

	return gb.GetValue()
}

func (gb *GameBuffer) GetValue() int {
	candidate := 0
	count := 0
	sum := 0
	for _, datum := range gb.Data {
		if count == 0 {
			candidate = datum
		}

		if datum == candidate {
			count++
		} else {
			count--
		}
		sum += datum
	}

	count = 0
	for _, num := range gb.Data {
		if num != candidate {
			continue
		}
		count++
	}

	if count > 1 {
		sum -= candidate * count
		sum += (candidate * count) * count
	}

	gb.Value = sum
	return gb.Value
}
