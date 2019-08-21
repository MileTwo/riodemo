FROM golang:1.12.6-alpine3.9 as builder
COPY main.go .
RUN go build -o /app main.go

FROM alpine:3.9
CMD ["./app"]
EXPOSE 80
COPY --from=builder /app .