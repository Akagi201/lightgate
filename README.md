# LightGate

A light weight API gateway

## Features

- [ ] Suport Admin API, and configs stores to [docker/libkv](https://github.com/docker/libkv).
- [ ] Support reversed proxy, via [vulcand/oxy](https://github.com/vulcand/oxy).
- [ ] Support standard http handler as plugin/middleware to add custom business logic or common staff.
- [ ] Support WebSocket proxy.
- [ ] Support HTTP proxy.
- [ ] A Cli tool to manage the admin staff.

## Middlewares
- [ ] JWT auth, support key server with [docker/libkv](https://github.com/docker/libkv).

## Inspired by

- [kong](https://github.com/Mashape/kong)
- [traefik](https://github.com/containous/traefik)
- [jwtproxy](https://github.com/coreos/jwtproxy)
- [vulcand](https://github.com/vulcand/vulcand)
