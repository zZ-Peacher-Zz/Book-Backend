package main

import (
	"book-backend/utils"
	"fmt"
)

func main() {
	var myVar uint = 22323222
	token := utils.GenerateJWT(myVar)
	fmt.Println(token)
}
