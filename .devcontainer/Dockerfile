FROM mcr.microsoft.com/devcontainers/base:noble

RUN apt update && apt upgrade -y

RUN apt install -y mysql-client-core-8.0

RUN curl -fsSL https://deb.nodesource.com/setup_26.x | bash \
  && apt install -y nodejs

USER vscode
