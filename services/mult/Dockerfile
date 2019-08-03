FROM golang:1.11.1
ENV GOPATH="/go"
RUN mkdir -p /demo
COPY * /demo/
WORKDIR /demo
RUN go build -o svr
CMD ["./svr"]    