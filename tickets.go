package main
import (
  "encoding/csv"
  "os"
  "strconv"
  "time"
  "fmt"
  "io/ioutil"
  "encoding/json"
)
var gTicketId = 1
var gTickets = Tickets{

}
func init(){
  if err := initTickets(); err != nil {
    panic(err)
  }

}
func initTickets() error {
  file, err := os.Open(DATABASE_CACHE_FILE)
  var cached = false
  if err == nil {
    defer file.Close()
    serilized, err := ioutil.ReadAll(file)
    if err = json.Unmarshal(serilized, gTickets); err == nil{
      cached = true
    }
  }
  if !cached{
    if err = loadViolations(); err != nil {
      return fmt.Errorf("Error Loading violations.csv: %v", err)
    }
    if err = loadCitations(); err != nil {
      return fmt.Errorf("Error Loading citations.csv: %v", err)
    }
  }
  return nil
}
func loadViolations() error {
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
  for i, record := range(rawCSVdata){
    if i == 0 {
      continue
    }
    var tempTicket = Ticket{}
    tempTicket.Id = gTicketId
    gTicketId += 1
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
    gTickets[i] = tempTicket
  }
  return nil
}
func loadCitations() error {
  //TODO: Use a more efficient data structure later
  var err error
  csvfile, err := os.Open("citations.csv")
  if err != nil {
    return fmt.Errorf("Failed to Open citations.csv: %v", err.Error())
  }
  defer csvfile.Close()
  reader := csv.NewReader(csvfile)
  reader.FieldsPerRecord = 13
  rawCSVdata, err := reader.ReadAll()
  if err != nil {
    return err
  }
  for j, record := range rawCSVdata{
    if j == 0 {
      continue
    }
    var citation_number int
    if citation_number,err  = strconv.Atoi(record[1]); err != nil{
      return fmt.Errorf("Failed to read citation number: %v", err)
    }
    for i, previous_record := range gTickets {
      if previous_record.CitationNumber == citation_number{
        gTickets[i].CitationDate, err = time.Parse("1/2/2006 15:04", record[2])
        gTickets[i].FirstName = record[3]
        gTickets[i].LastName = record[4]
        gTickets[i].DateOfBirth, err = time.Parse("1/2/2006 15:04", record[5])
        gTickets[i].DefendantAddress = record[6]
        gTickets[i].DefendantCity = record[7]
        gTickets[i].DefendantState = record[8]
        gTickets[i].DriverLicenseNumber = record[9]
        gTickets[i].CourtDate, err = time.Parse("1/2/2006 15:04", record[10])
        court, _ := RepoCreateCourt(Court{Address: record[12], City: record[11]})
        gTickets[i].Court = &court
      }
    }
  }
  return nil
}
