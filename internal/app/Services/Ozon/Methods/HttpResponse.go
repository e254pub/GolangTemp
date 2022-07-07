package Methods

import (
	"bytes"
	"encoding/json"
	"errors"
	"main/internal/app/Services/Ozon/Structures"
	"net/http"
	"os"
)

func OzonResponse(data interface{}, method Structures.OzonMethod) (*http.Response, error) {
	auth := Structures.OzonAuth{ApiKey: os.Getenv("OZON_API_KEY"), ClientId: os.Getenv("OZON_CLIENT_ID")}
	payloadBytes, err := json.Marshal(data)
	if err != nil {
		return nil, errors.New("json parse error")
	}
	body := bytes.NewReader(payloadBytes)
	req, err := http.NewRequest(method.Method, "https://api-seller.ozon.ru"+method.Url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Api-Key", auth.ApiKey)
	req.Header.Set("Client-Id", auth.ClientId)

	resp, err := http.DefaultClient.Do(req)

	return resp, err
}
