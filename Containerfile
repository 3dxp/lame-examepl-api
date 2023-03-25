FROM scratch
USER 1000
COPY simple-example-api /opt/app
EXPOSE 8080/tcp
ENTRYPOINT ["/opt/app"]