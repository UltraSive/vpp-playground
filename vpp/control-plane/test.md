root@prox:~# tcpdump
tcpdump: verbose output suppressed, use -v[v]... for full protocol decode
listening on enp1s0f0, link-type EN10MB (Ethernet), snapshot length 262144 bytes
23:30:37.728901 IP prox.anchored.host.ssh > syn-071-010-196-052.res.spectrum.com.51692: Flags [P.], seq 2595267422:2595267618, ack 3036684211, win 249, length 196
23:30:37.746234 IP 185.42.12.58.28244 > cloudpanel.cwgdata.center.ms-wbt-server: Flags [P.], seq 2261210840:2261210906, ack 4282930257, win 1021, length 66
23:30:37.746255 IP 185.42.12.58.28244 > cloudpanel.cwgdata.center.ms-wbt-server: Flags [P.], seq 66:104, ack 1, win 1021, length 38
23:30:37.746306 IP 185.42.12.58.28244 > cloudpanel.cwgdata.center.ms-wbt-server: Flags [R.], seq 104, ack 1, win 0, length 0
23:30:37.746363 IP cloudpanel.cwgdata.center.ms-wbt-server > 185.42.12.58.28244: Flags [.], ack 104, win 62622, length 0
23:30:37.757402 ARP, Request who-has 103.195.101.201 tell 103.195.101.122, length 46
23:30:37.758079 IP 185.42.12.58.1817 > cloudpanel.cwgdata.center.ms-wbt-server: Flags [SEW], seq 2368029308, win 64240, options [mss 1460,nop,wscale 8,nop,nop,sackOK], length 0
23:30:37.758209 IP cloudpanel.cwgdata.center.ms-wbt-server > 185.42.12.58.1817: Flags [S.E], seq 1834185827, ack 2368029309, win 64000, options [mss 1460,nop,wscale 0,nop,nop,sackOK], length 0
23:30:37.773175 IP 162.159.133.234.https > cloudpanel.cwgdata.center.55264: Flags [P.], seq 1444901594:1444901648, ack 4137875152, win 9, length 54
23:30:37.777987 IP syn-071-010-196-052.res.spectrum.com.51692 > prox.anchored.host.ssh: Flags [.], ack 196, win 1023, length 0
23:30:37.784848 IP cloudpanel.cwgdata.center.55264 > 162.159.133.234.https: Flags [.], ack 54, win 1026, length 0
23:30:37.803611 IP syn-071-010-196-052.res.spectrum.com.53628 > 103.195.102.87.ssh: Flags [.], ack 4259016568, win 1023, length 0
23:30:37.812523 IP serverarmour.com.27593 > 103.195.102.91.28016: Flags [.], ack 2573204667, win 1921, options [nop,nop,TS val 1319202979 ecr 2005403699], length 0
23:30:37.813193 IP 162.159.134.234.https > cloudpanel.cwgdata.center.56549: Flags [P.], seq 1034061395:1034061438, ack 17106822, win 9, length 43
23:30:37.813595 IP 162.159.130.234.https > cloudpanel.cwgdata.center.57085: Flags [P.], seq 3636082948:3636082991, ack 3988253106, win 9, length 43
23:30:37.816659 IP prox.anchored.host.60844 > dns.google.domain: 49523+ PTR? 52.196.10.71.in-addr.arpa. (43)
23:30:37.817447 ARP, Request who-has 104.238.205.229 tell 104.238.205.1, length 46
23:30:37.817485 ARP, Request who-has 104.238.205.228 tell 104.238.205.1, length 46
23:30:37.817496 ARP, Request who-has 45.126.211.91 tell 45.126.211.1, length 46
23:30:37.837948 IP cloudpanel.cwgdata.center.57085 > 162.159.130.234.https: Flags [.], ack 43, win 1023, length 0
23:30:37.837955 IP cloudpanel.cwgdata.center.56549 > 162.159.134.234.https: Flags [.], ack 43, win 1023, length 0
23:30:37.850428 ARP, Request who-has 169.254.169.254 tell 45.126.208.218.elitehost.com.br, length 46
23:30:37.869913 IP dns.google.domain > prox.anchored.host.60844: 49523 1/0/0 PTR syn-071-010-196-052.res.spectrum.com. (93)
23:30:37.870123 IP prox.anchored.host.42707 > dns.google.domain: 42088+ PTR? 101.110.93.172.in-addr.arpa. (45)
23:30:37.870127 IP prox.anchored.host.ssh > syn-071-010-196-052.res.spectrum.com.51692: Flags [P.], seq 196:408, ack 1, win 249, length 212
23:30:37.872401 ARP, Request who-has 191.101.251.75 tell 172.93.110.96, length 46
23:30:37.882479 IP dns.google.domain > prox.anchored.host.42707: 42088 1/0/0 PTR cloudpanel.cwgdata.center. (84)
23:30:37.882545 IP prox.anchored.host.53435 > dns.google.domain: 59579+ PTR? 58.12.42.185.in-addr.arpa. (43)
23:30:37.883393 ARP, Request who-has 169.254.151.22 tell 199.127.62.249, length 46
23:30:37.904559 IP 185.42.12.58.28244 > cloudpanel.cwgdata.center.ms-wbt-server: Flags [R], seq 2261210944, win 0, length 0
23:30:37.908562 IP 103.195.102.88.ms-wbt-server > 77.90.185.223.9599: Flags [P.], seq 3683634740:3683634771, ack 3965235103, win 62698, length 31
23:30:37.908885 IP dns.google.domain > prox.anchored.host.53435: 59579 NXDomain 0/1/0 (103)
23:30:37.908993 IP prox.anchored.host.ssh > syn-071-010-196-052.res.spectrum.com.51692: Flags [P.], seq 408:860, ack 1, win 249, length 452
23:30:37.909011 IP prox.anchored.host.ssh > syn-071-010-196-052.res.spectrum.com.51692: Flags [P.], seq 860:1024, ack 1, win 249, length 164
23:30:37.909012 IP prox.anchored.host.58477 > dns.google.domain: 5812+ PTR? 201.101.195.103.in-addr.arpa. (46)
23:30:37.911930 IP 185.42.12.58.1817 > cloudpanel.cwgdata.center.ms-wbt-server: Flags [.], ack 1, win 1026, length 0
23:30:37.911948 IP 185.42.12.58.1817 > cloudpanel.cwgdata.center.ms-wbt-server: Flags [P.], seq 1:46, ack 1, win 1026, length 45
23:30:37.914609 IP cloudpanel.cwgdata.center.ms-wbt-server > 185.42.12.58.1817: Flags [P.], seq 1:20, ack 46, win 63955, length 19
23:30:37.917911 IP serverarmour.com.57045 > 103.195.102.91.28016: Flags [P.], seq 2258532637:2258532704, ack 1951549570, win 7631, options [nop,nop,TS val 1319203085 ecr 2005403599], length 67
23:30:37.941853 IP serverarmour.com.61849 > 103.195.102.91.28016: Flags [P.], seq 4250294474:4250294541, ack 1402459404, win 846, options [nop,nop,TS val 1319203108 ecr 2005402999], length 67
23:30:37.941976 IP 103.195.102.91.28016 > serverarmour.com.61849: Flags [.], ack 67, win 501, options [nop,nop,TS val 2005403946 ecr 1319203108], length 0
23:30:37.958623 IP 103.195.102.91.28016 > serverarmour.com.57045: Flags [.], ack 67, win 501, options [nop,nop,TS val 2005403963 ecr 1319203085], length 0
23:30:37.958738 IP syn-071-010-196-052.res.spectrum.com.51692 > prox.anchored.host.ssh: Flags [.], ack 1024, win 1026, length 0
23:30:37.960155 IP dns.google.domain > prox.anchored.host.58477: 5812 NXDomain 0/1/0 (107)
23:30:37.960284 IP prox.anchored.host.35933 > dns.google.domain: 21334+ PTR? 122.101.195.103.in-addr.arpa. (46)
23:30:37.975807 IP 45.126.210.116.56771 > 239.255.255.250.1900: UDP, length 176
23:30:37.982401 ARP, Request who-has 104.238.205.230 tell 104.238.205.1, length 46
23:30:37.993438 ARP, Request who-has 104.238.205.216 tell 104.238.205.1, length 46
23:30:37.993449 ARP, Request who-has 104.238.204.46 tell 104.238.204.1, length 46
23:30:37.993453 IP6 2605:9880:400:100:0:a6:1240:1 > ff02::1:ff00:1: ICMP6, neighbor solicitation, who has 2605:9880:400::1, length 32
23:30:37.993462 ARP, Request who-has 172.96.143.158 tell 172.96.143.1, length 46
23:30:37.993470 ARP, Request who-has 104.238.205.179 tell 104.238.205.1, length 46
23:30:37.994257 IP cloudpanel.cwgdata.center.ms-wbt-server > 80.64.30.82.30771: Flags [P.], seq 1174347454:1174347491, ack 2986254480, win 62680, length 37
23:30:37.994955 IP 103.195.102.91.28016 > serverarmour.com.57045: Flags [P.], seq 1:87, ack 67, win 501, options [nop,nop,TS val 2005403999 ecr 1319203085], length 8623:30:37.995183 IP 103.195.102.91.28016 > serverarmour.com.61849: Flags [P.], seq 1:646, ack 67, win 501, options [nop,nop,TS val 2005403999 ecr 1319203108], length 645
23:30:38.009013 IP dns.google.domain > prox.anchored.host.35933: 21334 NXDomain 0/1/0 (107)
23:30:38.009166 IP prox.anchored.host.ssh > syn-071-010-196-052.res.spectrum.com.51692: Flags [P.], seq 1024:1156, ack 1, win 249, length 132
23:30:38.009186 IP prox.anchored.host.ssh > syn-071-010-196-052.res.spectrum.com.51692: Flags [P.], seq 1156:1576, ack 1, win 249, length 420
23:30:38.009219 IP prox.anchored.host.51598 > dns.google.domain: 46558+ PTR? 234.133.159.162.in-addr.arpa. (46)
23:30:38.011240 IP serverarmour.com.52832 > 103.195.102.91.28016: Flags [P.], seq 1366562501:1366562568, ack 3796494801, win 737, options [nop,nop,TS val 1319203177 ecr 2005403099], length 67
23:30:38.011333 IP 103.195.102.91.28016 > serverarmour.com.52832: Flags [.], ack 67, win 501, options [nop,nop,TS val 2005404015 ecr 1319203177], length 0
23:30:38.013150 IP serverarmour.com.6904 > 103.195.102.91.28016: Flags [P.], seq 372612658:372612725, ack 681329199, win 825, options [nop,nop,TS val 1319203179 ecr 2005403299], length 67
23:30:38.013224 IP 103.195.102.91.28016 > serverarmour.com.6904: Flags [.], ack 67, win 501, options [nop,nop,TS val 2005404017 ecr 1319203179], length 0
23:30:38.015408 IP6 :: > ff02::1:ff23:be01: ICMP6, neighbor solicitation, who has fe80::215:5dff:fe23:be01, length 32
23:30:38.015409 ARP, Request who-has 103.195.101.125 tell 103.195.101.122, length 46
23:30:38.022242 IP serverarmour.com.34204 > 103.195.102.91.28016: Flags [P.], seq 4049695733:4049695800, ack 3945535642, win 10087, options [nop,nop,TS val 1319203189 ecr 2005403099], length 67
23:30:38.022321 IP 103.195.102.91.28016 > serverarmour.com.34204: Flags [.], ack 67, win 501, options [nop,nop,TS val 2005404026 ecr 1319203189], length 0
23:30:38.026392 ARP, Request who-has 104.238.205.171 tell 104.238.205.1, length 46
23:30:38.036122 IP dns.google.domain > prox.anchored.host.51598: 46558 NXDomain 0/1/0 (108)
23:30:38.036265 IP prox.anchored.host.ssh > syn-071-010-196-052.res.spectrum.com.51692: Flags [P.], seq 1576:1772, ack 1, win 249, length 196
23:30:38.036286 IP prox.anchored.host.ssh > syn-071-010-196-052.res.spectrum.com.51692: Flags [P.], seq 1772:2064, ack 1, win 249, length 292
23:30:38.036291 IP prox.anchored.host.39203 > dns.google.domain: 46340+ PTR? 87.102.195.103.in-addr.arpa. (45)
23:30:38.041911 IP serverarmour.com.27947 > 103.195.102.91.28016: Flags [P.], seq 1653304046:1653304113, ack 126571842, win 3247, options [nop,nop,TS val 1319203209 ecr 2005403099], length 67
23:30:38.041991 IP 103.195.102.91.28016 > serverarmour.com.27947: Flags [.], ack 67, win 501, options [nop,nop,TS val 2005404046 ecr 1319203209], length 0
23:30:38.042225 IP 77.90.185.223.9599 > 103.195.102.88.ms-wbt-server: Flags [P.], seq 1:67, ack 31, win 251, length 66
23:30:38.051416 IP 77.90.185.223.9599 > 103.195.102.88.ms-wbt-server: Flags [P.], seq 67:105, ack 31, win 251, length 38
23:30:38.051489 IP 103.195.102.88.ms-wbt-server > 77.90.185.223.9599: Flags [.], ack 105, win 62594, length 0
23:30:38.059142 IP 77.90.185.223.9599 > 103.195.102.88.ms-wbt-server: Flags [R.], seq 105, ack 31, win 0, length 0
23:30:38.059655 IP syn-071-010-196-052.res.spectrum.com.51692 > prox.anchored.host.ssh: Flags [.], ack 1576, win 1024, length 0
23:30:38.061693 IP 77.90.185.223.36728 > 103.195.102.88.ms-wbt-server: Flags [SEW], seq 4107341495, win 8192, options [mss 1460,nop,wscale 8,nop,nop,sackOK], length 023:30:38.061828 IP 103.195.102.88.ms-wbt-server > 77.90.185.223.36728: Flags [S.E], seq 300349661, ack 4107341496, win 64000, options [mss 1460,nop,wscale 0,nop,nop,sackOK], length 0
23:30:38.068480 IP 185.42.12.58.1817 > cloudpanel.cwgdata.center.ms-wbt-server: Flags [P.], seq 46:193, ack 20, win 1026, length 147
23:30:38.068801 IP cloudpanel.cwgdata.center.ms-wbt-server > 185.42.12.58.1817: Flags [P.], seq 20:866, ack 193, win 63808, length 846
23:30:38.070394 ARP, Request who-has 104.238.205.243 tell 104.238.205.1, length 46
23:30:38.085110 IP syn-071-010-196-052.res.spectrum.com.51692 > prox.anchored.host.ssh: Flags [.], ack 2064, win 1022, length 0
23:30:38.092439 ARP, Request who-has 104.238.205.240 tell 104.238.205.1, length 46
23:30:38.092449 ARP, Request who-has 104.238.204.7 tell 104.238.204.1, length 46
23:30:38.095133 IP 103.195.102.91.28016 > serverarmour.com.52832: Flags [P.], seq 1:646, ack 67, win 501, options [nop,nop,TS val 2005404099 ecr 1319203177], length 645
23:30:38.095205 IP 103.195.102.91.28016 > serverarmour.com.6904: Flags [P.], seq 1:87, ack 67, win 501, options [nop,nop,TS val 2005404099 ecr 1319203179], length 86
23:30:38.095437 IP 103.195.102.91.28016 > serverarmour.com.34204: Flags [P.], seq 1:646, ack 67, win 501, options [nop,nop,TS val 2005404099 ecr 1319203189], length 645
23:30:38.095653 IP 103.195.102.91.28016 > serverarmour.com.27947: Flags [P.], seq 1:646, ack 67, win 501, options [nop,nop,TS val 2005404100 ecr 1319203209], length 645
23:30:38.095840 IP dns.google.domain > prox.anchored.host.39203: 46340 NXDomain 0/1/0 (106)
23:30:38.095999 IP prox.anchored.host.ssh > syn-071-010-196-052.res.spectrum.com.51692: Flags [P.], seq 2064:2244, ack 1, win 249, length 180
23:30:38.096016 IP prox.anchored.host.44724 > dns.google.domain: 30824+ PTR? 91.102.195.103.in-addr.arpa. (45)
23:30:38.099493 IP serverarmour.com.4244 > 103.195.102.91.28016: Flags [P.], seq 2407960977:2407961044, ack 2101578721, win 10106, options [nop,nop,TS val 1319203266 ecr 2005403399], length 67
23:30:38.099564 IP 103.195.102.91.28016 > serverarmour.com.4244: Flags [.], ack 67, win 501, options [nop,nop,TS val 2005404103 ecr 1319203266], length 0
23:30:38.112488 IP serverarmour.com.57045 > 103.195.102.91.28016: Flags [.], ack 87, win 7631, options [nop,nop,TS val 1319203279 ecr 2005403999], length 0
23:30:38.113141 IP serverarmour.com.61849 > 103.195.102.91.28016: Flags [.], ack 646, win 846, options [nop,nop,TS val 1319203279 ecr 2005403999], length 0
23:30:38.129742 IP 172.93.111.152.56365 > 255.255.255.255.8888: UDP, length 32
23:30:38.135060 IP cloudpanel.cwgdata.center.ms-wbt-server > 194.0.234.40.52453: Flags [P.], seq 2799638900:2799638931, ack 2678015657, win 62720, length 31
23:30:38.141083 IP dns.google.domain > prox.anchored.host.44724: 30824 NXDomain 0/1/0 (106)
23:30:38.141210 IP prox.anchored.host.35349 > dns.google.domain: 31650+ PTR? 185.176.90.157.in-addr.arpa. (45)
23:30:38.147516 IP serverarmour.com.34204 > 103.195.102.91.28016: Flags [P.], seq 67:134, ack 1, win 10087, options [nop,nop,TS val 1319203314 ecr 2005404026], length 67
23:30:38.147614 IP 103.195.102.91.28016 > serverarmour.com.34204: Flags [.], ack 134, win 501, options [nop,nop,TS val 2005404151 ecr 1319203314], length 0
23:30:38.151746 IP cloudpanel.cwgdata.center.59301 > 104.20.1.160.https: Flags [SEW], seq 320572285, win 64240, options [mss 1460,nop,wscale 8,nop,nop,sackOK], length 0
23:30:38.152048 IP 104.20.1.160.https > cloudpanel.cwgdata.center.59301: Flags [S.], seq 3748945095, ack 320572286, win 65535, options [mss 1400,nop,nop,sackOK,nop,wscale 13], length 0
23:30:38.152129 IP cloudpanel.cwgdata.center.59301 > 104.20.1.160.https: Flags [.], ack 1, win 8192, length 0
23:30:38.152483 IP cloudpanel.cwgdata.center.59301 > 104.20.1.160.https: Flags [P.], seq 1:518, ack 1, win 8192, length 517
23:30:38.152747 IP 104.20.1.160.https > cloudpanel.cwgdata.center.59301: Flags [.], ack 518, win 9, length 0
23:30:38.155958 IP 104.20.1.160.https > cloudpanel.cwgdata.center.59301: Flags [P.], seq 1:2354, ack 518, win 9, length 2353
23:30:38.156018 IP cloudpanel.cwgdata.center.59301 > 104.20.1.160.https: Flags [.], ack 2354, win 8192, length 0
23:30:38.157754 IP cloudpanel.cwgdata.center.59301 > 104.20.1.160.https: Flags [P.], seq 518:598, ack 2354, win 8192, length 80
23:30:38.157910 IP cloudpanel.cwgdata.center.59301 > 104.20.1.160.https: Flags [P.], seq 598:1181, ack 2354, win 8192, length 583
23:30:38.158184 IP 104.20.1.160.https > cloudpanel.cwgdata.center.59301: Flags [.], ack 1181, win 9, length 0
23:30:38.167272 IP 80.64.30.82.30771 > cloudpanel.cwgdata.center.ms-wbt-server: Flags [R.], seq 1, ack 37, win 0, length 0
23:30:38.184423 IP 77.90.185.223.9599 > 103.195.102.88.ms-wbt-server: Flags [R], seq 3965235207, win 0, length 0
23:30:38.185169 IP syn-071-010-196-052.res.spectrum.com.51692 > prox.anchored.host.ssh: Flags [.], ack 2244, win 1021, length 0
23:30:38.194899 IP 103.195.102.91.28016 > serverarmour.com.4244: Flags [P.], seq 1:87, ack 67, win 501, options [nop,nop,TS val 2005404199 ecr 1319203266], length 86
23:30:38.194941 IP 103.195.102.91.28016 > serverarmour.com.34204: Flags [P.], seq 646:732, ack 134, win 501, options [nop,nop,TS val 2005404199 ecr 1319203314], length 86
23:30:38.196301 IP 77.90.185.223.36728 > 103.195.102.88.ms-wbt-server: Flags [.], ack 1, win 256, length 0
23:30:38.196320 IP 77.90.185.223.36728 > 103.195.102.88.ms-wbt-server: Flags [P.], seq 1:48, ack 1, win 256, length 47
23:30:38.196746 IP serverarmour.com.64447 > 103.195.102.91.28016: Flags [P.], seq 3253594008:3253594075, ack 2226835250, win 770, options [nop,nop,TS val 1319203363 ecr 2005403299], length 67
23:30:38.196826 IP 103.195.102.91.28016 > serverarmour.com.64447: Flags [.], ack 67, win 501, options [nop,nop,TS val 2005404201 ecr 1319203363], length 0
23:30:38.198916 IP 103.195.102.88.ms-wbt-server > 77.90.185.223.36728: Flags [P.], seq 1:20, ack 48, win 63953, length 19
23:30:38.203404 ARP, Request who-has 104.238.205.219 tell 104.238.205.1, length 46
23:30:38.209660 IP serverarmour.com.30692 > 103.195.102.91.28016: Flags [P.], seq 4288226784:4288226851, ack 1702484432, win 7807, options [nop,nop,TS val 1319203376 ecr 2005403299], length 67
23:30:38.209737 IP 103.195.102.91.28016 > serverarmour.com.30692: Flags [.], ack 67, win 501, options [nop,nop,TS val 2005404214 ecr 1319203376], length 0
23:30:38.210222 IP serverarmour.com.6904 > 103.195.102.91.28016: Flags [.], ack 87, win 825, options [nop,nop,TS val 1319203377 ecr 2005404099], length 0
23:30:38.210291 IP serverarmour.com.52832 > 103.195.102.91.28016: Flags [.], ack 646, win 737, options [nop,nop,TS val 1319203376 ecr 2005404099], length 0
23:30:38.210383 IP serverarmour.com.6904 > 103.195.102.91.28016: Flags [P.], seq 67:134, ack 87, win 825, options [nop,nop,TS val 1319203377 ecr 2005404099], length 67
23:30:38.210439 IP 103.195.102.91.28016 > serverarmour.com.6904: Flags [.], ack 134, win 501, options [nop,nop,TS val 2005404214 ecr 1319203377], length 0
23:30:38.212896 IP serverarmour.com.34204 > 103.195.102.91.28016: Flags [.], ack 646, win 10087, options [nop,nop,TS val 1319203380 ecr 2005404099], length 0
23:30:38.213168 IP serverarmour.com.27947 > 103.195.102.91.28016: Flags [.], ack 646, win 3247, options [nop,nop,TS val 1319203380 ecr 2005404100], length 0
23:30:38.214407 ARP, Request who-has 103.195.101.223 tell mail.payroll.ph, length 46
23:30:38.214421 ARP, Request who-has unassigned.swiftnode.net tell 104.238.205.1, length 46
23:30:38.222877 IP 185.42.12.58.1817 > cloudpanel.cwgdata.center.ms-wbt-server: Flags [P.], seq 193:511, ack 866, win 1023, length 318
23:30:38.223719 IP cloudpanel.cwgdata.center.ms-wbt-server > 185.42.12.58.1817: Flags [P.], seq 866:917, ack 511, win 63490, length 51
^C^C23:30:38.230031 IP 112.239.142.5.9422 > 103.195.102.90.telnet: Flags [S], seq 1740858970, win 16806, length 0