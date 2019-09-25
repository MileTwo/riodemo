#
# build
#
FROM golang:1.12.6-alpine3.9 as builder
COPY src/main.go .
RUN go build -o /app main.go

#
# package up only the app
#
FROM alpine:3.9
# set color from build-ars on docker build command

ARG COLOR
ENV COLOR ${COLOR:-yellow}

EXPOSE 80
COPY --from=builder /app .
CMD ["./app"]