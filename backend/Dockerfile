FROM golang:alpine3.12 as builder
RUN mkdir /build 
ADD . /build/
WORKDIR /build 
RUN go get
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o main .

FROM scratch
COPY --from=builder /build/main /app/
WORKDIR /app
CMD ["./main","-addr", "0.0.0.0:8000"]
