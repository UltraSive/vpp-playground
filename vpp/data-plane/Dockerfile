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

# Install VPP and necessary plugins
RUN apt update && apt install -y \
    vpp vpp-plugin-core vpp-plugin-dpdk vpp-plugin-memif \
    && rm -rf /var/lib/apt/lists/*

# Default command to start VPP
CMD ["/usr/bin/vpp", "-c", "/etc/vpp/startup.conf"]