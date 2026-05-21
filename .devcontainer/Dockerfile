# Use Ubuntu as base OS
FROM ubuntu:24.04

# Install tools
RUN apt-get update
RUN apt-get install -y procps curl
RUN apt-get install wget -y

# Install exercism CLI
RUN wget https://github.com/exercism/cli/releases/download/v3.5.6/exercism-3.5.6-linux-arm64.tar.gz
RUN tar -xf exercism-3.5.6-linux-arm64.tar.gz
RUN mv exercism /usr/bin
RUN exercism configure --token=d1818b69-752e-41ef-8064-f6444b01991b
RUN exercism configure --workspace /development/

# Install GO
RUN wget https://go.dev/dl/go1.26.3.linux-arm64.tar.gz
RUN rm -rf /usr/local/go && tar -C /usr/local -xzf go1.26.3.linux-arm64.tar.gz


# Get and Install Node
# https://github.com/nodejs/node
RUN curl -sL https://deb.nodesource.com/setup_20.x | bash -
RUN apt-get update
RUN apt-get install -y nodejs

# Install Git
RUN apt-get install -y git
RUN apt-get install git-lfs

# Install GitHub CLI
# https://github.com/cli/cli
RUN curl -fsSL https://cli.github.com/packages/githubcli-archive-keyring.gpg | dd of=/etc/apt/trusted.gpg.d/githubcli-archive-keyring.gpg
RUN echo "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/trusted.gpg.d/githubcli-archive-keyring.gpg] https://cli.github.com/packages stable main" | tee /etc/apt/sources.list.d/github-cli.list >/dev/null
RUN apt-get update
RUN apt-get install gh

# Install Shell
RUN mkdir install
COPY /shell /install/shell
RUN chmod +x /install/shell/install.sh
RUN /bin/bash -c "./install/shell/install.sh"