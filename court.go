package main

import (
  "fmt"
)

type Court struct{
  Id int                `json:"id"`
  Muni *Municipality   `json:"municipality"`
  PhoneNumber string    `json:"phone"`
  OpenTimes string      `json:"open_times"`
  Description string    `json"description"`
  PaymentTypes []string `json:"payment_types"`

  //Fields from python object
  Initialized bool //if gotten by python
  CourtId int `json:"court_id"`
	City string `json:"city"`
	Address string `json:"address"`
	County string `json:"county"`
  State string `json:"state"`
	ZipCode int `json:"zip_code"`
  X float64 `json:"x"`
  Y float64 `json:"y"`
	Level string `json:"level"`
  WebSite string `json:"web"`
}

var DEFAULT_COURT = Court{
  Id:0,
  City:"St. Louis",
  Address:"Multiple Addresses; see website",
  County:"St. Louis",
  State:"MO",
  ZipCode:0,
  X:-360,
  Y:-360,
  Level:"county",
  PhoneNumber:"(314) 615-8760",
  WebSite:"http://www.stlouisco.com/LawandPublicSafety/MunicipalCourts/LocationsandHours",
}
const(
  COURT_FIELD_NUMS = 7
  ERROR_MISSING_ID = "ID"
  ERROR_MISSING_Address = "ADDRESS"
  ERROR_MISSING_Municipality = "MUNICIPALITY"
  ERROR_MISSING_PhoneNumber = "PHONE NUMBER"
  ERROR_MISSING_OpenTimes = "OPEN TIMES"
  ERROR_MISSING_Description = "DESCRIPTION"
  ERROR_MISSING_PaymentTypes = "PAYMENT TYPES"
)

func (c *Court) Validate () error{
  errors := make([]string, COURT_FIELD_NUMS)
  if c.Id <= 0 {
    errors = append(errors, ERROR_MISSING_ID)
  }
  if len(c.Address) == 0{
    errors = append(errors, ERROR_MISSING_Address)
  }
  if c.Muni != nil{
    errors = append(errors, ERROR_MISSING_Municipality)
  }
  if len(c.PhoneNumber) == 0{
    errors = append(errors, ERROR_MISSING_PhoneNumber)
  }
  if len(c.OpenTimes) == 0{
    errors = append(errors, ERROR_MISSING_OpenTimes)
  }
  if len(c.Description) == 0{
    errors = append(errors, ERROR_MISSING_Description)
  }
  if c.PaymentTypes == nil || len(c.PaymentTypes) == 0{
    errors = append(errors, ERROR_MISSING_PaymentTypes)
  }else{
    for _, payment := range c.PaymentTypes {
      if len(payment) == 0{
        errors = append(errors, ERROR_MISSING_PaymentTypes)
      }
    }
  }

  if (len(errors) == 0){
    return nil
  } else {
    return fmt.Errorf("Missing the follwing field(s): %v", errors)
  }

}

// func (c *Court) String() string{
//
// }"T""

type Courts []Court
