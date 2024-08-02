FROM golang

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

WORKDIR /app/cmd

RUN go build -o /app/makeup ./main.go

CMD ["/app/makeup"]