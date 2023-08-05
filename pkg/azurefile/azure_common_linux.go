//go:build linux
// +build linux

/*
Copyright 2020 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package azurefile

import (
	"os"

	"k8s.io/klog/v2"
	mount "k8s.io/mount-utils"
)

func SMBMount(m *mount.SafeFormatAndMount, source, target, fsType string, options, sensitiveMountOptions []string) error {
	return m.MountSensitive(source, target, fsType, options, sensitiveMountOptions)
}

func SMBUnmount(m *mount.SafeFormatAndMount, target string) error {
	return m.Unmount(target)
}

func RemoveStageTarget(m *mount.SafeFormatAndMount, target string) error {
	return os.Remove(target)
}

func CleanupSMBMountPoint(m *mount.SafeFormatAndMount, target string, extensiveMountCheck bool) error {
	// unmount first since if remote SMB directory is not found, linked path cannot be deleted if not mounted
	if err := m.Unmount(target); err != nil {
		klog.Errorf("Unmount(%s) failed with %v", target, err)
	}
	return CleanupMountPoint(m, target, extensiveMountCheck)
}

func CleanupMountPoint(m *mount.SafeFormatAndMount, target string, extensiveMountCheck bool) error {
	return mount.CleanupMountPoint(target, m, extensiveMountCheck)
}

func preparePublishPath(path string, m *mount.SafeFormatAndMount) error {
	return nil
}

func prepareStagePath(path string, m *mount.SafeFormatAndMount) error {
	return nil
}
