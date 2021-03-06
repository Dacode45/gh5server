package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	

	//Route of Court
  Route{
    "CourtIndex",
    "GET",
    "/courts",
    CourtIndex,
  },
	Route{
    //TODO: Add security so only authorized useser can create a court otherwise return htp status 500
    "CourtCreate",
    "POST",
    "/courts",
    CourtCreate,
  },
  Route{
    "CourtUpdate",
    "PUT",
    "/courts/{courtId}",
    CourtUpdate,
  },
  Route{
    //TODO: Add securty so only authorized users can delete a court otherwise return http status 500
    "CourtDelete",
    "DELETE",
    "/courts/{courtId}",
    CourtDelete,
  },
  Route{
    "CourtShow",
    "GET",
    "/courts/{courtId}",
    CourtShow,
  },

	//Route of Ticket
	Route{
		"TicketIndex",
		"GET",
		"/tickets",
		TicketIndex,
	},
	Route{
		"TicketShow",
		"GET",
		"/tickets/{ticketId}",
		TicketShow,
	},
	Route{
		"TicketCreate",
		"CREATE",
		"/tickets",
		TicketCreate,
	},
	Route{
		"TicketUpdate",
		"UPDATE",
		"/tickets/{ticketId}",
		TicketUpdate,
	},
	Route{
		"TicketDelete",
		"DELETE",
		"/tickets/{ticketId}",
		TicketDelete,
	},

//need work on TicketShow
	Route{
		"TicketShow",
		"GET",
		"/tickets/{ticketId}",
		TicketShow,
	},

	//Municipalities
	Route{
		"MunicipalityIndex",
		"GET",
		"/municipal/{lat}/{lon}",
		MunicipalIndex,
	},
	Route{
		"MunicipalityShow",
		"GET",
		"/municipal/{mId}",
		MunicipalShow,
	},

}
