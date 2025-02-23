FROM debian:bullseye

# Install required dependencies
RUN apt update && apt install -y \
    software-properties-common \
    curl gnupg lsb-release \
    && rm -rf /var/lib/apt/lists/*

# Add the FD.io repository key
RUN curl -fsSL https://packagecloud.io/fdio/release/gpgkey | gpg --dearmor -o /usr/share/keyrings/fdio-archive-keyring.gpg

# Add FD.io repository
RUN echo "deb [signed-by=/usr/share/keyrings/fdio-archive-keyring.gpg] https://packagecloud.io/fdio/release/debian $(lsb_release -cs) main" | tee /etc/apt/sources.list.d/fdio.list

# Ensure dependencies are installed first
RUN apt update && apt install -y procps

# Install VPP and necessary plugins
RUN DEBIAN_FRONTEND=noninteractive apt update && \
    mkdir -p /usr/lib/sysctl.d && \
    touch /usr/lib/sysctl.d/50-pid-max.conf && \
    touch /etc/sysctl.d/80-vpp.conf && \
    groupadd -r vpp && \
    useradd -r -g vpp vpp && \
    DEBIAN_FRONTEND=noninteractive apt install -y --no-install-recommends \
    vpp vpp-plugin-core vpp-plugin-dpdk || true && \
    rm -f /etc/sysctl.d/80-vpp.conf

# Create necessary directories
RUN mkdir -p /etc/vpp /var/log/vpp /run/vpp && \
    chown -R vpp:vpp /var/log/vpp /run/vpp

# Copy the VPP startup configuration file
COPY dpdk-startup.conf /etc/vpp/startup.conf

# Default command to start VPP
CMD ["/usr/bin/vpp", "-c", "/etc/vpp/startup.conf"]