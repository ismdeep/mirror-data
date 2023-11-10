package task

import "github.com/ismdeep/mirror-data/internal/github"

type EtcdManager struct {
}

func (receiver *EtcdManager) Run() {
	github.FetchReleases("etcd-manager", "gtamas", "etcdmanager")
}
