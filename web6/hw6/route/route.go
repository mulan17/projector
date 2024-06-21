package route

import (
    "fmt"
)

type PublicTransport interface {
    GetPassengers(passengerNumber int)
    PushPassengers(passengerNumber int)
    TransportName() string
}

type Route struct {
    transports []PublicTransport
}

func NewRoute() *Route {
    return &Route{}
}

func (r *Route) AddTransport(t PublicTransport) {
    r.transports = append(r.transports, t)
}

func (r *Route) ShowTransports() {
    fmt.Println("Transports:")
    for _, t := range r.transports {
        fmt.Printf("%v\n", t.TransportName())
    }
}

func (r *Route) Travel(passengerNumber int) {
    for _, t := range r.transports {
        t.GetPassengers(passengerNumber)
        t.PushPassengers(passengerNumber)
    }
}
