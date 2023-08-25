# Your Turn To Roll

YTTR (Your turn to roll) is a utility tool for each DnD dungeon master (DM).
It allows the host (ie. the DM) to set up an initiative order for an upcoming fight and allow their players to join the session to view the current and upcoming characters.

We all love pen and paper, but a little bit of technical support is useful especially in big fights when the players ask themselves 'when is my turn?', 'did you skip me?'

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

## Building

Either build the docker image or use `task build:server` which will generate a binary called `web` which contains the server and the frontend, which is being served by the server (on port `8081`).