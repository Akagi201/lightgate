# LightGate

[![Build Status](https://travis-ci.org/Akagi201/lightgate.svg)](https://travis-ci.org/Akagi201/lightgate)
[![Coverage Status](https://coveralls.io/repos/github/Akagi201/lightgate/badge.svg?branch=master)](https://coveralls.io/github/Akagi201/lightgate?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/Akagi201/lightgate)](https://goreportcard.com/report/github.com/Akagi201/lightgate)
[![GoDoc](https://godoc.org/github.com/Akagi201/lightgate?status.svg)](https://godoc.org/github.com/Akagi201/lightgate)

A light weight API gateway.

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
