FROM ubuntu:20.04

RUN apt-get update && apt-get install -y \
    curl \
    jq \
    net-tools \
    vim \
    git

RUN sh -ci "$(curl -fsSL https://storage.googleapis.com/flow-cli/install.sh)"
