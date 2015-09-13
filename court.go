package main

import (
  "fmt"
)

type Court struct{
  Id int                `json:"id"`
  Muni *Municipality   `json:"municipality"`
  PhoneNumber string    `json:"phone_number"`
  OpenTimes string      `json:"open_times"`
  Description string    `json"description"`
  PaymentTypes []string `json:"payment_types"`

  //Fields from python object
  Initialized bool //if gotten by python
  Municipali string //name of Municipality
	Address string `json:"address"`
	City string `json:"city"`
	State string `json:"state"`
	Zip_Code string `json:"zip_code"`
	Transparen string
	SymbolID string
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
