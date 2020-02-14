# Sider-OG [WIP]

![Go](https://github.com/deeper-x/siderog/workflows/Go/badge.svg)

## Project description

Http based sudo-like service.
Some random description notes: 

- Client is registered to be allowed to consume the service 

- Registered client asks for session starting, server allows it.

- Client asks if session is active, server replies.

- Server generates unique static server identity, basing on machine ID.

- Server component is delegated to store HASHES in memory, with its TTL.

- Token is machine-unique (machine ID), unpredictable (hashed 256sum).

### Calls

```bash
# start session
/start_session

# return: 16b9ee3151ee76fdf5af5c509f9c208865e5a398a660167b64554c4e51211b9
# Description: This is the client's token

# check session
/check_session?token=16b9ee3151ee76fdf5af5c509f9c208865e5a398a660167b64554c4e51211b9

# return: [true|false]
# Description: Client session is up/down
```

### Unittest

```bash
> go test -v -cover ./...
[...]
```

### Build

```bash
> export GOPATH=${HOME}/go
> export GOBIN=${GOPATH}/bin
> export PATH=${PATH}:${GOBIN}
> export GO111MODULE=on
....
> go build -o ${GOBIN}
> siderog
> Server running...
```

### Make

```bash
# Install
$ make install

# Run
$ make run

# Stop
$ make stop
```
