package greeting

import (
	"bufio"
	"fmt"
	"os"
)

func Hello() {
	fmt.Println("Hello from module!")
}

func GreetByName(name string) {
	fmt.Printf("Hello, %s", name)
}

func AskNameAndGreet() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Как тебя зовут: ")
	name, _ := reader.ReadString('\n')
	GreetByName(name)
}

func IntroduceMe(name string) {
	fmt.Printf("Ladies and gentlemens, this is my best friend %s!", name)
}
