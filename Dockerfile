# go backend builder
FROM golang:1.16 as gobuilder
WORKDIR /go/src/app
COPY . .
RUN go build -o ./bin/shellme ./cmd/shellme

# react frontend builder
FROM node:14-alpine as uibuilder
WORKDIR /src
COPY ui .
RUN npm install && npm run build


# Final docker image
FROM ubuntu:20.04
WORKDIR /app
COPY --from=gobuilder /go/src/app/bin/shellme .
COPY --from=uibuilder /src/build ui/build
EXPOSE 8000

ENV PRODUCTION=1
ENV GIN_MODE=release

CMD ["/app/shellme"]