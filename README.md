# Your Turn To Roll

YTTR (Your turn to roll) is a utility tool for each DnD dungeon master (DM).
It allows the host (ie. the DM) to set up an initiative order for an upcoming fight and allow their players to join the
session to view the current and upcoming characters.

We all love pen and paper, but a little bit of technical support is useful especially in big fights when the players ask
themselves 'when is my turn?', 'did you skip me?'

## Execution

```text
task: Available tasks for this project:
* build:frontend:        Executes 'npm run dev' to create a production build of the frontend
* build:server:          Compiles the server with an embedded frontend into a platform specific binary
* build:server:pi:       Compiles the server with an embedded frontend into a binary for raspberry pi OS 32bit
* deploy:pi:             Builds a binary for raspberry pi OS 32bit and stops the setup service 'yttr' from systemd, copies the new binary to the raspberry pi under 'raspberrypi.local' and starts the service again
```

### Raspberry Pi

#### Setup the YTTR service with systemd

Create a service file:

```shell
vim /lib/systemd/system/yttr.service
```

with the following content:

```
[Unit]
Description=Your Turn To Roll
After=multi-user.target

[Service]
Type=simple
ExecStart=/home/pi/web

[Install]
WantedBy=multi-user.target
```

After that you need to update the systemd daemon and enable the service:

```shell
sudo systemctl daemon-reload
sudo systemctl enable sample.service
```

and now reboot:

```shell
sudo reboot
```

View logs via:
```shell
journalctl -e -u yttr
```

## Development

Run the frontend via:

```shell
cd frontend
npm run dev
```

Run the server via:

```shell
go run cmd/main.go
```

Or use the run configurations for intellij within this repository.

### Local environment

create a `.env.local` file in the `frontend` folder with the following content:

```text
VITE_API_URL=http://{your-internal-ip}:8081/api
VITE_WS_API_URL=ws://{your-internal-ip}:8081/api
```

## Building

Either build the docker image or use `task build:server` which will generate a binary called `web` which contains the
server and the frontend. The server is available under the `8081` port.
