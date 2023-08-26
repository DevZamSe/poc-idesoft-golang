#####################################
#Configuracion imagen para compilación
#####################################
FROM devzamse/base-image:latest AS build-env

LABEL MAINTAINER = 'Devzamse'

#####################################
#Generar archivo UNIX
#####################################
RUN apk --no-cache add build-base git gcc
ENV D=/Go/src/test-idesoft
ADD . $D
RUN cd $D && go build -o msCoreServer -ldflags '-linkmode=external' cmd.go && cp msCoreServer /tmp/

#####################################
#Configuracion imagen final
#####################################
FROM devzamse/base-image:latest

RUN apk add --no-cache \python3 \py3-pip \&& pip3 install --upgrade pip \&& pip3 install --no-cache-dir \awscli \&& rm -rf /var/cache/apk/*

#####################################
#Exponer puerto HTTP
#####################################
ENV HTTP_PORT=8080
EXPOSE $HTTP_PORT

#####################################
#Cargar e iniciar archivo UNIX
#####################################
WORKDIR /app

COPY --from=build-env /tmp/msCoreServer /app/msCoreServer
COPY .env /app
CMD ["./msCoreServer"]