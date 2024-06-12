package transport

type PublicTransport interface {
    GetPassengers(passengerNumber int)
    PushPassengers(passengerNumber int)
    TransportName() string
}
