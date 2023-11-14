// OS interface wrappers

package utils

import "os"

type VulnmapOSUtil interface {
	UserCacheDir() (string, error)
	MkdirAll(path string, perm os.FileMode) error
	Stat(name string) (os.FileInfo, error)
	TempDir() string
}

type vulnmapOsUtilImpl struct{}

func (bd *vulnmapOsUtilImpl) UserCacheDir() (string, error) {
	return os.UserCacheDir()
}

func (bd *vulnmapOsUtilImpl) MkdirAll(path string, perm os.FileMode) error {
	return os.MkdirAll(path, perm)
}

func (bd *vulnmapOsUtilImpl) Stat(name string) (os.FileInfo, error) {
	return os.Stat(name)
}

func (bd *vulnmapOsUtilImpl) TempDir() string {
	return os.TempDir()
}

func NewVulnmapOSUtil() VulnmapOSUtil {
	return &vulnmapOsUtilImpl{}
}
