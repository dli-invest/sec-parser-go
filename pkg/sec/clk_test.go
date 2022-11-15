// test

package sec

import (
	"testing"
)

// make read the csv read back into the file is the same
func TestParseCik(t *testing.T) {
	cikData, err := ParseCik()
	if len(cikData) == 0 {
		t.Fatal("No cik data was returned")
	}
	if err != nil {
		t.Fatal("err")
	}
}
