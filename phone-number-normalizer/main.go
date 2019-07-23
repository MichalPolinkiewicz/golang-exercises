package main

import (
	"fmt"
	"github.com/MichalPolinkiewicz/golang-exercises/phone-number-normalizer/persistence"
	"regexp"
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
	defer persistence.DbInstance.Close()

	numbers := persistence.GetNumbers()
	for _, n := range numbers {
		fmt.Println(normalize(n))
	}
}

func normalize(phone string) string {
	re := regexp.MustCompile("\\D")
	return re.ReplaceAllString(phone, "")
}

