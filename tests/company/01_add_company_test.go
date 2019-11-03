package company

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

func TestAddCompany(t *testing.T) {

	fakeTest := testing.T{}
	g := Goblin(&fakeTest)

	client := &http.Client{
		Timeout: 5 * time.Second,
	}


	g.Describe("Add company test", func() {

		g.Describe("fail request", func() {
			// fail request
			requestBody, err := json.Marshal(map[string] string {
				"id": "0",
			})

			if err != nil {
				t.Fatal(err)
			}

			g.It("It should return error, wrong request", func() {

				req, err := http.NewRequest("POST", "http://localhost:8080/company/add", bytes.NewBuffer(requestBody))
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
			})

			if err != nil {
				t.Fatal(err)
			}

			g.It("It should pass", func() {

				req, err := http.NewRequest("POST", "http://localhost:8080/company/add", bytes.NewBuffer(requestBody))
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
