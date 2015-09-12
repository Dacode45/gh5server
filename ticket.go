package main
import (
  "fmt"
)
type Ticket struct{
  TIs TicketInfos `json:"tis"`
  UserID int `json:"userid"`
  Id int `json:"id"`
  Muni Municipality `json:"municipality"`
}

type TicketInfo struct{
  Violation string `json:"violation"`
  Fine string `json:"fine"`
  CourtCost string `json:"courtcost"`
  Total string `json:"total"`
  Id int `json:"id"`
}

//validate ticket
const(
  TICKET_FIELD_NUMS = 3
  ERROR_MISSING_UserID = "USERID"
  ERROR_MISSING_Id = "ID"
  ERROR_MISSING_Muni = "MUNI"
)

func (t *Ticket) Validate () error{
  errors := make([]string, TICKET_FIELD_NUMS)
  if t.UserID < 0{
    errors = append(errors, ERROR_MISSING_UserID)
  }
  if t.Id < 0{
    errors = append(errors, ERROR_MISSING_Id)
  }
  if t.Muni.Id > 0{
    errors = append(errors, ERROR_MISSING_Muni)
  }

  if (len(errors) == 0){
    return nil
  } else {
    return fmt.Errorf("Missing the follwing field(s): %v", errors)
  }

}

//if TicketInfo exists, update it; else, append it.
func (t *Ticket) addCitation(ti TicketInfo) TicketInfos{
  for index, i := range t.TIs{
    if i.Id == ti.Id{
      t.TIs[index] = ti
      return t.TIs
    }
  }
  t.TIs = append(t.TIs, ti)
  return t.TIs
}

// func (t *Ticket) String() string{
//
// }
type Tickets []Ticket
type TicketInfos []TicketInfo
