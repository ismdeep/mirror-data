package task

import "testing"

func TestElectronSsrBackup_Run(t *testing.T) {
	receiver := &ElectronSsrBackup{}
	receiver.Run()
}
