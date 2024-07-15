FROM golang:alpine as builder

WORKDIR /usr/local/src

COPY ["go.mod", "go.sum", "./"]
RUN go mod download

COPY ./ ./

RUN go build -o ./bin/app ./main.go

FROM alpine AS runner

COPY --from=builder /usr/local/src/bin/app /app


RUN ls -la /app
WORKDIR /

CMD ["/app"]
