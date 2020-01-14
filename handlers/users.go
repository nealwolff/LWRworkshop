package handlers

import (
	"LWRworkshop/crud"
	"LWRworkshop/types"
	"encoding/json"
	"net/http"
)

//UserHandler handels the user endpoint
func UserHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		var user types.User
		json.NewDecoder(r.Body).Decode(&user)

		if user.Name == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("The json was malformed"))
			return
		}

		ret, err := crud.Insert("users", user, w)
		if err != nil {
			return
		}

		json.NewEncoder(w).Encode(ret)

	}
}
