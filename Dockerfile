FROM golang:1.19-alpine as build

RUN mkdir /app
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . /app/
RUN apk add gcc g++
RUN CGO_ENABLED=1 GOOS=linux go build -a -ldflags '-linkmode external -extldflags "-static"' -o main

FROM scratch
COPY --from=build /app/main .
ENV DB_NAME=dev.db
EXPOSE 8080
CMD [ "./main" ]