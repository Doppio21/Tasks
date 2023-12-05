package main

import (
	"encoding/json"
	"fmt"
)

// Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
// Объединить данные значения в группы с шагом в 10 градусов. Последовательность в подмножноствах не важна.
// Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.

func main() {
	temperature := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	groups := make(map[int][]float64)
	step := 10

	for _, t := range temperature {
		group := int(t) - int(t)%step
		groups[group] = append(groups[group], t)
	}

	data, err := json.MarshalIndent(&groups, " ", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))
}
