package stats

import (
	"bytes"
	"encoding/json"
	. "github.com/franela/goblin"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

const token = "5672139asdaw"

func TestFirstGetStats(t *testing.T) {

	fakeTest := testing.T{}
	g := Goblin(&fakeTest)

	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	g.Describe("Get detail stats for the company", func() {

		g.Describe("fail request with invalid order by", func() {

			requestBody, err := json.Marshal(map[string]string{
				"companyid": "1",
				"order": "create",
				"from": "2019-08-01",
				"to": "2019-11-03",
			})

			if err != nil {
				t.Fatal(err)
			}

			g.It("It should return error, wrong request", func() {

				req, err := http.NewRequest("POST", "http://localhost:8080/statistic/stats", bytes.NewBuffer(requestBody))
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

		g.Describe("fail request with invalid order by", func() {

			requestBody, err := json.Marshal(map[string]string{
				"companyid": "1",
				"order": "opened",
				"from": "2019-08-01",
				"to": "2019-11-03",
			})

			if err != nil {
				t.Fatal(err)
			}

			g.It("It should pass", func() {

				req, err := http.NewRequest("POST", "http://localhost:8080/statistic/stats", bytes.NewBuffer(requestBody))
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

			})

		})
	})
}