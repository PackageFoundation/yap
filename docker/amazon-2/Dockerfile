FROM amazonlinux:2
LABEL maintainer="Yap <contact@packagefoundation.org>"

SHELL ["/bin/bash", "-o", "pipefail", "-c"]
RUN \
  yum -y install yum-plugins-core epel-release; \
  yum config-manager --set-enabled powertools; \
  yum config-manager --set-enabled appstream; \
  yum -y groupinstall 'Development Tools'; \
  yum -y install \
  automake bash-completion createrepo expect gcc git make openssl rpm-build rpm-sign rsync sudo tar zip; \
  curl -sL https://dl.google.com/go/go1.19.2.linux-amd64.tar.gz --output go.tar.gz; \
  echo "5e8c5a74fe6470dd7e055a461acda8bb4050ead8c2df70f227e3ff7d8eb7eeb6 go.tar.gz" | sha256sum -c - ; \
  tar -C /usr/lib -xf go.tar.gz; \
  rm -f go.tar.gz; \
  ln -sf /usr/lib/go/bin/go /usr/bin/go; \
  ln -sf /usr/lib/go/bin/gofmt /usr/bin/gofmt; \
  ln -sf /usr/share/doc/go /usr/lib/go/doc; \
  yum clean all

ENV GOROOT=/usr/lib/go
ENV PATH /usr/lib/go/bin:$PATH:/go/bin

COPY . /project
WORKDIR /project
RUN \
  go build -ldflags "-w -s" -trimpath -mod=readonly -modcacherw -o /usr/bin/; \
  yap completion bash > /etc/bash_completion.d/yap; \
  echo "source /usr/share/bash-completion/bash_completion" >> /etc/bashrc; \
  rm -rf /project

ENTRYPOINT ["yap"]
