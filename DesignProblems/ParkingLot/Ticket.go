package parkinglot

import "time"

type Ticket struct {
	TicketID  string
	Vehicle   Vehicle
	FloorID   int
	SlotID    int
	EntryTime time.Time
}
