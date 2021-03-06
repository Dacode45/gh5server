package main

import (
"fmt"
"log"
"os/exec"
"strings"
"encoding/json"
"strconv"
)

var gCourtId int
var courts Courts

// Give us some seed data
func init() {
	  RepoCreateCourt(Court{})
}

//Takes In String returns one with all whitespace replaced and lower case
func BareBones(str string) string{
		return strings.ToLower(strings.Replace(str, " ", "", -1))
}
//Repo Court
func RepoFindCourt(id int) Court{
  //range on an array index, object
  for _, c := range courts{
    if c.Id == id{
      return c
    }
  }
  // return empty Court if not found
  return Court{}
}

//Id of Court and Municipality will be 0
func GetCourtByAddress(lat, lon float64) (Court){

	cmd := exec.Command("python", "court_locator.py", strconv.FormatFloat(lat, 'f', 6, 64), strconv.FormatFloat(lon, 'f', 6, 64))
	out, err := cmd.Output()
	//fmt.Printf("%d, %d", lat, lon)
	if err != nil{
		fmt.Printf("Could Not Run Python Script, err: %v", err)
		return Court{}
	}

	cPython := Court{}
	//log.Printf("Court @ %s", out)
	outArr := strings.Split(string(out), ":^)")
	if err = json.Unmarshal([]byte(outArr[1]), &cPython); err == nil {
	}
	if cPython.CourtId != 0{
		cPython.Initialized = true
	}
	//fmt.Printf("%v", cPython)
	return cPython
}


	//Call Python script
	//Get Json stuff

func RepoFindCourtByAddress(lat, lon float64) Court{
	//range on an array index, object
	cPython := GetCourtByAddress(lat, lon)
	if !cPython.Initialized {
		return DEFAULT_COURT
	}
	var matched bool
	for i, court := range courts {
		if court.Id == cPython.CourtId {
			matched = true
			cPython = courts[i]
		}
	}
	if (!matched){
		cPython.Id = cPython.CourtId
		courts = append(courts, cPython)
	}
	matched = false
	for _, muni := range gMunicipalities {
		//fmt.Printf("%v vs %v = %v\n",BareBones(muni.Name),BareBones(cPython.City), strings.EqualFold(BareBones(muni.Name), BareBones(cPython.City)))
		if strings.EqualFold(BareBones(muni.Name), BareBones(cPython.City)) == true {
			matched = true
			cPython.Muni = &muni
			break
		}
	}
	if (!matched){
		gMunicipalityCounter += 1
		mid := gMunicipalityCounter
		muni := Municipality{Id: mid, Name: cPython.City, MunicipalCourtWebsite: cPython.WebSite}
		gMunicipalities = append(gMunicipalities, muni)
		cPython.Muni = &gMunicipalities[mid]
	}

	log.Printf("Got Court\n")
	return cPython

}

func RepoFindTicketByDriverLicenseNumber(driver_license_number string) Ticket{
	//TODO: Handlers should really get things from global arrays. Create a function in repo.go that returns a ticket from drivers license
	for _, tic := range gTickets {
		if strings.EqualFold(tic.DriverLicenseNumber, driver_license_number) {
				return tic
		}
	}
	return Ticket{}

}

func RepoFindMunicipality(mId int) Municipality{
	//range on an array index, object
  for _, m := range gMunicipalities{
    if m.Id == mId{
      return m
    }
  }
  // return empty Ticket if not found
  return Municipality{}
}


func RepoUpdateCourt(new_court *Court, id int) error{
  new_court.Id = id
	if err := new_court.Validate(); err != nil{
		return fmt.Errorf("Bad court: %v", err.Error())
	}
  for i, c := range courts{
    if c.Id == id{
      courts[i] = *new_court
			return nil
    }
  }
  return fmt.Errorf("Could not find a Court with id of %v to update", id)
}

func RepoCreateCourt(c Court) (Court, error){
  gCourtId += 1
  c.Id = gCourtId
	//Caution: Possibility of data racing.
	if err := c.Validate(); err != nil{
		gCourtId -= 1
		return Court{}, fmt.Errorf("Bad court: %v", err.Error())
	}
  courts = append(courts, c)
  return c, nil
}

func RepoDeleteCourt(id int) error{
  for i, c := range courts{
    if c.Id == id{
      courts = append(courts[:i], courts[i+1:]...)
      return nil
    }
  }
  return fmt.Errorf("Could not find a Court with id of %v to delete", id)
}

//Repo Ticket
func RepoFindTicket(id int) Ticket{
  //range on an array index, object
  for _, t := range gTickets{
    if t.Id == id{
      return t
    }
  }
  // return empty Ticket if not found
  return Ticket{}
}

func RepoUpdateTicket(new_ticket *Ticket, id int) error{
  new_ticket.Id = id
	if err := new_ticket.Validate(); err != nil{
		return fmt.Errorf("Bad ticket: %v", err.Error())
	}
  for i, t := range gTickets{
    if t.Id == id{
      gTickets[i] = *new_ticket
			return nil
    }
  }
  return fmt.Errorf("Could not find a Ticket with id of %v to update", id)
}

func RepoCreateTicket(t Ticket) (Ticket, error){
  gCourtId += 1
  t.Id = gCourtId
	//Caution: Possibility of data racing.
	if err := t.Validate(); err != nil{
		gCourtId -= 1
		return Ticket{}, fmt.Errorf("Bad ticket: %v", err.Error())
	}
  gTickets = append(gTickets, t)
  return t, nil
}

func RepoDeleteTicket(id int) error{
  for i, t := range gTickets{
    if t.Id == id{
      gTickets = append(gTickets[:i], gTickets[i+1:]...)
      return nil
    }
  }
  return fmt.Errorf("Could not find a Ticket with id of %v to delete", id)
}
