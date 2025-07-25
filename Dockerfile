FROM golang:1.24-alpine AS builder

RUN mkdir -p /home/backend
WORKDIR /home/backend

COPY ./go.mod ./
RUN go mod download
COPY ./ ./

RUN go build ./cmd/main.go

FROM alpine
COPY --from=builder /home/backend/main .
RUN chmod +x ./main

CMD ["./main", "--persist"]
