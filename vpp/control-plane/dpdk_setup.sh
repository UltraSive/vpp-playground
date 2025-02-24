# 1. Enable and Configure the Gigabit Ethernet Interface
vppctl set interface mac address GigabitEthernet0/13/0 bc:24:11:2c:bd:9a
vppctl set interface mtu 1500 GigabitEthernet0/13/0
vppctl set interface ip address GigabitEthernet0/13/0 172.93.110.120/24
vppctl ip route add 0.0.0.0/0 via 172.93.110.1
vppctl set interface state GigabitEthernet0/13/0 up

# 2. Create and Configure the GRE Tunnel
vppctl create gre tunnel src 172.93.110.120 dst 45.76.233.197

vppctl set interface ip address gre0 10.10.10.1/30
vppctl set interface state gre0 up

# 3. Add Static Route for Destination Subnet via GRE Tunnel
vppctl ip route add 103.195.102.92/24 via 10.10.10.2 gre0