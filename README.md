# Prototype with [OpenFGA](https://openfga.dev)

See [the entry in my wiki "Prototype OpenFGA with Recurse OAuth"](<~/tmpshare/personal-wiki-html/entries/html-wiki/stories/Prototype OpenFGA with Recurse OAuth.md>)

## Folder Structure

<dl>

<dt>[bin](./bin)</dt>
<dd>Scripts to execute</dd>

</dl>

## Tasks

[![xc compatible](https://xcfile.dev/badge.svg)](https://xcfile.dev)

### Setup

Pull down the CLI image, so future `docker run -it openfga/cli` calls will work instantly.

```sh
docker pull openfga/cli
```

### Run

`--ansi never` because `docker compose` aggressively tried to clear my Emcas `vterm` buffer and it was unreadable.

```sh
docker compose --ansi never up
```

### CLI_List_Stores

```sh
./bin/local-fga store list
```
