# react frontend builder
FROM node:14-alpine as uibuilder
WORKDIR /src
COPY ui .
RUN npm install && npm run build

# go backend builder
FROM golang:1.16 as gobuilder
WORKDIR /go/src/app
COPY . .
COPY --from=uibuilder /src/build ui/build
RUN go build -o ./bin/shellme ./cmd/shellme



# Final docker image
FROM ubuntu:20.04
WORKDIR /app
COPY --from=gobuilder /go/src/app/bin/shellme .
EXPOSE 8000

ENV GIN_MODE=release
CMD ["/app/shellme"]