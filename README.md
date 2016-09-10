# List processes running in a /proc pseudo file system

This is a small project that lists all running processes in a `/proc` pseudo file system (linux)

## Build and run test env
- docker build -t go .
- docker run -it -v /Users/kimpettersen/projects/gocode/src/github.com/kimpettersen/procs/src:/go/src  go bash
- sh test_processes.sh &
