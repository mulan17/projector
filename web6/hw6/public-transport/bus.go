package transport

import "fmt"

type Bus struct {
    name string
}

func NewBus(name string) *Bus {
    return &Bus{name: name}
}

func (b *Bus) GetPassengers(passengerNumber int) {
    fmt.Printf("Passenger %v get the bus %v\n", passengerNumber, b.name)
}

func (b *Bus) PushPassengers(passengerNumber int) {
    fmt.Printf("Passenger %v goes out the bus %v\n", passengerNumber, b.name)
}

func (b *Bus) TransportName() string {
    return b.name
}
