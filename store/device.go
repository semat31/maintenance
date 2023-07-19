// Package store provides types and methods for
// describing a unit of the production line
package store


// Device describes the unit of the production line
type Device struct {
	Code          string
	Name          string
	Categorie     string
	Serial        string
	Commissioning string
	model         string
	port          string
	aggregat      string
}


