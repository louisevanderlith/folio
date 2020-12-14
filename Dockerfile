FROM scratch

COPY cmd/cmd .

EXPOSE 8090

ENTRYPOINT [ "./cmd" ]