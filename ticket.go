package main

import (
  "time"
  "fmt"
)
const (
  CONT_FOR_PAYMENT  = "CONT FOR PAYMENT"
  FTA_WARRANT_ISSUED  = "FTA WARRANT ISSUED"
  DISMISS_WITHOUT_COSTS  = "DISMISS WITHOUT COSTS"
  CLOSED  = "CLOSED"
)
const (
  EXPIRED_LICENSE_PLATES_TAGS  = "Expired License Plates (Tags)"
  FAILURE_TO_OBEY_ELECTRIC_SIGNAL = "Failure to Obey Electric Signal"
  FAILURE_TO_YIELD = "Failure to Yield"
  IMPROPER_PASSING  = "Improper Passing"
  NO_BRAKE_LIGHTS  ="No Brake Lights"
  NO_DRIVERS_LICENSE  ="No Driver's License"
  NO_INSPECTION_STICKER  = "No Inspection Sticker"
  NO_INSURANCE_NO_COMPLIANCE  ="No Insurance [no compliance]"
  NO_LICENSE_PLATES  ="No License Plates"
  PARKING_IN_FIRE_ZONE  ="Parking in Fire Zone"
  PROHIBITED_UTURN  = "Prohibited U-Turn"


)
type Ticket struct{
  //TIs TicketInfos `json:"tis"`
  UserID int `json:"userid"`
  Id int `json:"id"`
  CitationNumber int `json:"citation_number"`
  ViolationNumber string `json:"violation_number"`
  Violation string `json:"violation"`
  WarrentStatus bool `json:"warrent_status"`
  WarrentNumber string `json:"warrent_number"`
  Status string `json:"status"`
  StatusDate time.Time `json:StatusDate`
  Muni Municipality `json:"municipality"`
  FineAmount string `json:"fine_amount"`
  CourtCost string `json:"court_cost"`
}

/*type TicketInfo struct{

  Id int `json:"id"`
}*/

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
/*func (t *Ticket) addCitation(ti TicketInfo) TicketInfos{
  for index, i := range t.TIs{
    if i.Id == ti.Id{
      t.TIs[index] = ti
      return t.TIs
    }
  }
  t.TIs = append(t.TIs, ti)
  return t.TIs
}*/

// func (t *Ticket) String() string{
//
// }
type Tickets []Ticket
