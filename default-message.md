# docker-on-scheduler

## Use

This container must mount config.yml.

## Sample

### docker

```
$ docker run -d -v ./your_config.yml:/docker-on-scheduler/config.yml iwanomoto/docker-on-scheduler
```

### docker-compose

```
version: "3"
services:
  docker-on-scheduler:
    image: iwanomoto/docker-on-scheduler
    volumes:
      - ./your_config.yml:/docker-on-scheduler/config.yml
```
