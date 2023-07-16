FROM golang:1.20-alpine AS build
WORKDIR /app
COPY go.* .
RUN go mod download && go mod verify
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go install -a -installsuffix cgo -ldflags '-extldflags "-static"' ./...

FROM scratch
COPY --from=build /go/bin/* /bin/
ENTRYPOINT ["/bin/tpl"]
