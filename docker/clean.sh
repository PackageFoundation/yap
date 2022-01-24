sudo docker rm "$(sudo docker ps -a -q)"

sudo docker rmi packagefoundation/yap-archlinux
sudo docker rmi packagefoundation/yap-amazonlinux-1
sudo docker rmi packagefoundation/yap-amazonlinux-2
sudo docker rmi packagefoundation/yap-centos-8
sudo docker rmi packagefoundation/yap-debian-jessie
sudo docker rmi packagefoundation/yap-debian-stretch
sudo docker rmi packagefoundation/yap-debian-buster
sudo docker rmi packagefoundation/yap-fedora-32
sudo docker rmi packagefoundation/yap-genkey
sudo docker rmi packagefoundation/yap-opensuse-tumbleweed
sudo docker rmi packagefoundation/yap-ubuntu-bionic
sudo docker rmi packagefoundation/yap-ubuntu-focal

sudo docker rmi base/archlinux
sudo docker rmi archlinux/base
sudo docker rmi amazonlinux:1
sudo docker rmi amazonlinux:2
sudo docker rmi centos:8
sudo docker rmi debian:jessie
sudo docker rmi debian:stretch
sudo docker rmi debian:buster
sudo docker rmi fedora:32
sudo docker rmi opensuse:tumbleweed
sudo docker rmi oraclelinux:8
sudo docker rmi ubuntu:bionic
sudo docker rmi ubuntu:focal

sudo docker rmi "$(sudo docker images -q -f dangling=true)"
