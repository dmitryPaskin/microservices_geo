FROM mirror.gcr.io/golang:alpine as builder

WORKDIR /app


COPY . .

#RUN go mod download
RUN go mod tidy

RUN go build -o ./main ./cmd/main.go

FROM mirror.gcr.io/alpine:latest

COPY --from=builder /app/main .
COPY --from=builder /app/.env .

EXPOSE 50051

CMD ["./main"]