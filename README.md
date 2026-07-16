# Prototype with [OpenFGA](https://openfga.dev)

See [the entry in my wiki "Prototype OpenFGA with Recurse OAuth"](<~/tmpshare/personal-wiki-html/entries/html-wiki/stories/Prototype OpenFGA with Recurse OAuth.md>)

## Folder Structure

<dl>

<dt><a href="./bin">bin/</a></dt>
<dd>Scripts to execute</dd>

</dl>

## Tasks

[![xc compatible](https://xcfile.dev/badge.svg)](https://xcfile.dev)

### Setup_Local_Dev

Pull down the FGA image, so future `docker run -it openfga/cli` calls will work instantly.

```sh
docker pull openfga/cli
```

### Run

`--ansi never` because `docker compose` aggressively tried to clear my Emcas `vterm` buffer and it was unreadable.

```sh
docker compose --ansi never up
```

### FGA_List_Stores

Can use this if you lose the store ID

```sh
./bin/local-networked-fga store list
```

### FGA_Create_Store

```sh
./bin/local-networked-fga store create --name "FGA Demo Store"
echo "Grab the above ID and put it in ./local.config.fga.yaml's store-id: field"
```

### FGA_Create_Model

See the note in [./bin/local-networked-fga](./bin/local-fga) about `/devdir`. That's where the file is in the FGA's Docker container.

```sh
./bin/local-networked-fga model write --file=/devdir/model.fga
echo "Grab the above ID and put it in ./local.config.fga.yaml's model-id: field"
```

### FGA_Get_Model

Just to check it out and verify it's there.

```sh
./bin/local-networked-fga model get --format "fga" # Can also try format json
```

### FGA_Test_Model_Networked

**Either** test against the model currently on the server which the FGA targets **OR** test against the file directly, bypassing the server, but still use the Docker network.

Which of the above occurs depends on whether the test file has a `model:` or `model_file:` top-level field. If `test.fga.yaml` file includes `model:` or `model_file:` then it will skip the server and test directly simulating the model within the FGA. If the file does NOT contain either of those fields, then it will hit the server and test against the model loaded in there. See the [docs](https://github.com/openfga/cli/#command-11).

```sh
./bin/local-networked-fga model test --tests="/devdir/test.fga.yaml"
```

### FGA_Test_Model_Networked_Clear

Utility for the below watch command

```sh
clear
xc FGA_Test_Model_Networked || cat <<-EOF
  !
  ! snuffing non-zero exit code $?, continuuing
  !
EOF
```

### FGA_Test_Model_Networked_Watch

Watches for file changes and immediately re-runs tests

Requires [`fswatch`](https://github.com/emcrisostomo/fswatch). Can test the name of the flag to put after `--event` in your system with `fswatch -x`.

```sh
xc FGA_Test_Model_Networked_Clear
fswatch --one-per-batch --recursive --exclude '.*.git' --exclude '.*.jj' --latency 0.8 --event-flags --event-flag-separator '%' --include "test.fga.yaml" --include "model.fga" --event Updated . | xargs -I{} xc FGA_Test_Model_Networked_Clear
```

### FGA_Test_Model_Local

Test with no server using the model file as only and the simulated OpenFGA within the FGA. See the [docs](https://github.com/openfga/cli/#command-11).

The `test.fga.yaml` file must include a `model:` or `model_file:` for this (so uncomment it or add it)

```sh
./bin/local-fga model test --tests="/devdir/test.fga.yaml"
```
