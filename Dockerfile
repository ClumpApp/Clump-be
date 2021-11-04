##
## Build
##
FROM golang:1.17.2 AS build

WORKDIR /go/src/app
COPY . .

RUN go mod download
RUN go build -o /clump

##
## Deploy
##
FROM ubuntu AS deploy

WORKDIR /
COPY --from=build /clump /clump

EXPOSE 8080

ENTRYPOINT ["/clump"]
