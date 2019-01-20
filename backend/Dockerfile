FROM golang AS builder
COPY . /go/src/NP2OP/backend/
WORKDIR /go/src/NP2OP/backend/
RUN go get . && go build .

FROM alpine
COPY --from=builder /go/src/NP2OP/backend/backend /app
CMD ./app/backend