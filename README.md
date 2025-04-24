# Docker Compose Setup: Go App + Redis

This repository contains a simple `docker-compose` setup to run a Go application alongside a Redis instance, both connected on the same Docker network.

---

## 🧱 Structure

```
├── Dockerfile           # Go app Dockerfile 
├── docker-compose.yml   # Multi-container config
├── main.go              # Go application
└── README.md            # You're here!
```

## 🚀 Quick Start

**Build and Run Containers**

```bash
docker-compose up --build
```

This will:
* Build the Go app container (from Dockerfile)
* Pull and run the official Redis image
* Launch both on a shared internal network


**Verify It's Working**

You should see logs from your app and Redis. The app includes a health check for Redis interaction.  You can check health of the system going to:
```
http://localhost:8080
```
