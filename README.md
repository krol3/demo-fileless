# Fileless demo

## Image Scanning

```
trivy image --severity HIGH,CRITICAL --security-checks vuln,secret,config krol/demo-memrun
```

## Linux Events

```
strace -c ls

strace -c docker run hello-world
```

## Run the container

```
docker run --name demo01 krol/demo-memrun

docker exec -t demo01 /memrun nginx /bin/date
```

## Runtime Security Tool
### [Tracee](https://github.com/aquasecurity/tracee)
```
docker run \
  --name tracee --rm -it \
  --pid=host --cgroupns=host --privileged \
  -v /etc/os-release:/etc/os-release-host:ro \
  -e LIBBPFGO_OSRELEASE_FILE=/etc/os-release-host \
  aquasec/tracee:0.7.0

```
