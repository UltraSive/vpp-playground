docker run -d --name data-plane \
    --privileged \
    -v /run/vpp/:/run/vpp/ \
    vpp-container