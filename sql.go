package aefire

type DatabaseParam struct {
	DriverName   string
	Url          string
	MaxOpenConns int
	MaxIdleConns int
}
