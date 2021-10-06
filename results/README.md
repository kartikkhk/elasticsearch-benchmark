## Local Loadtest Results

- Type: READ
- VUS: 10
- DURATION: 10m
```
checks.........................: 0.00%  ✓ 0          ✗ 269697
data_received..................: 1.2 GB 1.9 MB/s
data_sent......................: 27 MB  46 kB/s
http_req_blocked...............: avg=4.58µs  min=1µs    med=4µs     max=3.03ms  p(90)=6µs     p(95)=7µs
http_req_connecting............: avg=20ns    min=0s     med=0s      max=1.28ms  p(90)=0s      p(95)=0s
http_req_duration..............: avg=21.43ms min=4.5ms  med=19.32ms max=1.42s   p(90)=33.62ms p(95)=39.52ms
  { expected_response:true }...: avg=21.43ms min=4.5ms  med=19.32ms max=1.42s   p(90)=33.62ms p(95)=39.52ms
http_req_failed................: 0.00%  ✓ 0          ✗ 269697
http_req_receiving.............: avg=88.96µs min=16µs   med=81µs    max=24.98ms p(90)=133µs   p(95)=155µs
http_req_sending...............: avg=20.45µs min=5µs    med=18µs    max=10.17ms p(90)=30µs    p(95)=39µs
http_req_tls_handshaking.......: avg=0s      min=0s     med=0s      max=0s      p(90)=0s      p(95)=0s
http_req_waiting...............: avg=21.32ms min=4.41ms med=19.21ms max=1.42s   p(90)=33.5ms  p(95)=39.4ms
http_reqs......................: 269697 449.476839/s
iteration_duration.............: avg=22.23ms min=4.74ms med=19.98ms max=1.44s   p(90)=34.63ms p(95)=40.73ms
iterations.....................: 269697 449.476839/s
vus............................: 10     min=10       max=10
vus_max........................: 10     min=10       max=10

```
