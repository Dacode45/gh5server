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

func RepoUpdateCourt(new_court Court, id int) error{
  new_court.Id = id
  for _, c := range courts{
    if c.Id == id{
      c = new_court
    }
  }
  return fmt.Errorf("Could not find a Court with id of %v to update", id)
}

func RepoCreateCourt(c Court) Court{
  currentId += 1
  c.Id = currentId
  courts = append(courts, c)
  return c
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
