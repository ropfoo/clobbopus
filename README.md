<p align="center">
<img src="https://assets.pokemon.com/assets/cms2/img/pokedex/full/852.png" height="150" />
</p>
STILL WIP
Tool for downloading and serving html files.
Mainly useful for acting as a mock server for scraping services.

## Features

- download external pages by params
- serve them on a local server

## The Problem

Web scraping can be harmful when things go wrong.
When developing a web scraper it is not uncommon to send tons of requests to a server that is not under your control.
This might be because of a bug or just the nature of development espacially in a team.
Instead of flooding external servers with requests in development - we should flood our own.

## Setup

Create a config yml file called `clobbopus.yml`

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
