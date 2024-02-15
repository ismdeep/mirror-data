package task

import (
	"github.com/ismdeep/mirror-data/app/data/internal/github"
)

type EtcdManager struct {
}

func (receiver *EtcdManager) Run() {
	github.FetchReleases("etcd-manager", "gtamas", "etcdmanager", nil)
}
