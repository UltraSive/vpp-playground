# Enable promiscuous mode on host interface in order to directly recieve packets
ip link set dev ens20 promisc on up

# Create the interface
vppctl create int af_xdp host-if ens20 num-rx-queues all
---
vppctl create int af_xdp host-if ens20 num-rx-queues all prog /extras/bpf/af_xdp.bpfel.o

# 1. Enable and Configure the Gigabit Ethernet Interface
vppctl set interface mac address ens20/0 BC:24:11:ED:08:CA
# vppctl set interface mtu 1500 ens20/0
vppctl set interface ip address ens20/0 103.195.102.93/24
vppctl ip route add 0.0.0.0/0 via 103.195.102.1
vppctl set interface state ens20/0 up

# 2. Create and Configure the GRE Tunnel
vppctl create gre tunnel src 172.93.110.120 dst 45.76.233.197

vppctl set interface ip address gre0 10.10.10.1/30
vppctl set interface state gre0 up

# 3. Add Static Route for Destination Subnet via GRE Tunnel
vppctl ip route add 103.195.102.92/24 via 10.10.10.2 gre0