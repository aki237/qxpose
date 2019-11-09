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