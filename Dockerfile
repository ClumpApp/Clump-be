##
## Build
##
FROM golang:1.17.3 AS build

WORKDIR /go/src/app
COPY . .

RUN go build -ldflags="-s" -o /clump 

##
## Containerize
##
FROM ubuntu

WORKDIR /
COPY --from=build /clump /clump

EXPOSE 8080

ENTRYPOINT ["/clump"]
