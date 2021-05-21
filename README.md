# fastjson-serialization-bench

This repo was created in response to the following blog post:  
https://talawah.io/blog/extreme-http-performance-tuning-one-point-two-million/

It demonstrates a 245k req/s baseline using fastjson and fasthttp with a c5n.4xlarge client and c5n.xlarge server in a placement group on an unmodified AWS linux 2 AMI.

## Ground Zero

```
> $ ./twrk -t 16 -c 256 -D 2 -d 10 --latency --pin-cpus "http://172.31.35.0:8080/json" -H 'Host: server.tfb' -H 'Accept: application/json,text/html;q=0.9,application/xhtml+xml;q=0.9,application/xml;q=0.8,*/*;q=0.7' -H 'Connection: keep-alive'
Running 10s test @ http://172.31.35.0:8080/json
  16 threads and 256 connections
  Thread Stats   Avg     Stdev       Max       Min   +/- Stdev
    Latency     1.13ms    0.87ms   13.59ms   64.00us   77.10%
    Req/Sec    15.40k   343.76     16.42k    13.43k    71.65%
  Latency Distribution
  50.00%    0.97ms
  90.00%    2.18ms
  99.00%    4.12ms
  99.99%   11.04ms
  2452813 requests in 10.00s, 392.98MB read
Requests/sec: 245279.56
Transfer/sec:     39.30MB
```

Server setup

```
wget https://storage.googleapis.com/golang/go1.16.4.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.16.4.linux-amd64.tar.gz
sudo ln -s /usr/local/go/bin/go /usr/bin/go
sudo mkdir /usr/local/share/go
sudo mkdir /usr/local/share/go/bin
sudo chmod 777 /usr/local/share/go

git clone https://github.com/kevburnsjr/fastjson-serialization-bench.git
cd fastjson-serialization-bench
go build ldflags="-s -w" main.go
```

Running the server:

```
./main
```

Make sure you have 8080 open in your security group.
