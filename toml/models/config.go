package models

// Config Models
type Config struct {
	Database database
}
type database struct {
	Server   string
	Port     string
	Database string
	User     string
	Password string
}
