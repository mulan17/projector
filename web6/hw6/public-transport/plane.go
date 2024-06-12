package transport

import "fmt"

type Plane struct {
    name string
}

func NewPlane(name string) *Plane {
    return &Plane{name: name}
}

func (p *Plane) GetPassengers(passengerNumber int) {
    fmt.Printf("Passenger %v get the plane %v\n", passengerNumber, p.name)
}

func (p *Plane) PushPassengers(passengerNumber int) {
    fmt.Printf("Passenger %v goes out the plane %v\n", passengerNumber, p.name)
}

func (p *Plane) TransportName() string {
    return p.name
}
