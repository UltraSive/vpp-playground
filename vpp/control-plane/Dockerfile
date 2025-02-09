FROM debian:bookworm

# Add VPP repository and install required packages
RUN apt update && apt install -y curl gnupg procps && \
    curl -s https://packagecloud.io/install/repositories/fdio/release/script.deb.sh | bash

# Install VPP with proper configuration handling
RUN DEBIAN_FRONTEND=noninteractive apt update && \
    mkdir -p /usr/lib/sysctl.d && \
    touch /usr/lib/sysctl.d/50-pid-max.conf && \
    touch /etc/sysctl.d/80-vpp.conf && \
    DEBIAN_FRONTEND=noninteractive apt install -y vpp || true && \
    rm -f /etc/sysctl.d/80-vpp.conf

# Set entrypoint to vppctl
CMD ["/usr/bin/vppctl"]