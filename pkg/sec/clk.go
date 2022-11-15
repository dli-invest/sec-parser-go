package sec

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/dli-invest/sec-parser-go/pkg/types"
)

// make post request to https://efts.sec.gov/LATEST/search-index
func ParseCik() (cikData []types.ClkData, err error) {
	url := "https://www.sec.gov/include/ticker.txt"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		// Adjust error handling here to meet application requrirements.
		fmt.Println("Error: ", err)
		log.Fatal(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		// Adjust error handling here to meet application requrirements.
		fmt.Println("Error: ", err)
		log.Fatal(err)
	}

	b, err := io.ReadAll(resp.Body)
	// split the response into lines
	lines := strings.Split(string(b), "\n")
	var clkList []types.ClkData

	// loop through the lines
	for _, line := range lines {
		// split each line into fields
		fields := strings.Split(line, "\t")
		// field[0] is the company name and should only have alpha characters
		// check fields[0] to make sure it is a company name
		// brk-b
		// if !utils.IsStringAlpha(fields[0]) {
		// 	log.Println(fields[0])
		// 	log.Fatal("invalid data goes here: ", fields[0])
		// 	err = errors.New("company contains non letter characters")
		// 	return
		// }
		if !strings.ContainsAny(fields[1], "0123456789") {
			// field[1] is the CIK
			log.Println(fields[1])
			log.Fatal("invalid data goes here", fields[1])
			err = errors.New("cik contains non number characters")
		}
		// append the company name and cik to the cikList
		clkList = append(clkList, types.ClkData{Company: fields[0], Clk: fields[1]})
		// CikData
		// field[1] is the CIK number and should only have numeric characters

		// check if the first field is the company we are looking for
	}
	return clkList, nil
}
