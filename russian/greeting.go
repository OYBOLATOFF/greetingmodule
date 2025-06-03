package russian

import (
	"bufio"
	"fmt"
	"os"
)

func Hello() {
	fmt.Println("Передаю привет прямиком из модуля!")
}

func GreetByName(name string) {
	fmt.Printf("Привет, %s", name)
}

func AskNameAndGreet() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Как тебя зовут: ")
	name, _ := reader.ReadString('\n')
	GreetByName(name)
}

func IntroduceMe(name string) {
	fmt.Printf("Леди и джентельмены, это мой лучший друг %s!", name)
}
