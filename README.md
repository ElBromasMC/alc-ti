# go-webserver-template
All Last Computing webpage

## Local environment

### Prerequisites
* Go
* Node and npm
* PostgreSQL
* [Air](https://github.com/cosmtrek/air#installation)
* [Templ](https://templ.guide/quick-start/installation)
* inotify-tools

### Install build dependencies
```shell
$ npm install
```

### .env file example
```
ENV=development
PORT=8080
REL=1
```

### Load env variables
```shell
$ set -a
$ source .env
$ set +a
```

### Live reload
```shell
$ make live
```

## Docker

### Prerequisites
* [Traefik](https://doc.traefik.io/traefik/getting-started/quick-start/)

### Docker compose .env file example
```
ENV=production
PORT=8080
REL=1
WEBSERVER_HOSTNAME=domain.tld
```
