FROM golang:1.24-alpine AS builder

RUN mkdir -p /home/backend
WORKDIR /home/backend

COPY ./go.mod ./
RUN go mod download
COPY ./ ./

RUN go build ./cmd/main.go

FROM scratch
COPY --from=builder /home/backend/main .

CMD ["./main", "--persist"]
