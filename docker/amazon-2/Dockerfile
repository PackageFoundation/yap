FROM amazonlinux:2
MAINTAINER Yap <contact@packagefoundation.org>

RUN rpm -Uvh https://dl.fedoraproject.org/pub/epel/epel-release-latest-7.noarch.rpm
RUN yum -y upgrade
RUN yum -y groupinstall 'Development Tools'
RUN yum -y install tar expect rpm-build rpm-sign openssl createrepo rsync make automake gcc wget zip git mercurial

RUN wget https://dl.google.com/go/go1.17.2.linux-amd64.tar.gz
RUN echo "f242a9db6a0ad1846de7b6d94d507915d14062660616a61ef7c808a76e4f1676 go1.17.2.linux-amd64.tar.gz" | sha256sum -c -
RUN tar -C /usr/local -xf go1.17.2.linux-amd64.tar.gz
RUN rm -f go1.17.2.linux-amd64.tar.gz

ENV GOPATH /go
ENV PATH /usr/local/go/bin:$PATH:/go/bin

RUN go get github.com/packagefoundation/yap

ENTRYPOINT ["yap"]
CMD ["build", "amazonlinux-2"]