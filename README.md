# Sider-OG [WIP]

## Project description

Idea is to build a distribuited system for session management, similar to sudo temporary permission.

The server component is delegated to store HASHES in memory, with its TTL.
Client first set starting session token, then checks if session is active.

### Calls

```bash
# query start session
http://127.0.0.1:8080/start_session?token=0f39F48J938JF2D834DNCSDR4

# output: OK

# query check session
http://127.0.0.1:8080/check_session?token=0f39F48J938JF2D834DNCSDR4

# output: true/false
```

### Unittest

```bash
> go test -v -cover ./...
[...]
```
