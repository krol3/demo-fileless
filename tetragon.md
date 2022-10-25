# Tetragon Demo Fileless

Examples about Tetragon rules [here](https://github.com/cilium/tetragon/tree/main/crds/examples)

`kubectl run demo-memfd --image=krol/demo-memfd:v1.1.2 `

`kubectl logs -n kube-system ds/tetragon -c export-stdout -f | tetragon observe`

![](https://i.imgur.com/7fBpeAS.png)


## Tetragon policy

PR: https://github.com/cilium/tetragon/pull/484

```
apiVersion: cilium.io/v1alpha1
kind: TracingPolicy
metadata:
  name: "demo-memfd"
spec:
  kprobes:
# int close(int fd);
  - call: "__x64_sys_close"
    syscall: true
    args:
    - index: 0
      type: "int"
    selectors:
    - matchPIDs:
      - operator: NotIn
        followForks: true
        isNamespacePID: true
        values:
        - 0
        - 1
      matchActions:
      - action: UnfollowFD
        argFd: 0
        argName: 0
  # int memfd_create(const char *name, unsigned int flags);
  - call: "__x64_sys_memfd_create"
    syscall: true
    args:
    - index: 0
      type: "string"
    - index: 1
      type: "int"
    selectors:
    - matchPIDs:
      - operator: NotIn
        followForks: true
        isNamespacePID: true
        values:
        - 0
        - 1
# int execve(const char *pathname, char *const argv[],char *const envp[]);
  - call: "__x64_sys_execve"
    syscall: true
    args:
    - index: 0
      type: "string"
    selectors:
    - matchPIDs:
      - operator: NotIn
        followForks: false
        isNamespacePID: true
        values:
        - 0
        - 1
      matchArgs:
      - index: 0
        operator: "Prefix"
        values:
        - "/proc/self/fd/"
      matchActions:
      - action: Sigkill
```


## With kill enabled

````
🚀 process default/demo-memfd /demo-memfd nginx /bin/date
⁉️ syscall default/demo-memfd /demo-memfd __x64_sys_execve
💥 exit    default/demo-memfd /demo-memfd nginx /bin/date SIGKILL
````
Automatic detection by Tetragon: https://github.com/cilium/tetragon/pull/499
