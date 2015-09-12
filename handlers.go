package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")
}

func CourtIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(courts); err != nil {
		panic(err)
	}
}


func CourtCreate(w http.ResponseWriter, r *http.Request){

  var court Court
	var err error
  //Check if malicious user is trying to overload the server
  body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
  if err != nil{
    panic(err)
  }
  if err := r.Body.Close(); err != nil{
    panic(err)
  }
  if err := json.Unmarshal(body, &court); err != nil{
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")

		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}

  }
  if court, err = RepoCreateCourt(court); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}
  w.Header().Set("Content-Type", "application/json; charset=UTF-8")
  w.WriteHeader(http.StatusCreated)
  if err := json.NewEncoder(w).Encode(court); err != nil {
    panic(err)
  }


}

func CourtUpdate(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	var courtId int
	var err error

	if courtId, err = strconv.Atoi(vars["courtId"]); err != nil {
		panic(err)
	}

	var court Court
	//Check if malicious user is trying to overload the server
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &court); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}


  if err := RepoUpdateCourt(&court, courtId); err != nil{
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusNotFound) //TODO: Use right status code
    if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
      panic(err)
    }
  }
  w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK) // TODO: Use right status code

	if err := json.NewEncoder(w).Encode(court); err != nil {
		panic(err)
	}

}


func CourtDelete(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
  var courtId int
  var err error

  if courtId, err = strconv.Atoi(vars["courtId"]); err != nil{
    panic(err)
  }

	if err := RepoDeleteCourt(courtId); err!= nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusNotFound) //TODO: Use right status code
    if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
      panic(err)
    }
	} else {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusOK, Text: "Item Deleted"}); err != nil {
      panic(err)
    }
	}


}

func CourtShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var courtId int
	var err error

	if courtId, err = strconv.Atoi(vars["courtId"]); err != nil {
		panic(err)
	}

	court := RepoFindCourt(courtId)
	if court.Id > 0 {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(court); err != nil {
			panic(err)
		}
		return
	}

	// If we didn't find it, 404
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
		panic(err)
	}

}

/*
Example of JSON data structure
{
  "id":"1",
  "key":"value",
  "key2":["value1","value2"]
}
*/

/*
Test with this curl command:
curl -H "Content-Type: application/json" -d '{"name":"New Todo"}' http://localhost:8080/todos
*/
