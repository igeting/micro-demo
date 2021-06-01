# tools

## etcd
```
#stop firewall
systemctl stop firewalld
systemctl disable firewalld
systemctl status  firewalld
#start etcd
./etcd --listen-client-urls http://0.0.0.0:2379 --advertise-client-urls http://0.0.0.0:2379 --listen-peer-urls http://0.0.0.0:2380
```

## consul(default listen :8500)
```
#stop firewall
systemctl stop firewalld
systemctl disable firewalld
systemctl status  firewalld
#start consul(-dev & -server)
./consul agent -dev -client 0.0.0.0 -ui
```

## consul cluster(default listen :8500)
```
#cluster
consul agent -server -bootstrap-expect=3 -data-dir=/tmp/consul -node=10.200.110.91 -bind=10.200.110.91 -client=0.0.0.0 -datacenter=origin -ui
consul agent -server -bootstrap-expect=3 -data-dir=/tmp/consul -node=10.200.110.92 -bind=10.200.110.92 -client=0.0.0.0 -datacenter=origin -ui
consul agent -server -bootstrap-expect=3 -data-dir=/tmp/consul -node=10.200.110.93 -bind=10.200.110.93 -client=0.0.0.0 -datacenter=origin -ui
```
- server run as server
- bootstrap-expect cluster number
- data-dir data dir
- node node id
- bind listen ip, default 0.0.0.0
- client client ip, remote use 0.0.0.0
- ui view ui
- config-dir config dir
- datacenter center name
