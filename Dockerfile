FROM golang:1.17-alpine AS build
WORKDIR /app

COPY go.* ./
RUN go mod download

COPY *.go ./
RUN go build -o be-auth


FROM alpine

COPY --from=build /app/be-auth ./

ENV PORT 80
ENV GIN_MODE release
CMD /be-auth
