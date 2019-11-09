# qxpose

WIP : Expose local tunnels to www

### Architechture

```
                      _________________
 Behind NAT          |     Exposed     |
                     |                 |
Local <--- QUIC ---> |     Tunnel      | <--- TLS ---> Browser/cURL/openssl
                     |                 |
                     |_________________|
```

### Why QUIC?

Head of line blocking in TCP is a pain and ability to initiate streams from server
without the need of any signalling or control streams is a nice to have to achieve low
latency proxying.

### Try it

Go build first.

*Server*

```
# run qxpose as server mode with the configured sub domain as poniesareaweso.me
# and the idle time out for QUIC sessions as an hour (default is 1/2 hour)
qxpose server --domain poniesareaweso.me -i 3600
```

```
# run qxpose as client mode with the following options
#  1. Tunnel server: to the locally running one
#  2. Local: Which local server/TCP address to proxy to public.
#  3. Idle Timeout: idle time out for QUIC sessions as an hour (default is 1/2 hour)
qxpose client --tunnel "localhost:2723" --local "localhost:8100" -i 3600
```

The client spits out a new hostname for the tunnel. (something like fb6b5b1749f59e70.poniesareaweso.me)
For locally testing, edit the /etc/hosts to point
the host to 127.0.0.1. 

Something like this.
```
127.0.0.1   fb6b5b1749f59e70.poniesareaweso.me
```

Now try the address in the browser (insecure) or cURL (with `-k` flag).