FROM golang:latest

WORKDIR /go/src/app
COPY . .
RUN go build .
ENTRYPOINT ["go", "run", "."] 
