# Local Loadtest Results

### Golang Results

- **Type: READ**
- **VUS: 10**
- **DURATION: 10m**

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

- **Type: WRITE**
- **VUS: 10**
- **DURATION: 10m**

```
checks.........................: 100.00% ✓ 219261     ✗ 0
data_received..................: 75 MB   126 kB/s
data_sent......................: 91 MB   151 kB/s
http_req_blocked...............: avg=4.55µs  min=1µs    med=4µs     max=1.79ms  p(90)=6µs     p(95)=7µs
http_req_connecting............: avg=15ns    min=0s     med=0s      max=447µs   p(90)=0s      p(95)=0s
http_req_duration..............: avg=20.36ms min=5.63ms med=17.34ms max=1.03s   p(90)=30.05ms p(95)=36.94ms
 { expected_response:true }...: avg=20.36ms min=5.63ms med=17.34ms max=1.03s   p(90)=30.05ms p(95)=36.94ms
http_req_failed................: 0.00%   ✓ 0          ✗ 219261
http_req_receiving.............: avg=66.89µs min=16µs   med=62µs    max=14.17ms p(90)=99µs    p(95)=117µs
http_req_sending...............: avg=23.99µs min=7µs    med=22µs    max=6.15ms  p(90)=34µs    p(95)=43µs
http_req_tls_handshaking.......: avg=0s      min=0s     med=0s      max=0s      p(90)=0s      p(95)=0s
http_req_waiting...............: avg=20.27ms min=5.56ms med=17.25ms max=1.03s   p(90)=29.96ms p(95)=36.86ms
http_reqs......................: 219261  365.420649/s
iteration_duration.............: avg=27.34ms min=5.81ms med=19.61ms max=1.03s   p(90)=49.01ms p(95)=50.69ms
iterations.....................: 219261  365.420649/s
vus............................: 10      min=10       max=10
vus_max........................: 10      min=10       max=10
```

---

### NodeJS Results

- **Type: READ**
- **VUS: 10**
- **DURATION: 10m**

```
checks.........................: 100.00% ✓ 259918     ✗ 0
data_received..................: 1.1 GB  1.8 MB/s
data_sent......................: 26 MB   44 kB/s
http_req_blocked...............: avg=4.87µs  min=1µs    med=4µs     max=1.22ms  p(90)=6µs     p(95)=8µs
http_req_connecting............: avg=13ns    min=0s     med=0s      max=447µs   p(90)=0s      p(95)=0s
http_req_duration..............: avg=19.17ms min=4.59ms med=17.38ms max=1.05s   p(90)=28.48ms p(95)=33.55ms
 { expected_response:true }...: avg=19.17ms min=4.59ms med=17.38ms max=1.05s   p(90)=28.48ms p(95)=33.55ms
http_req_failed................: 0.00%   ✓ 0          ✗ 259918
http_req_receiving.............: avg=81.28µs min=15µs   med=73µs    max=6.66ms  p(90)=118µs   p(95)=139µs
http_req_sending...............: avg=20.38µs min=5µs    med=18µs    max=12.08ms p(90)=28µs    p(95)=36µs
http_req_tls_handshaking.......: avg=0s      min=0s     med=0s      max=0s      p(90)=0s      p(95)=0s
http_req_waiting...............: avg=19.07ms min=4.51ms med=17.28ms max=1.05s   p(90)=28.37ms p(95)=33.45ms
http_reqs......................: 259918  433.169926/s
iteration_duration.............: avg=23.06ms min=4.76ms med=19.27ms max=1.05s   p(90)=36.45ms p(95)=48.72ms
iterations.....................: 259918  433.169926/s
vus............................: 10      min=10       max=10
vus_max........................: 10      min=10       max=10
```
