FROM alpine:latest

COPY folio .
COPY conf conf

EXPOSE 8090

ENTRYPOINT [ "./folio" ]
