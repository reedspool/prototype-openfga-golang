# Prototype with [OpenFGA](https://openfga.dev)

See [the entry in my wiki "Prototype OpenFGA with Recurse OAuth"](<~/tmpshare/personal-wiki-html/entries/html-wiki/stories/Prototype OpenFGA with Recurse OAuth.md>)

## Folder Structure

<dl>

<dt><a href="./bin">bin/</a></dt>
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

Can use this if you lose the store ID

```sh
./bin/local-fga store list
```

### CLI_Create_Store

```sh
./bin/local-fga store create --name "FGA Demo Store"
echo "Grab the above ID and put it in ./local.fga.yaml's store-id: field"
```

### CLI_Create_Model

See the note in [./bin/local-fga](./bin/local-fga) about `/devdir`. That's where the file is in the CLI's Docker container.

```sh
./bin/local-fga model write --file=/devdir/model.fga
echo "Grab the above ID and put it in ./local.fga.yaml's model-id: field"
```

### CLI_Get_Model

Just to check it out and verify it's there.

```sh
./bin/local-fga model get --format "fga" # Can also try format json
```
