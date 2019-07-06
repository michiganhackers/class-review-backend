package controllers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

// returns [google id], [error message] if the token is valid
func authenticate(id_token string) (string, error) {
	resp, err := http.Get("https://oauth2.googleapis.com/tokeninfo?id_token=" + id_token)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != http.StatusOK {
		return "", errors.New(resp.Status)
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	var token map[string]interface{}
	// unpack response body as byte[] into a map
	json.Unmarshal(bodyBytes, &token)

	// I don't think we'll ever reach this but go makes you assert type
	// of an interface before returning
	aud, ok := token["aud"].(string)
	if !ok {
		return "", errors.New("token's underlying type is not a string")
	}

	// TODO: Authenticate aud (make sure aud is in the app's client IDs)
	// for reference, this is an example of an aud claim:
	// 407408718192.apps.googleusercontent.com

	return aud, nil
}
