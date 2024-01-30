FROM golang:alpine
LABEL authors="alemaoprochnow"
WORKDIR /app
COPY go.mod .
RUN go mod download
COPY . .
RUN go build -o ./build/app
EXPOSE 80
