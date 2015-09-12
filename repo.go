package main

import "fmt"

var currentId int
var courts Courts

// Give us some seed data
func init() {
	  RepoCreateCourt(Court{})
}

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
