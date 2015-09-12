package main

var gTicketInfos = TicketInfos{ //same as var ticketInfos = []TicketInfo{

  //MOVING TRAFFIC VIOLATION SCHEDULE
  TicketInfo{
    Violation: "Courner Cutting to Avoid Signal",
    Fine: "$70.50",
  },
  TicketInfo{
    Violation: "Driving on Closed Road",
    Fine: "$70.50",
  },
  TicketInfo{
    Violation: "Driving Over Curb",
    Fine: "$70.50",
  },
  TicketInfo{
    Violation: "Driving Wrong Side of Road",
    Fine: "$70.50",
  },
  TicketInfo{
    Violation: "Driving Wrong Way (One Way Street)",
    Fine: "$70.50",
  },
  TicketInfo{
    Violation: "Expired Operator’s License",
    Fine: "$70.50",
  },
  TicketInfo{
    Violation: "Failure to Dim Lights",
    Fine: "$70.50",
  },
  TicketInfo{
    Violation: "Failure to Keep to the Right",
    Fine: "$70.50",
  },
  TicketInfo{
    Violation: "Failure to Obey Electric Signal",
    Fine: "$70.50",
  },
  TicketInfo{
    Violation: "Failure to Obey Stop Sign",
    Fine: "$70.50",
  },
  TicketInfo{
    Violation: "Failure to Signal",
    Fine: "$70.50",
  },
  TicketInfo{
    Violation: "Failure to Sound Horn",
    Fine: "$70.50",
  },
  TicketInfo{
    Violation: "Failure to Yield",
    Fine: "$70.50",
  },
  TicketInfo{
    Violation: "Failure to Yield to Emergency Vehicle",
    Fine: "$70.50",
  },
  TicketInfo{
    Violation: "Following Too Closely",
    Fine: "$70.50",
  },
  TicketInfo{
    Violation: "Impeding/Obstructing Traffic Movement",
    Fine: "$70.50",
  },
  TicketInfo{
    Violation: "Improper Backing",
    Fine: "$70.50",
  },
  TicketInfo{
    Violation: "Improper Lane Usage",
    Fine: "$70.50",
  },
  TicketInfo{
    Violation: "Improper Passing",
    Fine: "$70.50",
  },
  TicketInfo{
    Violation: "Improper/Prohibited Turn",
    Fine: "$70.50",
  },
  TicketInfo{
    Violation: "No Through Traffic",
    Fine: "$70.50",
  },
  TicketInfo{
    Violation: "Prohibited U­Turn",
    Fine: "$70.50",
  },
  TicketInfo{
    Violation: "Traffic Turn Signal Violation",
    Fine: "$70.50",
  },

  //NON­ MOVING TRAFFIC VIOLATION SCHEDULE
  TicketInfo{
    Violation: "Blocking Driveway or Alley",
    Fine: "$50.50",
  },
  TicketInfo{
    Violation: "Child Restraint Seat Violation",
    Fine: "$24.50",
  },
  TicketInfo{
    Violation: "Excessive Vehicle Noise",
    Fine: "$50.50",
  },
  TicketInfo{
    Violation: "Expired License Plates (Tags)",
    Fine: "$50.50",
  },
  TicketInfo{
    Violation: "Failure to Dim Headlights",
    Fine: "$50.50",
  },
  TicketInfo{
    Violation: "Failure to Secure a Load",
    Fine: "$50.50",
  },
  TicketInfo{
    Violation: "Fake Temporary Tags",
    Fine: "$50.50",
  },
  TicketInfo{
    Violation: "Fictitious License Tabs",
    Fine: "$50.50",
  },
  TicketInfo{
    Violation: "Improper Exhaust or Muffler",
    Fine: "$50.50",
  },
  TicketInfo{
    Violation: "Improper Registration (Failure to Register)",
    Fine: "$50.50",
  },
  TicketInfo{
    Violation: "License Plate not Illuminated",
    Fine: "$50.50",
  },
  TicketInfo{
    Violation: "No Brake Lights",
    Fine: "$50.50",
  },
  TicketInfo{
    Violation: "No Headlight",
    Fine: "$50.50",
  },
  TicketInfo{
    Violation: "No Inspection Sticker",
    Fine: "$50.50",
  },
  TicketInfo{
    Violation: "No License Plates",
    Fine: "$50.50",
  },
  TicketInfo{
    Violation: "No Motorcycle Helmet",
    Fine: "$50.50",
  },
  TicketInfo{
    Violation: "No Taillights",
    Fine: "$50.50",
  },
  TicketInfo{
    Violation: "One Headlight and/or Taillight",
    Fine: "$50.50",
  },
  TicketInfo{
    Violation: "Operated MV w/Defective Parts",
    Fine: "$50.50",
  },
  TicketInfo{
    Violation: "Parking at Fire Hydrant",
    Fine: "$50.50",
  },
  TicketInfo{
    Violation: "Parking in Fire Zone",
    Fine: "$50.50",
  },
  TicketInfo{
    Violation: "Parking in Handicapped Zone",
    Fine: "$50.50",
  },
  TicketInfo{
    Violation: "All other Parking Violations",
    Fine: "$50.50",
  },
  TicketInfo{
    Violation: "Seat Belt Violation",
    Fine: "$10.00",
  },
  TicketInfo{
    Violation: "Tinted Windows",
    Fine: "$50.50",
  },

  //ORDINANCE VIOLATIONS SCHEDULE
  TicketInfo{
    Violation: "Abandoned Auto",
    Fine: "$50.50",
  },
  TicketInfo{
    Violation: "Dog at Large (not restrained)",
    Fine: "$50.50",
  },
  TicketInfo{
    Violation: "No Dog Tags",
    Fine: "$50.50",
  },
}

func init(){
  for i, t := range gTicketInfos{
    t.Id = i
  }
}
