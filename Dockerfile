FROM alpine:3.6

ARG app_env
ENV APP_ENV $app_env

COPY ./maas_server /app/maas_server
RUN chmod +x /app/maas_server
# Install beego and the bee dev tool
WORKDIR /app 

# Expose the application on port 8080
EXPOSE 8080

# Set the entry point of the container to the bee command that runs the
# application and watches for changes
ENTRYPOINT /app/maas_server


# FROM golang

# # Install beego and the bee dev tool
# RUN go get github.com/astaxie/beego && go get github.com/beego/bee
# RUN go get github.com/gorilla/schema && go get github.com/gorilla/mux

# # Expose the application on port 8080
# EXPOSE 8080

# # Set the entry point of the container to the bee command that runs the
# # application and watches for changes
# CMD ["bee", "run"]