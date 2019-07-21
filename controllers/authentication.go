package controllers

import (
	"class-review-backend/env"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

const TOKENINFO_ENDPOINT string = "https://oauth2.googleapis.com/tokeninfo?id_token="

// returns no error if the token is valid
func authenticate(IDToken string) error {
	_, retrieved := tokenCache.Get(IDToken)
	if retrieved {
		return nil
	}
	resp, err := http.Get(TOKENINFO_ENDPOINT + IDToken)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return errors.New(resp.Status)
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	var token map[string]interface{}
	// unpack response body from byte[] into a map
	json.Unmarshal(bodyBytes, &token)

	// I don't think we'll ever reach this but go makes you assert type
	// of an interface before returning
	aud, ok := token["aud"].(string)
	if !ok {
		return errors.New("aud claim's underlying type is not a string")
	}
	ClientId := env.Variables.ClientId
	if aud != ClientId {
		return errors.New("application client Id and aud claim " + aud + " do not match")
	}

	// TODO: Authenticate aud (make sure aud is in the app's client IDs)
	// for reference, this is an example of an aud claim:
	// 407408718192.apps.googleusercontent.com

	// once a user is validated, return their id (contained in the sub claim), and no error
	_, valid := token["sub"].(string)
	if !valid {
		return errors.New("sub claim's underlying type is not a string")
	}
	return nil
}
