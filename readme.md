# POC-GOLANG-IDESOFT

## Tabla de Contenido

- [Instalación](#instalación)
- [Uso](#uso)

## Instalación

Instrucciones para instalar el proyecto:

### Install go

```shell
$ brew install go
```

### Prepare ~/.bash_profile or ~/.zshrc or terminal session for Go development

```shell
export GOPATH=$HOME/go
export GOROOT=/usr/local/opt/go/libexec
export GOBIN=$GOPATH/bin
export PATH=$PATH:$GOPATH/bin
export PATH=$PATH:$GOROOT/bin
```

## Uso

Es una PoC donde se realiza la integración de un API de festivos con una capa de persistencia que se maneja mediante una lista sincronizada.

Para poder realizar la ejecución del proyecto podemos levantar el ambiente en docker o levantar el ambiente local en el puerto 8080
