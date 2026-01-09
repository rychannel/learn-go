# Use Debian as the base image
FROM debian:bookworm-slim

# Avoid prompts from apt
ENV DEBIAN_FRONTEND=noninteractive

# Update package lists and install essential tools including Go
RUN apt-get update && apt-get upgrade -y && apt-get install -y \
    curl \
    wget \
    git \
    ca-certificates \
    gnupg \
    build-essential \
    golang-go \
    && rm -rf /var/lib/apt/lists/*

# Set Go environment variables for GOPATH workflow
ENV GOPATH="/workspaces/learn-go"
ENV GO111MODULE=off
ENV PATH="${GOPATH}/bin:${GOROOT}/bin:${PATH}"

# Create a non-root user for development
RUN useradd -m -s /bin/bash -u 1000 vscode

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
