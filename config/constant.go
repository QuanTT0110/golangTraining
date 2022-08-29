package config

type Database struct {
	URI  string
	Name string
}

type ENV struct {
	AppPort    string
	Database   Database
	SigningKey string
}
