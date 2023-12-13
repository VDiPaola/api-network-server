package models

type NodeResponseType struct {
	Status     int32
	StatusText string
	JSON       interface{}
	Text       string
}
