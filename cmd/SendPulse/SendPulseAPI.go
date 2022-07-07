package SendPulse

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

var grantType string
var clientID string
var clientSecret string
var accessToken string
var fromContact Sender

func initialize(rClientID string, rClientSecret string, rName string, rFromEmail string) {
	grantType = "client_credentials"
	clientID = rClientID
	clientSecret = rClientSecret
	fromContact = Sender{
		Name:  rName,
		Email: rFromEmail,
	}
}

func getKey() (string, error) {
	requestBody, err := json.Marshal(map[string]string{
		"grant_type":    grantType,
		"client_id":     clientID,
		"client_secret": clientSecret,
	})

	if err != nil {
		return "", errors.New("marshalling request payload gave an error")
	}

	resp, err := http.Post("https://api.sendpulse.com/oauth/access_token",
		"application/json",
		bytes.NewBuffer(requestBody))

	if err != nil {
		return "", errors.New("making the request gave an error")
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", errors.New("reading the response gave an error")
	}

	var response oauthTokenResponse
	err = json.Unmarshal([]byte(body), &response)

	if response.ErrorCode != 0 {
		return "", errors.New("sendPulse sent an error")
	}

	accessToken = response.AccessToken
	return response.AccessToken, nil
}

func sendEmail(template Template, subject string, to []Recipient) error {
	mailObj := emailArray{
		email{
			Template: template,
			Subject:  subject,
			From:     fromContact,
			To:       to,
		},
	}

	mailByteSlice, err := json.Marshal(mailObj)
	if err != nil {
		return errors.New("something wrong with email object -> string")
	}

	req, err := http.NewRequest(
		"POST",
		"https://api.sendpulse.com/smtp/emails",
		bytes.NewBuffer([]byte(mailByteSlice)),
	)

	if err != nil {
		return errors.New("something wrong with string -> request")
	}

	token, err := getKey()

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return errors.New("something wrong sending the email")
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	var emailResponse sendEmailResponse
	err = json.Unmarshal([]byte(body), &emailResponse)
	if err != nil {
		return err
	}
	if !emailResponse.Result {
		return errors.New("something went wrong at SendPulse")
	}

	return nil
}
