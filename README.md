<p align="center">
<img src="https://assets.pokemon.com/assets/cms2/img/pokedex/full/852.png" height="150" />
</p>
WIP tool for downloading and serving html files.
Mainly useful for acting as a mock server for scraping services.

## Setup

Create a config yml file called `clobbopus.yml`

```yml
domains:
  sample:
    params:
      page: 1-20
    urls:
      - "http://localhost:8080/sample1"
      - "http://localhost:8080/sample2"
```
