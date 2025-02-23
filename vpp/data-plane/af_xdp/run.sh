docker run -d --name data-plane \
    --privileged \
    --network=host \
    -v /run/vpp/:/run/vpp/ \
    vpp-container