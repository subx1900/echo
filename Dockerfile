# builder image
FROM golang:alpine as builder
RUN mkdir /build
ADD *.go /build/
WORKDIR /build
RUN CGO_ENABLED=0 GOOS=linux go build -a -o echo .


# generate clean, final image for end users
FROM alpine:3.12
COPY --from=builder /build/echo .

# executable
EXPOSE 3000
ENTRYPOINT [ "./echo" ]
