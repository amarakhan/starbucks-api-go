package model

type Food struct {
	ID			int64	`json:"id"`
	Name 		string	`json:"name"`
	Price 		int		`json:"price"`
	DateAdded	string	`json:"dateAdded"`
}

type User struct {
	ID			int64	`json:"id"`
	FirstName	string	`json:"firstName"`
	LastName	string	`json:"lastName"`
	Email		string	`json:"email"`
	Staff		int		`json:"staff"`
	Address1	string	`json:"address1"`
	Address2	string	`json:"address2"`
	Zip			string	`json:"zip"`
	State		string	`json:"state"`
	Country		string	`json:"country"`
	AddDate		string	`json:"addDate"`
	ModDate		string	`json:"modDate"`
}
