package main

import (
    "hw6/route"
    "hw6/public-transport"
)

func main() {
    bus := transport.NewBus("Bohdan")
    train := transport.NewTrain("Ukrzaliznytsia")
    plane := transport.NewPlane("Mria")

    route := route.NewRoute()
    route.AddTransport(bus)
    route.AddTransport(train)
    route.AddTransport(plane)

    route.ShowTransports()
    route.Travel(1708)
}
