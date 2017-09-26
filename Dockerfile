FROM alpine:3.6

ARG app_env
ENV APP_ENV $app_env

COPY ./maas_api /app/maas_api
RUN chmod +x /app/maas_api
# Install beego and the bee dev tool
WORKDIR /app 

# Expose the application on port 8080
EXPOSE 8080

# Set the entry point of the container to the bee command that runs the
# application and watches for changes
ENTRYPOINT /app/maas_api



