package handlers

const (
	FLIGHT_READY     = "FLIGHT_READY"
	FLIGHT_IN_AIR    = "FLIGHT_IN_AIR"
	FLIGHT_COMPLETED = "FLIGHT_COMPLETED"
)

var Airports = map[string]string{
	"DEL": "Indira Gandhi International Airport, Delhi",
	"BOM": "Chhatrapati Shivaji Maharaj International Airport, Mumbai",
	"BLR": "Kempegowda International Airport, Bengaluru",
	"MAA": "Chennai International Airport",
	"HYD": "Rajiv Gandhi International Airport, Hyderabad",
	"CCU": "Netaji Subhas Chandra Bose International Airport, Kolkata",
	"GOI": "Goa International Airport, Dabolim",
	"IXC": "Shaheed Bhagat Singh International Airport, Chandigarh",
	"AMD": "Sardar Vallabhbhai Patel International Airport, Ahmedabad",
	"PNQ": "Pune Airport",
	"LKO": "Chaudhary Charan Singh International Airport, Lucknow",
	"JAI": "Jaipur International Airport",
	"COK": "Cochin International Airport",
	"TRV": "Trivandrum International Airport",
	"IXB": "Bagdogra Airport, Siliguri",
	"GAU": "Lokpriya Gopinath Bordoloi International Airport, Guwahati",
	"SXR": "Srinagar International Airport",
	"IXR": "Birsa Munda Airport, Ranchi",
	"BBI": "Biju Patnaik International Airport, Bhubaneswar",
	"NAG": "Dr. Babasaheb Ambedkar International Airport, Nagpur",
}
