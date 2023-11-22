////go:build integration

package todo_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"testing"
)

type Todo struct {
	Id          string `json:"id"`
	Description string `json:"description"`
}

func TestIntegrationRequests(t *testing.T) {
	type expected struct {
		httpCode int
		Payload  interface{}
	}
	testCases := []struct {
		Method   string
		Uri      string
		Params   url.Values
		Payload  interface{} // Payload for POST requests
		Expected expected
	}{
		//{
		//	Method: http.MethodGet,
		//	Params: url.Values{
		//		"param1": []string{"value1"},
		//		"param2": []string{"value2"},
		//	},
		//},
		{
			Method: http.MethodPost,
			Uri:    "/todo",
			Payload: Todo{
				Description: "Have a breakfast",
			},
			Expected: expected{
				httpCode: http.StatusCreated,
				Payload: Todo{
					Description: "Have a breakfast",
				},
			},
		},
		// Add more test cases as needed
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Testing %s %s", tc.Method, tc.Uri), func(t *testing.T) {
			requestURL := fmt.Sprintf("http://localhost:8086/api%s", tc.Uri)
			var req *http.Request
			var err error

			if tc.Method == http.MethodGet {
				requestURL += "?" + tc.Params.Encode()
				req, err = http.NewRequest(tc.Method, requestURL, nil)
			} else if tc.Method == http.MethodPost {
				payload, _ := json.Marshal(tc.Payload)
				req, err = http.NewRequest(tc.Method, requestURL, bytes.NewBuffer(payload))
				req.Header.Set("Content-Type", "application/json")
			}

			if err != nil {
				t.Fatalf("failed to create request: %v", err)
			}

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				t.Fatalf("request failed: %v", err)
			}
			defer resp.Body.Close()
			assert.Equal(t, tc.Expected.httpCode, resp.StatusCode)
		})
	}
}
