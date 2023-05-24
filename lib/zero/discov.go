package zero

import "github.com/zeromicro/go-zero/core/discov"

type (
	// EtcdConf is the config item with the given key on etcd.
	// Hosts etcd host array 127.0.0.1:2379
	// Key rpc registration key
	EtcdConf = discov.EtcdConf
)
