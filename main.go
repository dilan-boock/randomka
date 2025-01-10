package main

import (
	"fmt"
	"math/rand"
	"os"
)

var level, lev, chislo int
var str string

func hello() {
	fmt.Println("Выберите 1, 2 или 3")
	fmt.Fscan(os.Stdin, &level)
	randomizer()
}

func play() {
	for i := 10; i > 0; i-- {
		fmt.Println(str)
		fmt.Fscan(os.Stdin, &chislo)
		if chislo < lev {
			fmt.Println("Слишком мало, попробуйте еще раз. У вас осталось ", i-1, " попыток")
		}
		if chislo > lev {
			fmt.Println("Слишком много, попробуйте еще раз. У вас осталось ", i-1, " попыток")
		}
		if chislo == lev {
			fmt.Println("Вы выиграли!")
			fmt.Println("Вы потратили ", 10-i, " попыток")
			break
		}
	}
	fmt.Println("Вы проиграли, начните еще раз!")
}
func randomizer() {
	switch lev = level; lev {
	case 1:
		lev = rand.Intn(100)
		//fmt.Println("Ваш уровень ", level, "Ваш рандом = ", lev)
		str = "Введите число от 0 до 100"
	case 2:
		lev = rand.Intn(500)
		//fmt.Println("Ваш уровень ", level, "Ваш рандом = ", lev)
		str = "Введите число от 0 до 500"
	case 3:
		lev = rand.Intn(1000)
		//fmt.Println("Ваш уровень ", level, "Ваш рандом = ", lev)
		str = "Введите число от 0 до 1000"
	default:
		fmt.Println("Вы ошиблись цифрой")
		hello()
	}
}

func main() {
	hello()
	play()
}
