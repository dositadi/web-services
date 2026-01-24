package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

var roman map[int]string = map[int]string{
	1: "I", 2: "II", 3: "III", 4: "IV", 5: "V", 6: "VI", 7: "VII", 8: "VIII", 9: "IX", 10: "X",
}

var word map[int]string = map[int]string{
	1: "One", 2: "Two", 3: "Three", 4: "Four", 5: "Five", 6: "Six", 7: "Seven", 8: "Eight", 9: "Nine", 10: "Ten",
}

func ReturnWord(key int) string {
	if key <= 0 {
		return "The value is either 0 or less than 0!"
	}

	if key > 10 {
		return "The value is greater than 10!"
	}

	if _, ok := roman[key]; ok {
		return word[key]
	}
	return "Incorrect value!"
}

func ReturnRomanNumeral(key int) string {
	if key <= 0 {
		return "The value is either 0 or less than 0!"
	}

	if key > 10 {
		return "The value is greater than 10!"
	}

	if _, ok := roman[key]; ok {
		return roman[key]
	}
	return "Incorrect value!"
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Path

		values := strings.Split(query, "/")[1:]

		var value string
		var db string

		for i, val := range values {
			if i == 0 {
				db = val
			}
			if i == 1 {
				value = val
			}
		}

		number, err := strconv.Atoi(value)
		if err != nil {
			log.Fatal("Integer convert error: ", err)
		}

		switch db {
		case "word":
			fmt.Fprintf(w, " Decimal digit: %d, Word value: %s", number, ReturnWord(number))
		case "roman":
			fmt.Fprintf(w, " Decimal digit: %d, Roman Numeral value: %s", number, ReturnRomanNumeral(number))
		default:
			fmt.Fprintf(w, "Invalid path request, %s", r.URL.Path)
		}
	})

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		fmt.Println("Listen and serve error: ", err)
		return
	}
}
