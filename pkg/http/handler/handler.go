package http

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/katzien/go-structure-examples/domain/listing"
	"github.com/katzien/go-structure-examples/domain/reviewing"
)

func NewRouter(a auth.Service) http.Handler {
	router := mux.NewRouter()
	router.POST("/join", addUser(a))
	return router
}


func addUser(s adding.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var newUser auth.User
		err := json.NewDecoder(r.Body).Decode(&newBeer)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		s.AddUser(newUser)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"msg": "New user added"})
	}
}

