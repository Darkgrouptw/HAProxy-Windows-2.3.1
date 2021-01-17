# ReadMe
## Introduction
I use HAProxy to do load balancer, and use go server to test it.

### Start HAProxy
```Command
cd HAProxy2.3.1
haproxy.exe -f confg.conf -d
```

### Start Golang Server
```Command
cd Test
go run Main_10000.go    # Start server on port 10000
go run Main_10001.go    # Start server on port 10001
```