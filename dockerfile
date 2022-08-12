FROM golang:1.19-alpine as dev

WORKDIR /work

FROM golang:1.19-alpine as build

WORKDIR /llm
COPY ./llm/* /llm/
RUN go build -o llm

FROM alpine as runtime
COPY --from=build /llm/llm /usr/local/bin/llm
COPY ./llm/orders.json /