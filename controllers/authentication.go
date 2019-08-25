package controllers

import (
	"class-review-backend/env"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

const TOKENINFO_ENDPOINT string = "https://oauth2.googleapis.com/tokeninfo?id_token="

// returns no error if the token is valid
func authenticate(IDToken string) (string, error) {
	_, retrieved := tokenCache.Get(IDToken)
	// if the user is in the cache, they've logged in within the last 30 mins and don't need to be revalidated
	if retrieved {
		return "", nil
	}
	resp, err := http.Get(TOKENINFO_ENDPOINT + IDToken)
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
	host, ok := token["hd"].(string)
	if !ok || host != "umich.edu" {
		return "", errors.New("invalid host -- user must be logged in with a University of Michigan email")
	}
	email, ok := token["email"].(string)
	if !ok {
		return "", errors.New("error casting email to string")
	}
	uniqname := strings.TrimSuffix(email, "@umich.edu")
	return uniqname, nil
}

func doesIdMatch(resourceId int64, uniqname string, targetTable string, targetIdField string, db *sqlx.DB) bool {
	var id int64
	rows, err := db.Query("SELECT id FROM ? WHERE ? = ?", targetTable, targetIdField, uniqname)
	if err != nil {
		log.Println(err.Error())
		return false
	}
	err = rows.Scan(&id)
	if err != nil {
		log.Println(err.Error())
		return false
	}
	return resourceId == id
}

func isAdmin(uniqname string, db *sqlx.DB) bool {
	// this probably isn't the best way to check, but I think it should work
	rows, err := db.Query("SELECT * FROM admins WHERE uniqname = ?", uniqname)
	if err != nil {
		log.Println(err.Error())
		return false
	}
	cols, err := rows.Columns()
	if err != nil || len(cols) == 0 {
		log.Println(err.Error())
		return false
	}
	return true
}

func hashUniqname(uniqname []byte) string {
	hash, err := bcrypt.GenerateFromPassword(uniqname, bcrypt.DefaultCost)
	if err != nil {
		log.Println(err.Error())
	}
	return string(hash)
}
