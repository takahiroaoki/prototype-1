FROM amazonlinux:2023.1.20230809.0

# dnf fixes tool versions
RUN dnf install -y \
     git \
     sudo \
     && rm -rf /var/cache/dnf/* \
     && dnf clean all

# useradd
RUN useradd "dev-user" && echo "dev-user ALL=NOPASSWD: ALL" >> /etc/sudoers
USER dev-user
