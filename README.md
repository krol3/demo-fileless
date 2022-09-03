# Fileless demo

## Demo using Tracee and Postee

Receive a notification if you find suspicious behaviour in your cluster, at the runtime stage.

![gif-demo](./img/tracee-postee-slack.gif)

Install Tracee and Postee in Kubernetes [here](https://aquasecurity.github.io/tracee/v0.8.0/installing/kubernetes/)

---

## Image Scanning

```
trivy image --severity HIGH,CRITICAL --security-checks vuln,secret,config krol/demo-memrun
```

## Linux Events

```
strace -c ls
```

```
strace -c docker run hello-world
```

## Run the container

### Using docker

```
docker run --name demo01 krol/demo-memrun
```
Calling the fileless program

```
docker exec -t demo01 /memrun nginx /bin/date
```

```
docker run -it --rm krol/demo-memrun /memrun nginx /bin/date
```
### Using Kubernetes

```
kubectl run nginx-fileless --image=krol/demo-memrun 

```

Calling the fileless program `kubectl exec -ti nginx-fileless -- /memrun nginx /bin/date`
```
kubectl exec -ti nginx-fileless -- /memrun nginx /bin/date
Sat Sep  3 16:15:26 UTC 2022
```

## Runtime Security
### [Tracee](https://github.com/aquasecurity/tracee)
```
docker run \
   --name tracee --rm -it \
   --pid=host --cgroupns=host --privileged \
   -v /etc/os-release:/etc/os-release-host:ro \
   -e LIBBPFGO_OSRELEASE_FILE=/etc/os-release-host \
   aquasec/tracee:0.8.0


```

[![Tracee Demo Video](./img/fileless-tracee-final.gif)](https://github.com/aquasecurity/tracee)

## More ELFs

```
curl -o /tmp/elf-fileless https://raw.githubusercontent.com/DenizBasgoren/elf32-hello-world/master/a.out && ./memrun nginx /tmp/elf-fileless
```
