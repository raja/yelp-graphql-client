package yelp

// Business represents a Yelp business
type Business struct {
	Name         string
	ID           string
	Alias        string
	Is_Claimed   bool
	Rating       float64
	Review_Count int
	Phone        string
	Price        string
	URL          string
	Categories   []Categories
	Photos       []string
	Location     Location
	Coordinates  Coordinates
	Reviews      []Reviews
	Hours        []Hours
}

// Location
type Location struct {
	Address1          string
	Address2          string
	Address3          string
	City              string
	State             string
	Zip_Code          string
	Country           string
	Formatted_Address string
}

type Coordinates struct {
	Latitude  float64
	Longitude float64
}

type Categories struct {
	Title string
	Alias string
}

type Reviews struct {
	ID           string
	Text         string
	Rating       int
	Time_Created string
	URL          string
	User         User
}

type User struct {
	ID          string
	Profile_Url string
	Name        string
	Image_Url   string
}

// Hours returns information describing the business's operating times.
type Hours struct {
	Hours_Type  string
	Is_Open_Now bool
	Open        []OpenHours
}

type OpenHours struct {
	Is_Overnight bool
	Day          int //0..6 == mon...sun
	Start        string
	End          string
}
