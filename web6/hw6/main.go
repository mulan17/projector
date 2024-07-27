// Створити інтерфейс «Публічний транспорт», який має методи «Приймати пасажирів» та «Висаджувати пасажирів», 
// і реалізувати його для типів «Автобус», «Потяг», «Літак». 
// Створити тип «Маршрут», який містить список транспортних засобів, які необхідні для проходження по заданому маршруту. 
// Тип «Маршрут» має мати методи «Додавати транспортний засіб до маршруту» та «Показувати список транспортних засобів на маршруті». 
// Тепер цей маршрут мусить пройти ваш подорожувальник («Пасажир») із виводом його подорожі на екран. 
// Файли різних груп об‘єктів зберігати в різних пакетах. 



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
