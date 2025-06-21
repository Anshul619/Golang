# Parking Lot

The Parking Lot is a classic LLD (Low-Level Design) interview question. 

It evaluates your ability to:
- Design systems using OOP principles
- Use interfaces, structs, and concurrency primitives (in Go)
- Handle state management, modularity, and extensibility

# Typical Requirements (Interview Variant)
- Multiple parking floors
- Each floor has multiple parking slots
- Support for different vehicle types (car, bike, truck)

# Basic operations
- Park(vehicle): Assign nearest available slot
- Unpark(ticketID): Free slot
- DisplayAvailableSlots()
- GetTicket(vehicle): Return a ticket with info

# Optional
- Payment calculation (based on duration)
- Entry/exit gates
- Concurrent vehicle entry/exits