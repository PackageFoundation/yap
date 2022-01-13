FROM rockylinux/rockylinux:8
MAINTAINER Pacur <contact@pacur.org>

RUN \
  dnf -y install dnf-plugins-core epel-release; \
  dnf config-manager --set-enabled powertools; \
  dnf config-manager --set-enabled appstream; \
  dnf -y upgrade; \
  dnf -y groupinstall 'Development Tools'; \
  dnf -y install \
  bash-completion tar expect rpm-build rpm-sign openssl createrepo rsync make automake gcc zip git mercurial sudo \
  java-1.8.0-openjdk-headless.x86_64; \
  curl -sL https://golang.org/dl/go1.17.2.linux-amd64.tar.gz --output go.tar.gz; \
  echo "f242a9db6a0ad1846de7b6d94d507915d14062660616a61ef7c808a76e4f1676 go.tar.gz" | sha256sum -c - ; \
  tar -C /usr/lib -xf go.tar.gz; \
  rm -f go.tar.gz; \
  ln -sf /usr/lib/go/bin/go /usr/bin/go; \
  ln -sf /usr/lib/go/bin/gofmt /usr/bin/gofmt; \
  ln -sf /usr/share/doc/go /usr/lib/go/doc

ENV GOROOT=/usr/lib/go

# Custom stuff from here
ADD pacur /usr/bin/

ENV \
# These variables are saved as ENV 'cause they can be useful to future
# scripts
    CI_USER=agent \
    CI_UID=1000 \
    CI_GID=1000 \
    JENKINS_AGENT=4.2 \
    JAVA_OPTS="-XX:+UnlockExperimentalVMOptions -XX:+UseCGroupMemoryLimitForHeap -Dfile.encoding=UTF8" \
    TIMEZONE="Europe/Rome" \
    LC_ALL="C.UTF-8"

RUN pacur completion bash > /etc/bash_completion.d/pacur \
 && echo "source /usr/share/bash-completion/bash_completion" >> /etc/bashrc \
 && mkdir /root/.autosetup \
 && mkdir /usr/share/zextras-ci-agent \
 && curl -sL https://sentry.io/get-cli/ | bash \
 && curl --create-dirs -fsSLo /usr/share/jenkins/slave.jar https://repo.jenkins-ci.org/public/org/jenkins-ci/main/remoting/${JENKINS_AGENT}/remoting-${JENKINS_AGENT}.jar \
 && chmod 755 /usr/share/jenkins \
 && chmod 644 /usr/share/jenkins/slave.jar \
 && groupadd -r -g 233 docker \
 && groupadd -r -g ${CI_GID} ${CI_USER} \
 && useradd -m -u ${CI_UID} -g ${CI_GID} -G docker -s /bin/bash ${CI_USER} \
 #&& usermod -aG sudo ${CI_USER} \
 && echo %${CI_USER} ALL=NOPASSWD:ALL > /etc/sudoers.d/${CI_USER} \
 && chmod 0440 /etc/sudoers.d/${CI_USER} \
 && echo "pacur build centos-8" > /usr/bin/pacur-build \
 && chmod a+x /usr/bin/pacur-build

ADD docker/jenkins/launcher.sh /usr/bin/launcher
ADD docker/jenkins/logging.sh /usr/share/zextras-ci-agent/logging.sh
ADD docker/jenkins/jenkins.sh /usr/share/zextras-ci-agent/jenkins.sh
ADD docker/jenkins/utils.sh /usr/share/zextras-ci-agent/utils.sh


ENTRYPOINT ["/usr/bin/launcher"]
