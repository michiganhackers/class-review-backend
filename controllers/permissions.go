package controllers

const OWN = "OWN"
const ANY = "ANY"

var permissions = map[string]string{
	// no need to check permissions for public endpoints

	// Can anybody logged in post a new professor? I feel like this should be an admin endpoint, right?
	"POST /professor/names":          OWN,
	"PUT /professor/names":           OWN,
	"DELETE /professor/names":        OWN,
	"GET /professor/stats":           ANY,
	"GET /professor/stats/:uniqname": ANY,
}

// making this a map for future generalization
var ownerKeys = map[string]string{
	"reviews": "uniqname",
}
