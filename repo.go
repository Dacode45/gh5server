package main

import "fmt"

var currentId int
var courts Courts

// Give us some seed data
func init() {
	  RepoCreateCourt(Court{})
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

func GetMuncipalityByAddress(lat, lon) (Municipality, Court, error){

	cmd := exec.Command("municipal.py")
	out, err := cmd.Output()

	if err != nil{
		return Municipality{}, error
	}


	outArr := strings.Split(out, delim)
	MunicipalityJSON,


	//Call Python script
	//Get Json stuff
}

func RepoFindMunicipalityByAddress(lat, lon float) Municipality{
	//range on an array index, object
	mPython, cPython, err := GetMuncipalityByAddress(lat, lon)
	//check if this data matches municipality object if it doesn't add municipality to lis,
	//if it does, update it
	// return empty Ticket if not found
	return Municipality{}
}

func RepoFindMunicipality(mId) Municipality{
	//range on an array index, object
  for _, m := range gMunicipalities{
    if m.Id == id{
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
  currentId += 1
  c.Id = currentId
	//Caution: Possibility of data racing.
	if err := c.Validate(); err != nil{
		currentId -= 1
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
  currentId += 1
  t.Id = currentId
	//Caution: Possibility of data racing.
	if err := t.Validate(); err != nil{
		currentId -= 1
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
