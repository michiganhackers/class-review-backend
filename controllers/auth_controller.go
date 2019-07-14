package controllers

import (
	"class-review-backend/env"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/thoas/go-funk"
)

func AuthenticationRequired(auths ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get("user")
		id_token := c.GetHeader("id_token") // not sure what we'll call the key, id_token seems fine
		if user == nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "user needs to be signed in to access this service"})
			c.Abort()
			return
		}
		if len(auths) != 0 {
			authType := session.Get("authType")
			if authType == nil || !funk.ContainsString(auths, authType.(string)) {
				c.JSON(http.StatusForbidden, gin.H{"error": "invalid request, restricted endpoint"})
				c.Abort()
				return
			}
		}
		// add session verification here, like checking if the user and authType
		// combination actually exists if necessary. Try adding caching this (redis)
		// since this middleware might be called a lot
		_, err := authenticate(id_token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "could not authenticate id token -- " + err.Error()})
			c.Abort()
			return
		}
		c.Next()
	}
}

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
	// unpack response body from byte[] into a map
	json.Unmarshal(bodyBytes, &token)

	// I don't think we'll ever reach this but go makes you assert type
	// of an interface before returning
	aud, ok := token["aud"].(string)
	if !ok {
		return "", errors.New("aud claim's underlying type is not a string")
	}
	ClientId := env.Variables.ClientId
	if aud != ClientId {
		return "", errors.New("application client Id and aud claim " + aud + " do not match")
	}

	// TODO: Authenticate aud (make sure aud is in the app's client IDs)
	// for reference, this is an example of an aud claim:
	// 407408718192.apps.googleusercontent.com

	// once a user is validated, return their id (contained in the sub claim), and no error
	sub, ok := token["sub"].(string)
	if !ok {
		return "", errors.New("sub claim's underlying type is not a string")
	}
	return sub, nil
}
