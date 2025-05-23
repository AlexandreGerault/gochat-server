FROM golang:1.24-alpine AS golang

WORKDIR /app
COPY . .

RUN go mod download
RUN go mod verify

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /server .

FROM gcr.io/distroless/static-debian11

COPY --from=golang /server .

EXPOSE 10000

CMD ["/server"]

