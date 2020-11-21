# File System Experiment

1. `go run main.go`
1. `px aux | grep main.go`
1. Grab pid from ^
1. `pgrep -P [pid]`
1. Grab child pid from ^
1. `lsof -p [pid]`
1. Sort it: `lsof -p [pid] | sort -n -r -k7`
1. `rm output.txt`
1. Watch the file count continue to go up from previous step because the file descriptor is still open.

Sort all file handlers by size:
```
lsof | sort -n -r -k7 | head -10
```
