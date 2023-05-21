<p align="center">
<img src="https://assets.pokemon.com/assets/cms2/img/pokedex/full/852.png" height="150" />
</p>
STILL WIP
Tool for downloading and serving html files.
Mainly useful for acting as a mock server for scraping services.

## Motivation

Web scraping can be harmful when things go wrong. When developing a web scraper it is not uncommon to send tons of requests to a server that is not under your control. This might be because of a bug or just the nature of development, especially in a team. Instead of flooding external servers with requests in development - we should flood our own.

## Features

- download external pages by params
- serve them on a local server

## 🐙 Usage

Run the follwing commands in order

download and install

```bash
curl -L https://github.com/ropfoo/clobbopus/releases/download/Latest/install.sh -o install.sh && bash install.sh
```

create a config yml file called `clobbopus.yml`

```yml
port: 3000 # default
dist: pages # default
domains:
  sample:
    url: "www.sample.com/results/"
    params:
      - result/with?query=test&page=~1-7~ # range from page=1 to page=7
      - result/without
```

get the initial page data

```bash
./clobbopus_data
```

run the server

```bash
./clobbopus_server
```

### 🐳 Docker

`Dockerfile`

```Dockerfile
FROM alpine:latest
RUN apk update
RUN apk add bash
RUN apk add curl
WORKDIR /app
RUN curl -L https://github.com/ropfoo/clobbopus/releases/download/Latest/install.sh -o install.sh && bash install.sh
# update with your path
COPY ./clobbopus/clobbopus.yml .
CMD ./clobbopus_data; ./clobbopus_server
```

`docker-compose.yml`

```yml
version: "3.8"
services:
  clobbopus:
    build:
      context: .
      dockerfile: ./clobbopus/Dockerfile
    ports:
      - 3000:3000
    volumes:
      - ./clobbopus:/app/clobbopus
```
