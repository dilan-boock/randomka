package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
)

func play(chislo int, str string) bool {
	number := 0
	for i := 10; i > 0; i-- {
		fmt.Println(str)
		fmt.Fscan(os.Stdin, &number)
		switch {
		case number < chislo:
			fmt.Println("Слишком мало, попробуйте еще раз. У вас осталось ", i-1, " попыток")
		case number > chislo:
			fmt.Println("Слишком много, попробуйте еще раз. У вас осталось ", i-1, " попыток")
		case number == chislo:
			fmt.Println("Вы выиграли!")
			fmt.Println("Вы потратили ", 10-i, " попыток")
			return true
		}
	}
	fmt.Println("Вы проиграли, начните еще раз!")
	return false
}
func randomizer(level int) (int, string, error) {
	switch level {
	case 1:
		return rand.Intn(100), "Введите число от 0 до 100", nil
	case 2:
		return rand.Intn(500), "Введите число от 0 до 500", nil
	case 3:
		return rand.Intn(1000), "Введите число от 0 до 1000", nil
	default:
		fmt.Println("Вы ошиблись цифрой")
		return 0, "", errors.New("error happened")
	}
}

func main() {
	var level, chislo int
	var str string
	var err error
	for {
		fmt.Println("Выберите 1, 2 или 3")
		fmt.Fscan(os.Stdin, &level)
		chislo, str, err = randomizer(level)
		if err == nil {
			break
		}
		fmt.Println(err)
	}
	if win := play(chislo, str); win {
		main()
	}
}
