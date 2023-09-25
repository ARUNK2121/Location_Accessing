package helpers

import (
	"encoding/json"
	"io"
	"location/utils"
	"net/http"
)

func SendRequestAndFindPlace(targetURL string) (utils.Response, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", targetURL, nil)
	if err != nil {
		return utils.Response{}, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return utils.Response{}, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return utils.Response{}, err
	}

	var responseData utils.Response

	if err := json.Unmarshal(body, &responseData); err != nil {
		return utils.Response{}, err
	}

	return responseData, nil
}
