package helpers

import (
	"encoding/json"
	"fmt"
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

	var data map[string]interface{}

	if err := json.Unmarshal(body, &data); err != nil {
		return utils.Response{}, err
	}

	features := data["features"].([]interface{})

	var result string

	if len(features) > 0 {
		place_name := features[0].(map[string]any)["place_name"]
		s, ok := place_name.(string)
		if ok {
			result = s
		}

		fmt.Println("place name:", place_name)
	} else {
		fmt.Println("not found")
	}

	return utils.Response{
		Place: result,
	}, nil
}
