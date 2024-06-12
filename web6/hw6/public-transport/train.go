package transport

import "fmt"

type Train struct {
    name string
}

func NewTrain(name string) *Train {
    return &Train{name: name}
}

func (t *Train) GetPassengers(passengerNumber int) {
    fmt.Printf("Passenger %v get the train %v\n", passengerNumber, t.name)
}

func (t *Train) PushPassengers(passengerNumber int) {
    fmt.Printf("Passenger %v goes out the train %v\n", passengerNumber, t.name)
}

func (t *Train) TransportName() string {
    return t.name
}
