package controllers

import (
	"class-review-backend/env"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

const TOKENINFO_ENDPOINT string = "https://oauth2.googleapis.com/tokeninfo?id_token="

// returns no error if the token is valid
func authenticate(IDToken string) error {
	val, retrieved := tokenCache.Get(IDToken)
	// if the user is in the cache, they've been active in the last hour and don't need to be revalidated
	if retrieved {
		// refresh the token expiration time
		tokenCache.Replace(IDToken, val, time.Hour)
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
	host, ok := token["hd"].(string)
	if !ok || host != "umich.edu" {
		return errors.New("invalid host -- user must be logged in with a University of Michigan email")
	}
	return nil
}
