# fastjson-serialization-bench

This repo does 230k+ req/s out of the box with a c5n.xlarge client and c5n.xlarge server on an unmodified AWS linux 2 AMI.

## Ground Zero

```
> ./twrk -t 4 -c 256 -D 2 -d 10 --latency --pin-cpus "http://54.213.147.90:8080/json" -H 'Host: server.tfb' -H 'Accept: application/json,text/html;q=0.9,application/xhtml+xml;q=0.9,application/xml;q=0.8,*/*;q=0.7' -H 'Connection: keep-alive'
Running 10s test @ http://54.213.147.90:8080/json
  4 threads and 256 connections
  Thread Stats   Avg     Stdev       Max       Min   +/- Stdev
    Latency     1.11ms    1.94ms  209.41ms  165.00us   97.99%
    Req/Sec    59.59k    12.28k    72.07k    37.29k    74.75%
  Latency Distribution
  50.00%    0.94ms
  90.00%    2.01ms
  99.00%    3.58ms
  99.99%  116.62ms
  2371816 requests in 10.00s, 380.01MB read
Requests/sec: 237179.84
Transfer/sec:     38.00MB
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
