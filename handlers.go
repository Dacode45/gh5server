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

//handlers for municipality
func MunicipalIndex(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var lat float64
	var lon float64
	var err error

	if lat, err =  strconv.ParseFloat(vars["lat"], 64); err != nil{
		panic(err)
	}
	if lon, err = strconv.ParseFloat(vars["lon"], 64); err != nil{
		panic(err)
	}

	m := RepoFindCourtByAddress(lat, lon)
	if m.Id > 0 {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(m); err != nil {
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

func MunicipalShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var mId int
	var err error

	if mId, err = strconv.Atoi(vars["mId"]); err != nil {
		panic(err)
	}

	m := RepoFindMunicipality(mId)
	if m.Id > 0 {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(m); err != nil {
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

//handlers for court
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

//handlers for Ticket
func TicketIndex(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	driver_licenses, ok := params["driver_license"]
	if ok {
		driver_license := string(driver_licenses[0])
		var found bool = false
		for _, tic := range gTickets {
			if tic.DriverLicenseNumber == driver_license {
				found = true
				w.Header().Set("Content-Type", "application/json; charset=UTF-8")
				w.WriteHeader(http.StatusOK)
				if err := json.NewEncoder(w).Encode(tic); err != nil {
					panic(err)
				}
				break
			}
		}
		if (!found){
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusNotFound)
			if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
				panic(err)
			}
		}
		return
	} else {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(gTickets); err != nil {
			panic(err)
		}
	}
}

func TicketCreate(w http.ResponseWriter, r *http.Request){

  var ticket Ticket
	var err error
  //Check if malicious user is trying to overload the server
  body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
  if err != nil{
    panic(err)
  }
  if err := r.Body.Close(); err != nil{
    panic(err)
  }
  if err := json.Unmarshal(body, &ticket); err != nil{
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")

		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}

  }
  if ticket, err = RepoCreateTicket(ticket); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}
  w.Header().Set("Content-Type", "application/json; charset=UTF-8")
  w.WriteHeader(http.StatusCreated)
  if err := json.NewEncoder(w).Encode(ticket); err != nil {
    panic(err)
  }
}

func TicketUpdate(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	var ticketId int
	var err error

	if ticketId, err = strconv.Atoi(vars["ticketId"]); err != nil {
		panic(err)
	}

	var ticket Ticket
	//Check if malicious user is trying to overload the server
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &ticket); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

  if err := RepoUpdateTicket(&ticket, ticketId); err != nil{
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusNotFound) //TODO: Use right status code
    if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
      panic(err)
    }
  }
  w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK) // TODO: Use right status code

	if err := json.NewEncoder(w).Encode(ticket); err != nil {
		panic(err)
	}
}

func TicketDelete(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
  var ticketId int
  var err error

  if ticketId, err = strconv.Atoi(vars["ticketId"]); err != nil{
    panic(err)
  }

	if err := RepoDeleteTicket(ticketId); err!= nil {
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


func TicketShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var ticketId int
	var err error

	if ticketId, err = strconv.Atoi(vars["ticketId"]); err != nil {
		panic(err)
	}

	ticket := RepoFindTicket(ticketId)
	if ticket.Id > 0 {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(ticket); err != nil {
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
