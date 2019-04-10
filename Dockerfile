#build stage
FROM golang:stretch AS builder
WORKDIR /go/src/app
COPY . .
ENV GO111MODULE=on
RUN go build -o server

#final stage
FROM gcr.io/distroless/base
COPY --from=builder /go/src/app/server /app
ENV PORT=${PORT}
ENTRYPOINT [ "/app" ]
CMD [ "/server" ]