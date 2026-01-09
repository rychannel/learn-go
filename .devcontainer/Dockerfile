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
ENV GO_CHECKSUM=e2bc0b3e4b64111ec117295c088bde5f00eeed1567999ff77bc859d7df70078e

# Download and install Go
RUN wget -q https://go.dev/dl/go${GO_VERSION}.linux-amd64.tar.gz && \
    echo "${GO_CHECKSUM}  go${GO_VERSION}.linux-amd64.tar.gz" | sha256sum -c - && \
    tar -C /usr/local -xzf go${GO_VERSION}.linux-amd64.tar.gz && \
    rm go${GO_VERSION}.linux-amd64.tar.gz

# Set Go environment variables
ENV PATH="/usr/local/go/bin:${PATH}"
ENV GOPATH="/go"
ENV PATH="${GOPATH}/bin:${PATH}"

# Create GOPATH directory
RUN mkdir -p ${GOPATH}/src ${GOPATH}/bin

# Create a non-root user for development
RUN useradd -m -s /bin/bash -u 1000 vscode && \
    chown -R vscode:vscode ${GOPATH}

# Verify Go installation
RUN go version

# Set the working directory
WORKDIR /workspace

# Change ownership of workspace
RUN chown vscode:vscode /workspace

# Reset DEBIAN_FRONTEND
ENV DEBIAN_FRONTEND=

# Default command
CMD ["/bin/bash"]
