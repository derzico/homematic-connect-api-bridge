# Homematic Connect API Bridge (Loxone ↔ Homematic)

Ein kleiner Go-basierter HTTP-zu-WebSocket-Proxy, der Loxone HTTP-Befehle entgegennimmt und an eine Homematic IP Zentrale weiterleitet.

## Features

- Einfache Steuerung von Homematic-Geräten über Loxone
- Minimaler Ressourcenverbrauch
- Bereit für Docker
- Konfigurierbar über Umgebungsvariablen

---

## API

```http
GET /setSwitch?device=<deviceId>&state=on|off

Beispiel:
GET http://<bridge-ip>:8080/setSwitch?device=HmIP-ABC1234567&state=on
