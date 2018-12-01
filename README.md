# go-udp-service-discovery
UDP based service discovery written in Go. Painfully simple

# Start the server

Run `go-udp-service-discovery -h` to see command line options

```
go-udp-service-discovery -port 1234                                                                                                               ✘ 1
Starting udp server on port 1234
```

Stdout will report the following for connections. For example, if you ran `echo test | ncat -u localhost 1234`, the server would spit this out into stdout:

```
Received  test
  from  [::1]:51402

```