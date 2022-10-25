
# Demo Falco fileless - memfd_create

## Demo

kubectl run demo-memfd --image=krol/demo-memfd:v3

To generate a fileless call.

k exec -ti demo-memfd -- /demo-memfd nginx /bin/date

![](https://i.imgur.com/sDy8Tjy.png)
 
## Rule

`k edit cm falco`

Falco default rules [here](https://github.com/falcosecurity/falco/blob/master/rules/falco_rules.yaml): 

### Custom Rule in Falco

```
  falco_rules.local.yaml: |
    - rule: Fileless detection
      desc: This rule detect memfd_create malicious activity
      condition: evt.type in (execve, execveat) and evt.dir=> and (evt.arg.path startswith /proc/self/fd/ or evt.arg.name startswith /proc/self/fd/ or fd.name startswith /proc/self/fd or evt.arg.filename startswith /proc/self/fd/)
      output: "---> Fileless detection (image=%container.image.repository evt.type=%evt.type evt.arg.path=%evt.arg.path evt.arg.name=%evt.arg.name evt.arg.filename=%evt.arg.filename fd.name=%fd.name proc.name=%proc.name proc.cmdline=%proc.cmdline)"
      priority: CRITICAL
      tags: [mitre_execution]
  
```

## Output


````
k logs -f falco-fhl6n
Defaulted container "falco" out of: falco, falco-driver-loader (init)
Tue Oct 18 02:07:57 2022: Falco version 0.32.2
Tue Oct 18 02:07:57 2022: Falco initialized with configuration file /etc/falco/falco.yaml
Tue Oct 18 02:07:57 2022: Loading rules from file /etc/falco/falco_rules.yaml:
Tue Oct 18 02:07:57 2022: Loading rules from file /etc/falco/falco_rules.local.yaml:
Tue Oct 18 02:07:58 2022: Starting internal webserver, listening on port 8765

02:08:00.509507302: Critical ---> Fileless detection (image=<NA> evt.type=execve evt.arg.path=<NA> evt.arg.name=<NA> evt.arg.filename=/proc/self/fd/5 fd.name=<NA> proc.name=exe proc.cmdline=exe init) k8s.ns=<NA> k8s.pod=<NA> container=host
02:08:01.611672787: Critical ---> Fileless detection (image=<NA> evt.type=execve evt.arg.path=<NA> evt.arg.name=<NA> evt.arg.filename=/proc/self/fd/5 fd.name=<NA> proc.name=exe proc.cmdline=exe init) k8s.ns=<NA> k8s.pod=<NA> container=host
02:08:10.513496600: Critical ---> Fileless detection (image=<NA> evt.type=execve evt.arg.path=<NA> evt.arg.name=<NA> evt.arg.filename=/proc/self/fd/5 fd.name=<NA> proc.name=exe proc.cmdline=exe init) k8s.ns=<NA> k8s.pod=<NA> container=host
02:08:11.635821457: Critical ---> Fileless detection (image=<NA> evt.type=execve evt.arg.path=<NA> evt.arg.name=<NA> evt.arg.filename=/proc/self/fd/5 fd.name=<NA> proc.name=exe proc.cmdline=exe init) k8s.ns=<NA> k8s.pod=<NA> container=host
02:08:15.716874207: Critical ---> Fileless detection (image=<NA> evt.type=execve evt.arg.path=<NA> evt.arg.name=<NA> evt.arg.filename=/proc/self/fd/6 fd.name=<NA> proc.name=exe proc.cmdline=exe init) k8s.ns=<NA> k8s.pod=<NA> container=host
02:08:15.751225857: Critical ---> Fileless detection (image=krol/demo-memfd evt.type=execve evt.arg.path=<NA> evt.arg.name=<NA> evt.arg.filename=/proc/self/fd/3 fd.name=<NA> proc.name=demo-memfd proc.cmdline=demo-memfd nginx /bin/date) k8s.ns=default k8s.pod=demo-memfd container=c56befd1fe4d
02:08:20.523472031: Critical ---> Fileless detection (image=<NA> evt.type=execve evt.arg.path=<NA> evt.arg.name=<NA> evt.arg.filename=/proc/self/fd/5 fd.name=<NA> proc.name=exe proc.cmdline=exe init) k8s.ns=<NA> k8s.pod=<NA> container=host
02:08:21.634042785: Critical ---> Fileless detection (image=<NA> evt.type=execve evt.arg.path=<NA> evt.arg.name=<NA> evt.arg.filename=/proc/self/fd/5 fd.name=<NA> proc.name=exe proc.cmdline=exe init) k8s.ns=<NA> k8s.pod=<NA> container=host
02:08:30.504928868: Critical ---> Fileless detection (image=<NA> evt.type=execve evt.arg.path=<NA> evt.arg.name=<NA> evt.arg.filename=/proc/self/fd/5 fd.name=<NA> proc.name=exe proc.cmdline=exe init) k8s.ns=<NA> k8s.pod=<NA> container=host
````
## Falco notes about memfd

Falco detect memfd + exec: https://hackmd.io/@leogr/SJKUMEbWo#loresuso-proposal

syscall memfd_create is not supported by Falco: 

- https://github.com/falcosecurity/libs/pull/595#issuecomment-1247743183
- https://github.com/falcosecurity/falco/issues/1998
