package main

import(
	"fmt"
	"os"
	"github.com/joho/godotenv"


)

func main() {
	godotenv.Load()
	linha1 := fmt.Sprintf(" %s ", os.Getenv("LINHA1"))
	linha2 := fmt.Sprintf(" %s ", os.Getenv("LINHA2"))
	linha3 := fmt.Sprintf(" %s ", os.Getenv("LINHA3"))

	fmt.Println(linha1)	
	fmt.Println(linha2)	
	fmt.Println(linha3)
}