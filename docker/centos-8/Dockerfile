FROM centos:8
MAINTAINER Yap <contact@packagefoundation.org>

RUN \
  yum -y install epel-release; \
  sed -i "s/enabled=0/enabled=1/g" /etc/yum.repos.d/CentOS-Linux-PowerTools.repo; \
  sed -i "s/enabled=0/enabled=1/g" /etc/yum.repos.d/CentOS-Linux-AppStream.repo; \
  yum -y upgrade; \
  yum -y groupinstall 'Development Tools'; \
  yum -y install \
  bash-completion tar expect rpm-build rpm-sign openssl createrepo rsync make automake gcc zip git mercurial sudo \
  curl -sL https://golang.org/dl/go1.17.2.linux-amd64.tar.gz --output go.tar.gz; \
  echo "f242a9db6a0ad1846de7b6d94d507915d14062660616a61ef7c808a76e4f1676 go.tar.gz" | sha256sum -c - ; \
  tar -C /usr/lib -xf go.tar.gz; \
  rm -f go.tar.gz; \
  ln -sf /usr/lib/go/bin/go /usr/bin/go; \
  ln -sf /usr/lib/go/bin/gofmt /usr/bin/gofmt; \
  ln -sf /usr/share/doc/go /usr/lib/go/doc

ENV GOROOT=/usr/lib/go
ENV PATH /usr/lib/go/bin:$PATH:/go/bin

RUN go get github.com/packagefoundation/yap; \
 yap completion bash > /etc/bash_completion.d/yap; \
 echo "source /usr/share/bash-completion/bash_completion" >> /etc/bashrc \

ENTRYPOINT ["yap"]
CMD ["build", "centos-8"]
