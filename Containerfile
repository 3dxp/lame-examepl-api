FROM scratch
ARG APP_PATH="./simple-example-api"
USER 1000
COPY  x$APP_PATH /opt/app
EXPOSE 8080/tcp
ENTRYPOINT ["/opt/app"]