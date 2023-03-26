FROM scratch
ARG APP_PATH="./simple-example-api"
USER 1000
COPY  $APP_PATH /opt/app
EXPOSE 8080/tcp
ENTRYPOINT ["/opt/app"]