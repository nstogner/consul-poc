# Consul Proof of Concept

## Resources

https://medium.com/zendesk-engineering/making-docker-and-consul-get-along-5fceda1d52b9

### 1. Spin up new VM with Ubuntu 16.04

### 2. Install docker

```
sudo apt-get update

sudo apt-get install \
    apt-transport-https \
    ca-certificates \
    curl \
    software-properties-common

curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -

sudo add-apt-repository \
   "deb [arch=amd64] https://download.docker.com/linux/ubuntu \
   $(lsb_release -cs) \
   stable"

sudo apt-get update

sudo apt-get install docker-ce
```

### 3. Run consul

```
# sudo docker run -d --net=host consul:latest
sudo docker run --net=host consul:latest consul agent --client=169.254.1.1 --dev
```

### 4. Setup Dummy Interface

```
$ sudo ip link add dummy0 type dummy
$ sudo ip link set dev dummy0 up
$ ip link show type dummy
25: dummy0: <BROADCAST,NOARP> mtu 1500 qdisc noop state DOWN mode DEFAULT qlen 1000
    link/ether 2a:bb:3e:f6:50:1c brd ff:ff:ff:ff:ff:ff
```

```
$ sudo ip addr add 169.254.1.1/32 dev dummy0 
$ sudo ip link set dev dummy0 up
$ ip addr show dev dummy0
25: dummy0: <BROADCAST,NOARP,UP,LOWER_UP> mtu 1500 qdisc noqueue state UNKNOWN qlen 1000
    link/ether 2a:bb:3e:f6:50:1c brd ff:ff:ff:ff:ff:ff
    inet 169.254.1.1/32 scope global dummy0
       valid_lft forever preferred_lft forever
    inet6 fe80::28bb:3eff:fef6:501c/64 scope link 
       valid_lft forever preferred_lft forever
```

### 5. Configure interface

Place the following file into `/etc/systemd/network/dummy0.netdev`:

```
[NetDev]
Name=dummy0
Kind=dummy
```

Then place the following file into `/etc/systemd/network/dummy0.network`:

```
[Match]
Name=dummy0

[Network]
Address=169.254.1.1/32
```

`sudo systemctl restart systemd-networkd`

### 6. Install & Setup dnsmasq

`sudo apt-get install dnsmasq`

Create `/etc/dnsmasq.d/consul.conf`:

```
server=/consul/169.254.1.1#8600
listen-address=127.0.0.1
listen-address=169.254.1.1
```

