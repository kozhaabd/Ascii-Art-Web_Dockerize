FROM golang:alpine AS builder
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o app .

FROM alpine
WORKDIR /app
COPY --from=builder /app .
EXPOSE 3000
CMD ["./app"]
