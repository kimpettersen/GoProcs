# List and kill processes in a linux env

This is a small project that lists all running processes in a `/proc` pseudo file system (linux) and to easily kill it

## Build and run test env
- docker build -t go .
- docker run -it -v /Users/kimpettersen/projects/gocode/src/github.com/kimpettersen/procs/src:/go/src  go bash

# Usage

```
root@283bc8f632a1:/go/src# ./procs
(1) bash  - 1
(2) sh test_process.sh  - 793
(3) sleep 10  - 794
(4) ./procs  - 795

Enter number of process to kill
2
Killed sh test_process.sh
[1]+  Killed                  sh test_process.sh
root@283bc8f632a1:/go/src#
```
