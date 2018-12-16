package unittest

import (
	"net/http"
	"testing"
)

const checkMark = "\u2713"
const ballotX = "\u2717"

func TestDownload(t *testing.T) {
	url := "https://golang.org/pkg/fmt"
	statusCode := 200

	t.Log("Given the need to test downloading content")
	{
		t.Logf("\tWhen checking \"%s\" for status code \"%d\"", url, statusCode)
		{
			res, err := http.Get(url)
			if err != nil {
				t.Fatal("\t\tShould be able to make the Get call.", ballotX, err)
			}
			t.Log("\t\tShould be able to make the Get call.", checkMark)

			defer res.Body.Close()

			if res.StatusCode == statusCode {
				t.Logf("\t\tShould receive a \"%d\" status. %v", statusCode, checkMark)
			} else {
				t.Errorf("\t\tShould receive a \"%d\" status. %v %d", statusCode, ballotX, res.StatusCode)
			}
		}
	}
}
