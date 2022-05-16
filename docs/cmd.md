## cmd usage

## builder

```
go build ctl.go &&  ./ctl  builder --write=true --builder-path=/home/aidan --builder-name=test
```

this will build a new systemd file named `test` with a path of `/home/aidan`

outputs

```
------------------------------
[Unit]
Description=aidans-service Service
After=network.target
[Service]
User=aidan
WorkingDirectory=/home/aidan
ExecStart=/usr/bin/python3 something.py
Restart=always
[Install]
WantedBy=multi-user.target
------------------------------
------------------------------
build and add new file here: /home/aidan/test.service
------------------------------
```

## service manager (systemd)

### add a new service

steps

- `add` a new service
- `start` the service
- `enable` the service for auto restart

```
go build ctl.go  && sudo ./ctl service  --add=true --path=/home/aidan/test.service
```

#### remove a service

```
go build ctl.go  && sudo ./ctl service  --remove=true --service=test
```

### status

```
go build ctl.go && sudo ./ctl  service --service=test --status=true
```
