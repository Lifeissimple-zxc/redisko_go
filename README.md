# Redisko (GO)

This repo contains code i wrote following a redis tutorial for GO to get familiar with the tool.

## Setup

Pull the latest redis image from the docker registry.
```bash
docker pull redis 
```

Run the container & map our machine's port to the one used by the container.
```bash
docker run --name redis-tutorial -p 6379:6379 -d redis
```

`-p` maps ports.<br>
`1` means detached mode.<br>
`redis` is the container name we pulled.

Make sure the new container is running with `docker ps`.