# Sider-OG [WIP]

![Go](https://github.com/deeper-x/siderog/workflows/Go/badge.svg)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/90c2bae02c784850b0961fbdc7acd9c9)](https://www.codacy.com/manual/deeper-x/siderog?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=deeper-x/siderog&amp;utm_campaign=Badge_Grade)
[![GoDoc](https://godoc.org/github.com/deeper-x/siderog?status.svg)](https://godoc.org/github.com/deeper-x/siderog)

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
# create role
/admin/new_role?value=938579384579348579347

# return: 938579384579348579347
# Description: register role 938579384579348579347 in order to allow to consume /start_session call

# start session
/start_session?role=938579384579348579347

# return: 16b9ee3151ee76fdf5af5c509f9c208865e5a398a660167b64554c4e51211b9
# Description: This is the client's token

# check session
/check_session?token=16b9ee3151ee76fdf5af5c509f9c208865e5a398a660167b64554c4e51211b9e

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

### TLS setup

```bash
$ mkdir -p tls/cert tls/key
# private key (.key)
$ openssl genrsa -out ./tls/key/server.key 2048
# Key considerations for algorithm "ECDSA" ≥ secp384r1
# List ECDSA the supported curves (openssl ecparam -list_curves)
$ openssl ecparam -genkey -name secp384r1 -out ./tls/key/server.key
# self-signed(x509) public key (PEM-encodings .pem|.crt) based on the private (.key)
$ openssl req -new -x509 -sha256 -key ./tls/key/server.key -out ./tls/cert/server.crt -days 3650


```
