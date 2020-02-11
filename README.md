# Sider-OG [WIP]

## SKETCH

Idea is to build a distribuited system for session management, similar to sudo temporary permission.

The server component is delegated to store HASHES in memory, with its TTL.
Client first set starting session token, then checks if session is active.
