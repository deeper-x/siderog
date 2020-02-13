# Sider-OG [WIP]

## Project description

Idea is to build a distribuited system for session management, similar to sudo temporary permission.

The server component is delegated to store HASHES in memory, with its TTL.
Client first set starting session token, then checks if session is active.

### Calls

```bash
# start session
/start_session?token=0f39F48J938JF2D834DNCSDR4

# return: OK

# check session
/check_session?token=0f39F48J938JF2D834DNCSDR4

# return: true|false
```

### Unittest

```bash
> go test -v -cover ./...
[...]
```
