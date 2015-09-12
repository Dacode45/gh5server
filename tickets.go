package main
import (
  "encoding/csv"
  "os"
  "strconv"
  "time"
  "fmt"
)
var gTickets = Tickets{

}
func initTickets() error {
  var err error
  csvfile, err := os.Open("violations.csv")

	if err != nil {
		return fmt.Errorf("Failed to Open violations.csv: %v", err.Error())
	}
  defer csvfile.Close()
	reader := csv.NewReader(csvfile)
	reader.FieldsPerRecord = 10
  rawCSVdata, err := reader.ReadAll()
  if err != nil {
		return err
	}
  gTickets = make(Tickets, len(rawCSVdata))
  for _, record := range(rawCSVdata){
    var tempTicket = Ticket{}
    if tempTicket.CitationNumber,err  = strconv.Atoi(record[1]); err != nil{
      return fmt.Errorf("Failed to read citation number: %v", err)
    }
    tempTicket.ViolationNumber = record[2]
    switch record[3] {
    case EXPIRED_LICENSE_PLATES_TAGS:
      tempTicket.Violation = EXPIRED_LICENSE_PLATES_TAGS
    case FAILURE_TO_OBEY_ELECTRIC_SIGNAL:
      tempTicket.Violation = FAILURE_TO_OBEY_ELECTRIC_SIGNAL
    case FAILURE_TO_YIELD:
      tempTicket.Violation = FAILURE_TO_YIELD
    case IMPROPER_PASSING:
      tempTicket.Violation = IMPROPER_PASSING
    case NO_BRAKE_LIGHTS:
      tempTicket.Violation = NO_BRAKE_LIGHTS
    case NO_DRIVERS_LICENSE:
      tempTicket.Violation = NO_DRIVERS_LICENSE
    case NO_INSPECTION_STICKER:
      tempTicket.Violation = NO_INSPECTION_STICKER
    case NO_INSURANCE_NO_COMPLIANCE:
      tempTicket.Violation = NO_INSURANCE_NO_COMPLIANCE
    case NO_LICENSE_PLATES:
      tempTicket.Violation = NO_LICENSE_PLATES
    case PARKING_IN_FIRE_ZONE:
      tempTicket.Violation = PARKING_IN_FIRE_ZONE
    case PROHIBITED_UTURN:
      tempTicket.Violation = PROHIBITED_UTURN
    }
    if tempTicket.WarrentStatus, err = strconv.ParseBool(record[4]); err != nil{
      return fmt.Errorf("Failed to read citation number: %v", err)
    }
    tempTicket.WarrentNumber = record[5]
    switch record[6] {
    case CONT_FOR_PAYMENT:
      tempTicket.Status = CONT_FOR_PAYMENT
    case FTA_WARRANT_ISSUED:
      tempTicket.Status = FTA_WARRANT_ISSUED
    case DISMISS_WITHOUT_COSTS:
      tempTicket.Status = DISMISS_WITHOUT_COSTS
    case CLOSED:
      tempTicket.Status = CLOSED
    }
    if tempTicket.StatusDate, err = time.Parse("1/2/2006", record[7]); err != nil{
      return fmt.Errorf("Failed to read status data: %v", err)
    }
    tempTicket.FineAmount = record[8]
    tempTicket.CourtCost = record[9]
    gTickets = append(gTickets, tempTicket)
  }
  return nil
}
