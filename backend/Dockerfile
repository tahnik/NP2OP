FROM golang
COPY . /go/src/NP2OP/backend/
WORKDIR /go/src/NP2OP/backend/
RUN go get .
RUN go run main.go
