// Пакет personaldata описывает персональные данные пользователя.
package personaldata

import "fmt"

// Personal хранит имя, вес (кг) и рост (м) пользователя.
type Personal struct {
	Name   string
	Weight float64
	Height float64
}

// Print выводит персональные данные на экран.
func (p Personal) Print() {
	fmt.Printf("Имя: %s\n", p.Name)
	fmt.Printf("Вес: %.2f\n", p.Weight)
	fmt.Printf("Рост: %.2f\n", p.Height)
}
