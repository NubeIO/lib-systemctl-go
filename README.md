# lib-systemctl-go

## install

clone the repo

```
go mod tidy
cd cmd
```

run as sudo

```
go build ctl.go && sudo ./ctl  --help
```

## state enabled and running

```
    "state": "enabled",
    "active_state": "active",
    "sub_state": "running",
```

## state enabled but has an error (trying to start)

```
    "state": "enabled",
    "active_state": "activating",
    "sub_state": "auto-restart",
```

## state enabled and stopped

```
    "state": "enabled",
    "active_state": "inactive",
    "sub_state": "dead",
```

## state disabled and running

```
    "state": "disabled",
    "active_state": "active",
    "sub_state": "running",
```

## state disabled and stopped

```
    "state": "disabled",
    "active_state": "inactive",
    "sub_state": "dead",
```

## docs

[CLI](docs/cmd.md)

[TESTS](docs/tests.md)