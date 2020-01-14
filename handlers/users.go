package handlers

import (
	"LWRworkshop/crud"
	"LWRworkshop/types"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
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

	if r.Method == http.MethodGet {
		params := mux.Vars(r)
		ID := params["id"]
		LoanID := params["loanId"]

		rawUser, err := crud.GetOne("users", ID, w)
		if err != nil {
			return
		}

		user := types.User{}
		json.Unmarshal(rawUser, &user)

		userLoan := user.Loan[LoanID]
		asched := userLoan.CalculateASchedule()

		aschedRaw, _ := json.MarshalIndent(asched, "", "   ")
		w.Write(aschedRaw)

	}
}
