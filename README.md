# Helmsman - The Chart Repository

At a high level, a chart repository is a location where packaged charts can be
stored and shared.

You will usually need to create a private repository for your charts. Helmsman provides a simple way to accomplish this task.

## Prerequisites

There are no prerequisites.

**Note:** It is not necessary to have the helm command

## Docker

```console
docker run -p 8000:8000 -e HELMSMAN_REPO_URL=127.0.0.1:8000 -v /home/ubuntu/charts:/app/charts andrepinto/helmsman
```

## Run

```console
$  ./helmsman -h

NAME:
   helmsman - A new cli application

USAGE:
   helmsman [global options] command [command options] [arguments...]

VERSION:
   1.1.0

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --debug           Debug mode default to false
   --port value      server port (default: 8000)
   --env value       (default: "stable")
   --repo.dir value  helm repo dir (default: "./charts")
   --repo.url value  helm repo url (default: "localhost:8000/charts/")
   --help, -h        show help
   --version, -v     print the version

```

## Api

Get Charts


```console
$ curl http://localhost:8000/envs/stable/charts/index.yaml

apiVersion: v1
entries:
  demo:
  - apiVersion: v1
    created: 2017-09-10T15:51:42.251080332+01:00
    description: A Helm chart for Kubernetes
    digest: ee8d3e7ee1fc682461cf1f10715504e6f6dabf5f8c023cfbcaecfbcc0a02753c
    name: demo
    urls:
    - localhost:8000/charts/demo-v1.0.0.tgz
    version: v1.0.0
generated: 2017-09-10T15:51:42.238429632+01:00

```


Upload Chart

```console
$ curl -v -T repo.tar.gz -X PUT http://localhost:8000/envs/stable/charts/upload/
```

## Build

```console
$  glide install
```

```console
$  go build -o helmsman cmd/server/main.go
```