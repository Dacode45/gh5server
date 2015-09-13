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
func GetCourtByAddress(lat, lon float64) (Municipality, Court){

	cmd := exec.Command("python", "court_locator.py", strconv.FormatFloat(lat, 'g', 1, 64), strconv.FormatFloat(lon, 'g', 1, 64))
	out, err := cmd.Output()

	if err != nil{
		fmt.Printf("Could Not Run Python Script, err: %v", err)
		return Municipality{}, Court{}
	}
	log.Printf("Court @ %s", out)
	delim := ":^)"
	outArr := strings.Split(string(out), delim)
	MunicipalityJSON, CourtJSON := []byte(outArr[0]), []byte(outArr[1])
	mPython := Municipality{}
	cPython := Court{}

	if err = json.Unmarshal(MunicipalityJSON, &mPython); err == nil {
		mPython.Name = mPython.MUNICIPALITY
		mPython.Initialized = true
	}
	if err = json.Unmarshal(CourtJSON, &cPython); err == nil {
		if cPython.Address != ""{
			cPython.Initialized = true
		}
	}
	return mPython, cPython
}


	//Call Python script
	//Get Json stuff

func RepoFindCourtByAddress(lat, lon float64) Court{
	//range on an array index, object
	mPython, cPython := GetCourtByAddress(lat, lon)
	if !cPython.Initialized || !mPython.Initialized {
		return Court{}
	}
	var matched bool
	for i, muni := range gMunicipalities {
		if BareBones(muni.Name) == BareBones(mPython.Name) {
			matched = true
			mPython.Id = muni.Id
			gMunicipalities[i] = mPython
		}
	}
	if (!matched){
		mPython.Id = len(gMunicipalities) + 1
		gMunicipalities = append(gMunicipalities, mPython)
	}
	cPython.Muni = &mPython
	matched = false
	for i, court := range courts {
		if BareBones(court.Address) == BareBones(cPython.Address) {
			matched = true
			cPython.Id = court.Id
			courts[i] = cPython
		}
	}
	if (!matched){
		cPython.Id = len(courts) + 1
		courts = append(courts, cPython)
	}
	return cPython

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
