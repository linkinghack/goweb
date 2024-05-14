FROM golang:1.22.3-bullseye as builder
WORKDIR /app
COPY . /app
RUN go build -o goweb

FROM ubuntu
COPY --from=builder /app/goweb /app/goweb
EXPOSE 8088
CMD ["/app/goweb"]