package tests

import (
	"encoding/json"
	"fmt"
	. "github.com/franela/goblin"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

func TestGetStats(t *testing.T) {

	fakeTest := testing.T{}
	g := Goblin(&fakeTest)

	client := &http.Client{
		Timeout: 5 * time.Second,
	}


	g.Describe("Get company stats", func() {

		g.Describe("fail request", func() {

			id := 0

			g.It("It should return error, id = 0 ", func() {

				req, err := http.NewRequest("GET", fmt.Sprintf("http://localhost:8080/company/stats?id=%d",id), nil)
				if err != nil {
					t.Fatal(err)
				}

				req.Header.Set("Content-Type", "application/json")
				req.Header.Set("Token", token)

				resp, err := client.Do(req)
				if err != nil {
					t.Fatal(err)
				}

				defer resp.Body.Close()

				body, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					t.Fatal(err)
				}

				//unmarshal json

				var data map[string]interface{}

				if err := json.Unmarshal(body, &data); err != nil {
					t.Fatal(err)
				}

				g.Assert(data["error"]).Equal("400 Bad Request")

			})

		})

		g.Describe("correct request", func() {

			id := 1

			g.It("It should pass", func() {

				req, err := http.NewRequest("GET", fmt.Sprintf("http://localhost:8080/company/stats?id=%d",id), nil)
				if err != nil {
					t.Fatal(err)
				}

				req.Header.Set("Content-Type", "application/json")
				req.Header.Set("Token", token)

				resp, err := client.Do(req)
				if err != nil {
					t.Fatal(err)
				}

				defer resp.Body.Close()

				body, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					t.Fatal(err)
				}

				//unmarshal json

				var data map[string]interface{}

				if err := json.Unmarshal(body, &data); err != nil {
					t.Fatal(err)
				}

				g.Assert(data["error"]).Equal("")
				g.Assert(data["totallocations"].(float64)).Equal(float64(0))

			})

		})

	})

}
