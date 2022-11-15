package sec

import (
	"github.com/dli-invest/sec-parser-go/pkg/types"
	"github.com/dli-invest/sec-parser-go/pkg/utils"
)

// make post request to https://efts.sec.gov/LATEST/search-index
func ParseCik() (cikData []types.CikData, err error) {
	url := "https://www.sec.gov/include/ticker.txt"
	jsonValue, _ := json.Marshal(values)
	req, err := http.NewRequest("GET", url, bytes.NewBuffer(jsonValue))
	if err != nil {
		// Adjust error handling here to meet application requrirements.
		fmt.Println("Error: ", err)
		log.Fatal(err)
	}
	req.Header.Set("origin", "https://www.sec.gov")
	req.Header.Set("referer", "https://www.sec.gov/edgar/searchedgar/companysearch.html")
	req.Header.Set("User-Agent", "Sample dli-invest dlcoding@gmail.com")
	req.Header.Set("Accept-Encoding", "gzip, deflate")
	req.Header.Set("Host", "www.sec.gov")
	defer req.Body.Close()

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		// Adjust error handling here to meet application requrirements.
		fmt.Println("Error: ", err)
		log.Fatal(err)
	}

	b, err := io.ReadAll(resp.Body)
	// split the response into lines
	lines := strings.Split(string(b))
	var cikList []CikData
	// loop through the lines
	for _, line := range lines {
		// split each line into fields
		fields := strings.Split(line, "\t")
		// field[0] is the company name and should only have alpha characters
		// check fields[0] to make sure it is a company name
		if !IsStringAlpha(fields[0]) {
			log.Println(fields[0])
			log.Fatal("invalid data goes here", fields[0])
			return [], errors.New("company contains non letter characters")
		}
		if !strings.ContainsAny(fields[1], "0123456789") {
			// field[1] is the CIK
			log.Println(fields[1])
			log.Fatal("invalid data goes here", fields[1])
			return [], errors.New("cik contains non number characters")
		}
		// append the company name and cik to the cikList
		cikList = append(cikList, CikData{company: fields[0], cik: fields[1]})
		// CikData
		// field[1] is the CIK number and should only have numeric characters

		// check if the first field is the company we are looking for
	}
	return
}