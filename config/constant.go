package config

type Database struct {
	URI  string
	Name string
}

// ENV .env struct
type ENV struct {
	// App port
	AppPort string

	// Database
	Database Database
}
