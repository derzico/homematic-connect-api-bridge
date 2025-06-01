# Loxone ↔ Homematic IP Bridge

Diese Go-Anwendung stellt eine Middleware bereit, die HTTP-Befehle von Loxone entgegennimmt und über WebSocket an eine Homematic IP Zentrale weiterleitet.

## Features

- Leichte Ansteuerung von Homematic-Geräten per HTTP
- Läuft als eigenständiges Binary oder in Docker
- Minimaler Wartungsaufwand (nur 1 externe Abhängigkeit)

## API

GET /setSwitch?device=<deviceId>&state=on|off


## Konfiguration

Passe in `main.go` folgende Zeilen an:

```go
const (
    wsURL     = "ws://<homematic-ip>:<port>/api/ws"
    token     = "<dein_token>"
)

## Kompilieren

go build -o homematic-bridge

## Docker

docker build -t homematic-bridge .
docker run -p 8080:8080 --env HOMEMATIC_URL=... --env HOMEMATIC_TOKEN=... homematic-bridge
