package main

import (
	"fmt"

	"github.com/CyberRoute/learn_thirdpartypackages/mapper"
)

func main() {
	fmt.Println(mapper.Greet("Howdy, what's new?"))
	fmt.Println(mapper.Greet("Comment allez vous?"))
	fmt.Println(mapper.Greet("Wie geht es Ihnen?"))
}
