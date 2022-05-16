## testing

- Add a new file in /tmp named `/tmp/rubix-updater.service`
- go build main.go && sudo ./main add /tmp/rubix-updater.service
- the command above copied the file too /lib/systemd/system/rubix-updater.service
- confirm its there
  -- `cat /lib/systemd/system/rubix-updater.service`

### add the tmp service file

```
sudo nano /tmp/rubix-updater.service
```

```
[Unit]
Description=nubeio-rubix-updater

[Service]
Type=simple
Restart=always
RestartSec=5s
WorkingDirectory=/home/aidan/rubix-updater/backend
ExecStart=/home/aidan/rubix-updater/backend/main

[Install]
WantedBy=multi-user.target
```

set the service name to make it easy when adding deleting

```
export NAME=rubix-updater.service
```

```
sudo systemctl start $NAME
sudo systemctl stop $NAME
sudo systemctl status $NAME
sudo systemctl disable $NAME
```

if you want to delete the service

```
sudo systemctl status $NAME
sudo rm /lib/systemd/system/$NAME
sudo rm /usr/lib/systemd/system/$NAME
sudo systemctl daemon-reload
sudo systemctl reset-failed
```

## for testing make a test server

```go
import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Welcome to Go and Gin!")
	})
	r.Run()
}
```
