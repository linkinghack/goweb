FROM golang:1.14-alpine as builder
WORKDIR /app
COPY . /app
RUN go build -o goweb

FROM alpine
COPY --from=builder /app/goweb /app/goweb
EXPOSE 8088
CMD ["/app/goweb"]