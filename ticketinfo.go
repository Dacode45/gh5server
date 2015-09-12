package main

var gTicketInfos = TicketInfos{ //same as var ticketInfos = []TicketInfo{

  //MOVING TRAFFIC VIOLATION SCHEDULE
  TicketInfo{
    "Courner Cutting to Avoid Signal",
    "$70.50",
  },
  TicketInfo{
    "Driving on Closed Road",
    "$70.50",
  },
  TicketInfo{
    "Driving Over Curb",
    "$70.50",
  },
  TicketInfo{
    "Driving Wrong Side of Road",
    "$70.50",
  },
  TicketInfo{
    "Driving Wrong Way (One Way Street)",
    "$70.50",
  },
  TicketInfo{
    "Expired Operator’s License",
    "$70.50",
  },
  TicketInfo{
    "Failure to Dim Lights",
    "$70.50",
  },
  TicketInfo{
    "Failure to Keep to the Right",
    "$70.50",
  },
  TicketInfo{
    "Failure to Obey Electric Signal",
    "$70.50",
  },
  TicketInfo{
    "Failure to Obey Stop Sign",
    "$70.50",
  },
  TicketInfo{
    "Failure to Signal",
    "$70.50",
  },
  TicketInfo{
    "Failure to Sound Horn",
    "$70.50",
  },
  TicketInfo{
    "Failure to Yield",
    "$70.50",
  },
  TicketInfo{
    "Failure to Yield to Emergency Vehicle",
    "$70.50",
  },
  TicketInfo{
    "Following Too Closely",
    "$70.50",
  },
  TicketInfo{
    "Impeding/Obstructing Traffic Movement",
    "$70.50",
  },
  TicketInfo{
    "Improper Backing",
    "$70.50",
  },
  TicketInfo{
    "Improper Lane Usage",
    "$70.50",
  },
  TicketInfo{
    "Improper Passing",
    "$70.50",
  },
  TicketInfo{
    "Improper/Prohibited Turn",
    "$70.50",
  },
  TicketInfo{
    "No Through Traffic",
    "$70.50",
  },
  TicketInfo{
    "Prohibited U­Turn",
    "$70.50",
  },
  TicketInfo{
    "Traffic Turn Signal Violation",
    "$70.50",
  },

  //NON­ MOVING TRAFFIC VIOLATION SCHEDULE
  TicketInfo{
    "Blocking Driveway or Alley",
    "$50.50",
  },
  TicketInfo{
    "Child Restraint Seat Violation",
    "$24.50",
  },
  TicketInfo{
    "Excessive Vehicle Noise",
    "$50.50",
  },
  TicketInfo{
    "Expired License Plates (Tags)",
    "$50.50",
  },
  TicketInfo{
    "Failure to Dim Headlights",
    "$50.50",
  },
  TicketInfo{
    "Failure to Secure a Load",
    "$50.50",
  },
  TicketInfo{
    "Fake Temporary Tags",
    "$50.50",
  },
  TicketInfo{
    "Fictitious License Tabs",
    "$50.50",
  },
  TicketInfo{
    "Improper Exhaust or Muffler",
    "$50.50",
  },
  TicketInfo{
    "Improper Registration (Failure to Register)",
    "$50.50",
  },
  TicketInfo{
    "License Plate not Illuminated",
    "$50.50",
  },
  TicketInfo{
    "No Brake Lights",
    "$50.50",
  },
  TicketInfo{
    "No Headlight",
    "$50.50",
  },
  TicketInfo{
    "No Inspection Sticker",
    "$50.50",
  },
  TicketInfo{
    "No License Plates",
    "$50.50",
  },
  TicketInfo{
    "No Motorcycle Helmet",
    "$50.50",
  },
  TicketInfo{
    "No Taillights",
    "$50.50",
  },
  TicketInfo{
    "One Headlight and/or Taillight",
    "$50.50",
  },
  TicketInfo{
    "Operated MV w/Defective Parts",
    "$50.50",
  },
  TicketInfo{
    "Parking at Fire Hydrant",
    "$50.50",
  },
  TicketInfo{
    "Parking in Fire Zone",
    "$50.50",
  },
  TicketInfo{
    "Parking in Handicapped Zone",
    "$50.50",
  },
  TicketInfo{
    "All other Parking Violations",
    "$50.50",
  },
  TicketInfo{
    "Seat Belt Violation",
    "$10.00",
  },
  TicketInfo{
    "Tinted Windows",
    "$50.50",
  },

  //ORDINANCE VIOLATIONS SCHEDULE
  TicketInfo{
    "Abandoned Auto",
    "$50.50",
  },
  TicketInfo{
    "Dog at Large (not restrained)",
    "$50.50",
  },
  TicketInfo{
    "No Dog Tags",
    "$50.50",
  },
}

func init(){
  for i, t := range gTicketInfos{
    t.Id = i
  }
}
