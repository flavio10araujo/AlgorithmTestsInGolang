package main

type RequestType int

const (
	PickupUp RequestType = iota
	PickupDown
	Destination
)

type Request struct {
	Floor int
	Type  RequestType
}
