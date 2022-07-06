package integrationtest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type MathClient struct {
	Token string
	Host  string
}

type AddResult struct {
	Result int `json:"result"`
}

func (c *MathClient) APISum(i, j string) (int, int, error) {
	query := fmt.Sprintf("http://%s/add?a=%s&b=%s&authtoken=%v", c.Host, i, j, c.Token)

	var statusCode int = http.StatusBadRequest
	response, err := http.Get(query)
	if err != nil {
		return 0, statusCode, err
	}
	defer response.Body.Close()
	statusCode = response.StatusCode
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return 0, statusCode, err
	}

	a := AddResult{}
	err = json.Unmarshal(data, &a)
	if err != nil {
		return 0, statusCode, err
	}

	return a.Result, statusCode, nil
}
