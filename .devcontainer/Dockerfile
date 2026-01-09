# Use Debian as the base image
FROM debian:bookworm-slim

# Avoid prompts from apt
ENV DEBIAN_FRONTEND=noninteractive

# Update package lists and install essential tools
RUN apt-get update && apt-get install -y \
    curl \
    wget \
    git \
    ca-certificates \
    gnupg \
    build-essential \
    && rm -rf /var/lib/apt/lists/*

# Set Go version
ENV GO_VERSION=1.21.5

# Download and install Go
RUN wget -q https://go.dev/dl/go${GO_VERSION}.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go${GO_VERSION}.linux-amd64.tar.gz && \
    rm go${GO_VERSION}.linux-amd64.tar.gz

# Set Go environment variables
ENV PATH="/usr/local/go/bin:${PATH}"
ENV GOPATH="/go"
ENV PATH="${GOPATH}/bin:${PATH}"

# Create GOPATH directory
RUN mkdir -p ${GOPATH}/src ${GOPATH}/bin

# Verify Go installation
RUN go version

# Set the working directory
WORKDIR /workspace

# Reset DEBIAN_FRONTEND
ENV DEBIAN_FRONTEND=

# Default command
CMD ["/bin/bash"]
