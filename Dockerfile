FROM golang:latest as build
RUN mkdir /magazines
ADD . /magazines/
WORKDIR /magazines
RUN go build -o main .

FROM debian:jessie-slim

RUN groupadd -g 999 magazines \
    && useradd -r -u 999 -g magazines magazines \
    && mkdir -p /magazines/logs \
    && chown -R magazines:magazines /magazines
USER magazines
WORKDIR /magazines

COPY --from=build /magazines/main .
COPY --from=build /magazines/config.json .

EXPOSE 8088
ENTRYPOINT [ "./main" ]


