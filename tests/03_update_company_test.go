package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	. "github.com/franela/goblin"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

func TestUpdateCompany(t *testing.T) {

	fakeTest := testing.T{}
	g := Goblin(&fakeTest)

	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	g.Describe("Update company test", func() {

		g.Describe("fail request with id = 0 and stats <= 0", func() {
			// fail request
			requestBody, err := json.Marshal(map[string] string {
				"id": "0",
				"totallocations": "0",
				"totaldoctors": "1",
				"totalusers": "1",
				"totalinvitations": "2011",
				"totalcreatedreviews": "0",
				"totalopenedreviews": "-1",
			})

			if err != nil {
				t.Fatal(err)
			}

			g.It("It should return error, wrong request", func() {

				req, err := http.NewRequest("POST", "http://localhost:8080/company/update", bytes.NewBuffer(requestBody))
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

			requestBody, err := json.Marshal(map[string] string {
				"id": "1",
				"totallocations": "0",
				"totaldoctors": "1",
				"totalusers": "1",
				"totalinvitations": "2011",
				"totalcreatedreviews": "0",
				"totalopenedreviews": "1",
			})

			if err != nil {
				t.Fatal(err)
			}

			g.It("It should pass", func() {

				req, err := http.NewRequest("POST", "http://localhost:8080/company/update", bytes.NewBuffer(requestBody))
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

		g.Describe("test that company will be correct modified if one of stats is less than 0", func() {

			requestBody, err := json.Marshal(map[string] string {
				"id": "1",
				"totallocations": "0",
				"totaldoctors": "1",
				"totalusers": "1",
				"totalinvitations": "-1",
				"totalcreatedreviews": "0",
				"totalopenedreviews": "1",
			})

			if err != nil {
				t.Fatal(err)
			}

			g.It("It should pass and set 0 where data was negative", func() {

				req, err := http.NewRequest("POST", "http://localhost:8080/company/update", bytes.NewBuffer(requestBody))
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


		g.Describe("get stats after correct and failed updates", func() {

			id := 1

			g.It("It should return correct stats", func() {

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
				g.Assert(data["totalusers"].(float64)).Equal(float64(2))
				g.Assert(data["totalinvitations"].(float64)).Equal(float64(2011))

			})

		})

	})

}
