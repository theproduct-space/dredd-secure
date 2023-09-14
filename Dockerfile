FROM --platform=linux ubuntu:22.04
ARG BUILDARCH

# Change your versions here
ENV GO_VERSION=1.20.3
ENV IGNITE_VERSION=0.27.1
ENV NODE_VERSION=18.x

ENV LOCAL=/usr/local
ENV GOROOT=$LOCAL/go
ENV HOME=/root
ENV GOPATH=$HOME/go
ENV PATH=$GOROOT/bin:$GOPATH/bin:$PATH

RUN mkdir -p $GOPATH/bin

ENV PACKAGES curl gcc jq nano
RUN apt-get update
RUN apt-get install -y $PACKAGES

# Install Go
RUN curl -L https://go.dev/dl/go${GO_VERSION}.linux-$BUILDARCH.tar.gz | tar -C $LOCAL -xzf -

# Install Ignite
RUN curl -L https://get.ignite.com/cli@v${IGNITE_VERSION}! | bash

# Install Node
RUN curl -fsSL https://deb.nodesource.com/setup_${NODE_VERSION} | bash -
RUN apt-get install -y nodejs

# Install Hermes
RUN mkdir -p $HOME/.hermes/bin
RUN curl -L https://github.com/informalsystems/hermes/releases/download/v1.6.0/hermes-v1.6.0-x86_64-unknown-linux-gnu.tar.gz | tar -C $HOME/.hermes/bin/ -vxzf -
RUN echo 'export PATH="$HOME/.hermes/bin:$PATH"' >> $HOME/.bashrc

EXPOSE 1317 3000 4500 9090 26657

WORKDIR /dredd-secure

# COPY go.mod /dredd-secure/go.mod
# RUN go mod download
# RUN rm /dredd-secure/go.mod