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
    url: "www.sample.com/results/"
    queries:
      - result/with?query=test
      - result/without
```
