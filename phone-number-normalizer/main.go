package main

import (
	"fmt"
	"github.com/MichalPolinkiewicz/golang-exercises/phone-number-normalizer/persistence"
)

func main() {
	//err := persistence.CreatePhonesTable()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//err = persistence.InsertSomePhones()
	//if err != nil {
	//	log.Fatal(err)
	//}

	numbers := persistence.GetNumbers()

	for _, n := range numbers {
		normalize(&n)
	}
}

func normalize(number *string){
	fmt.Println(*number)
}

