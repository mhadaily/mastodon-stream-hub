FROM golang:1.21.6 as builder

RUN apt-get -y update && apt-get install -y gcc

WORKDIR /build

COPY . .

RUN go mod download

ARG project
RUN echo "Building $project..."
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o server ./cmd/${project}

FROM scratch
WORKDIR /app
COPY --from=builder /build/server /server

# Run the executable
CMD ["./server"]