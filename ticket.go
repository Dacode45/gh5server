package main

type Ticket struct{
  TIs TicketInfos `json:"tis"`
  UserID int `json:"userid"`
  Id int `json:"id"`
}

type TicketInfo struct{
  Violation string `json:"violation"`
  Fine string `json:"fine"`
  CourtCost string `json:"courtcost"`
  Total string `json:"total"`
  Id int `json:"id"`
}

//if TicketInfo exists, update it; else, append it.
func (t *Ticket) addCitation(ti TicketInfo){
  haveTicketInfo bool = false
  for index, i := range t.TIs{
    if i.Id == ti.Id{
      haveTicketInfo = true
    }
  }
  if haveTicketInfo == true{
    t.TIs[index] = ti
    return t.TIs
  }else{
    t.TIs = append(t.TIs, ti)
    return t.TIs
  }
}

// func (t *Ticket) String() string{
//
// }
type Tickets []Ticket
type TicketInfos []TicketInfo
