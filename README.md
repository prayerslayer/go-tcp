# tcp

It's called `tcp` because I originally started with the TCP syscalls, but then I built it to talk a little HTTP too. It can only answer with `"PONG"` and immediately close the connection though (essentially doing HTTP/1.0). 

## Run

~~~ bash
export TCP_DIR=$GOPATH/src/github.com/prayerslayer/tcp
git clone https://github.com/prayerslayer/tcp.git $TCP_DIR
cd $TCP_DIR
go run main.go
# then elsewhere 'curl localhost:3000'
~~~

~~~
curl -vvv localhost:3000
* Rebuilt URL to: localhost:3000/
*   Trying ::1...
* connect to ::1 port 3000 failed: Connection refused
*   Trying 127.0.0.1...
* Connected to localhost (127.0.0.1) port 3000 (#0)
> GET / HTTP/1.1
> Host: localhost:3000
> User-Agent: curl/7.43.0
> Accept: */*
> 
< HTTP/1.1 200 OK
< Connection: close
< 
* Closing connection 0
"PONG"%    
~~~

## Benchmark

Just for the fun of it.

~~~
> wrk -c 2 -t 2 -d 10s http://127.0.0.1:3000
Running 10s test @ http://127.0.0.1:3000
  2 threads and 2 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   374.11us    1.96ms  43.32ms   98.59%
    Req/Sec   720.05    703.22     2.77k    78.48%
  6340 requests in 10.10s, 272.42KB read
  Socket errors: connect 0, read 6349, write 1, timeout 0
Requests/sec:    627.53
Transfer/sec:     26.96KB
~~~