# Docker Compose Setup: Go App + Redis

This repository contains a simple `docker-compose` setup to run a Go application alongside a Redis instance, both connected on the same Docker network.

It was designed to be deployed using docker-compose on a developer's machine and using kubernetes.

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

## Design Considerations
This repo is meant to build a single docker image (apple) for use in kubernetes.  

That being said, a docker-compose yaml is included to allow developers to stand up the service with the attendant dependencies (in this case, redis) for development purposes.  

It is not intended for docker-compose to play any role in production.   This was chosen as kubernetes is too cumbersome for most development, but docker-compose isn't really a great container orchestration solution.

Any kubernetes cluster will need to include a central redis deployment for app pods to communicate with.

## CI/CD
On PR creation, PR commit, and merge events, this repository will build go, run go tests, build the docker image and push the image to docker hub.  The Docker user name and password are configured as github secrets.

Tests can be found in the main_test.go file.

The built docker images are pushed to the alfonsol/apple-alpha registry in docker hub. 

For more details on that look in the .github/workflows directory.

## Deployment
For developer deployment see the quick start section.   For production this repo was built with kubernetes in mind, though any container orchestration stac
k would be fine.   

Included in the repo is "deployment.yaml" which can be applied to a kubernetes cluter via kubectl or any number of automations.

**I suggest we make use of Github actions to do actual deployments. An action would be triggered when we create a release,but await manual approval before proceeding.**

The service provides a `/health` endpoint for kubernetes to check on for load balancing and routing between running instances.  It currently only checks for a redis connection.

**Important Note**

* **Within Kubernetes the dns name "redis" needs to resolve to the production deployment redis instance.**   