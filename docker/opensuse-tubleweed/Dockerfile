FROM opensuse/tumbleweed
MAINTAINER M0Rf30

RUN zypper update \
&& zypper -n install -y tar expect rpm-build rsync make automake gcc wget zip git bzr mercurial unzip \
&& zypper -n install -y go

ENV GOPATH /go
ENV PATH /usr/local/go/bin:$PATH:/go/bin

RUN go get github.com/packagefoundation/yap

ENTRYPOINT ["yap"]
CMD ["build", "opensuse"]
