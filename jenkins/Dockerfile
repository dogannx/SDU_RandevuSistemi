FROM jenkins/jenkins:lts-jdk17

USER root

ARG GO_VERSION=1.25.0
ARG NODE_MAJOR=22

# Go (mimari otomatik tespit: amd64 / arm64)
RUN ARCH=$(dpkg --print-architecture) \
    && curl -fsSL "https://go.dev/dl/go${GO_VERSION}.linux-${ARCH}.tar.gz" -o /tmp/go.tgz \
    && tar -C /usr/local -xzf /tmp/go.tgz \
    && rm /tmp/go.tgz

# Node.js
RUN curl -fsSL "https://deb.nodesource.com/setup_${NODE_MAJOR}.x" | bash - \
    && apt-get install -y nodejs \
    && apt-get clean && rm -rf /var/lib/apt/lists/*

ENV PATH="/usr/local/go/bin:${PATH}"
ENV GOPATH=/var/jenkins_home/go
ENV GOCACHE=/var/jenkins_home/.cache/go-build

# Pluginleri önceden yükle
COPY plugins.txt /usr/share/jenkins/ref/plugins.txt
RUN jenkins-plugin-cli --plugin-file /usr/share/jenkins/ref/plugins.txt

# Kurulum sihirbazını atla + admin/job oluşturan init scriptleri
ENV JAVA_OPTS="-Djenkins.install.runSetupWizard=false"
COPY init.groovy.d/ /usr/share/jenkins/ref/init.groovy.d/

USER jenkins
