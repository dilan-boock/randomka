package main

import (
	"fmt"
	"math/rand"
	"os"
)

type Randomizer struct {
	Level  int
	Lev    int
	Chislo int
	I      int
	Str    string
}

var num = Randomizer{0, 0, 0, 10, " "}

func hello() {
	fmt.Println("Выберите 1, 2 или 3")
	fmt.Fscan(os.Stdin, &num.Level)
}

func play() {
	for i := 10; i > 0; i-- {
		fmt.Println(num.Str)
		fmt.Fscan(os.Stdin, num.Chislo)
		if num.Chislo < num.Lev {
			fmt.Println("Слишком мало, попробуйте еще раз. У вас осталось ", i-1, " попыток")
		}
		if num.Chislo > num.Lev {
			fmt.Println("Слишком много, попробуйте еще раз. У вас осталось ", i-1, " попыток")
		}
		if num.Chislo == num.Lev {
			fmt.Println("Вы выиграли!")
			fmt.Println("Вы потратили ", 10-i, " попыток")
			break
		}
	}
	fmt.Println("Вы проиграли, начните еще раз!")
}
func randomizer() {
	switch num.Lev = num.Level; num.Lev {
	case 1:
		num.Lev = rand.Intn(100)
		//fmt.Println("Ваш уровень ", num.Level, "Ваш рандом = ", num.Lev)
		num.Str = "Введите число от 0 до 100"
	case 2:
		num.Lev = rand.Intn(500)
		//fmt.Println("Ваш уровень ", num.Level, "Ваш рандом = ", num.Lev)
		num.Str = "Введите число от 0 до 500"
	case 3:
		num.Lev = rand.Intn(1000)
		//fmt.Println("Ваш уровень ", level, "Ваш рандом = ", lev)
		num.Str = "Введите число от 0 до 1000"
	default:
		fmt.Println("Вы ошиблись цифрой")
		hello()
	}
}

func main() {
	hello()
	randomizer()
	play()
}
