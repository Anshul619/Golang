package parkinglot

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

type Slot struct {
	ID      int
	IsFree  bool
	Vehicle *Vehicle
	Type    VehicleType
}

type Floor struct {
	ID    int
	Slots []*Slot
}

type ParkingLot struct {
	Floors []*Floor
	mu     sync.Mutex
}

func (pl *ParkingLot) Park(vehicle Vehicle) (*Ticket, error) {
	pl.mu.Lock()
	defer pl.mu.Unlock()

	for _, floor := range pl.Floors {
		for _, slot := range floor.Slots {
			if slot.IsFree && slot.Type == vehicle.Type {
				slot.IsFree = false
				slot.Vehicle = &vehicle
				ticket := &Ticket{
					TicketID:  generateTicketID(vehicle.Number, floor.ID, slot.ID),
					Vehicle:   vehicle,
					FloorID:   floor.ID,
					SlotID:    slot.ID,
					EntryTime: time.Now(),
				}
				return ticket, nil
			}
		}
	}
	return nil, errors.New("No available slots")
}

func (pl *ParkingLot) Unpark(ticketID string) error {
	pl.mu.Lock()
	defer pl.mu.Unlock()

	// Parse ticketID to extract floor and slot info (or use map for lookup)
	for _, floor := range pl.Floors {
		for _, slot := range floor.Slots {
			if !slot.IsFree && generateTicketID(slot.Vehicle.Number, floor.ID, slot.ID) == ticketID {
				slot.IsFree = true
				slot.Vehicle = nil
				return nil
			}
		}
	}
	return errors.New("Invalid ticket")
}

func generateTicketID(vehicleNumber string, floorID, slotID int) string {
	return fmt.Sprintf("T-%s-%d-%d", vehicleNumber, floorID, slotID)
}

func (pl *ParkingLot) DisplayAvailableSlots(vehicleType VehicleType) {
	for _, floor := range pl.Floors {
		fmt.Printf("Floor %d: ", floor.ID)
		for _, slot := range floor.Slots {
			if slot.IsFree && slot.Type == vehicleType {
				fmt.Printf("%d ", slot.ID)
			}
		}
		fmt.Println()
	}
}

func NewParkingLot(floors, slotsPerFloor int) *ParkingLot {
	pl := &ParkingLot{}

	for f := 0; f < floors; f++ {
		floor := &Floor{ID: f}
		for s := 0; s < slotsPerFloor; s++ {
			var vt VehicleType
			if s == 0 {
				vt = Truck
			} else if s <= 2 {
				vt = Bike
			} else {
				vt = Car
			}
			floor.Slots = append(floor.Slots, &Slot{
				ID:     s,
				IsFree: true,
				Type:   vt,
			})
		}
		pl.Floors = append(pl.Floors, floor)
	}

	return pl
}

func main() {
	pl := NewParkingLot(2, 6)

	v1 := Vehicle{Number: "KA-01-HH-1234", Type: Car}
	ticket1, _ := pl.Park(v1)
	fmt.Println("Ticket:", ticket1)

	pl.DisplayAvailableSlots(Car)

	pl.Unpark(ticket1.TicketID)
	fmt.Println("Unparked.")

	pl.DisplayAvailableSlots(Car)
}
