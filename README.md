# docker-on-scheduler

Docker on Scheduler tool.
Periodically execute commands within Docker.

## Installing

```
$ docker pull iwanomoto/docker-on-scheduler
```

## Use

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
