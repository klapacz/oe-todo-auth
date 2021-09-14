FROM golang:1.17-alpine AS build
WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . .
RUN go build  -o be-auth *.go


FROM alpine

COPY --from=build /app/be-auth ./

ENV PORT 80
ENV GIN_MODE release
CMD /be-auth
