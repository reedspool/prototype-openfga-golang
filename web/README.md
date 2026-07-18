# Web server for OpenFGA prototype

## Tasks

[![xc compatible](https://xcfile.dev/badge.svg)](https://xcfile.dev)

### Run

```sh
go run main.go
```

### Test

```sh
go test ./...
```

### Test_clear

Utility for the below watch command

```sh
clear
xc test || cat <<-EOF
  !
  ! snuffing non-zero exit code $?, continuuing
  !
EOF
```

### Test_watch

Requires [`fswatch`](https://github.com/emcrisostomo/fswatch). Can test the name of the flag to put after `--event` in your system with `fswatch -x`.

```sh
xc Test_clear
fswatch --one-per-batch --recursive --exclude '.*.git' --exclude '.*.jj' --latency 0.8 --event-flags --event-flag-separator '%' --event Updated . | xargs -I{} xc Test_clear
```
