package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

type Equity struct {
	SCCode string `json:"sc_code"`
	SCName string `json:"sc_name"`
	SCGoup string `json:"sc_type"`
}

func main() {
	csvFile, _ := os.Open("EQ150219.CSV")
	reader := csv.NewReader(bufio.NewReader(csvFile))
	var equity []Equity
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		equity = append(equity, Equity{
			SCCode: line[0],
			SCName: line[1],
			SCGoup: line[2],
		})
	}
	// peopleJson, _ := json.Marshal(equity)
	// fmt.Println(string(peopleJson))

	for i := 0; i < len(equity) && i < 10; i++ {
		fmt.Printf("----------------------\n")
		fmt.Printf("|%-20s|\n", equity[i].SCName)
	}
	fmt.Printf("----------------------\n")

}
