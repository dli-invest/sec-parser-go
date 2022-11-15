// test 

package sec

import (
	"testing"
	"github.com/dli-invest/sec-parser-go/pkg/sec"
)

// make read the csv read back into the file is the same
func TestParseCik(t *testing.T) {
	cikData := ParseCik()
	if len(cikData) == 0 {
		t.Error("No cik data was returned")
	}
}
