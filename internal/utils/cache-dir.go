package utils

import (
	"os"
	"path"
)

func VulnmapCacheDir() (string, error) {
	osutil := NewVulnmapOSUtil()
	return VulnmapCacheDirImpl(osutil)
}

func VulnmapCacheDirImpl(osUtil VulnmapOSUtil) (string, error) {
	baseDirectory, err := osUtil.UserCacheDir()
	subDir := path.Join("vulnmap", "vulnmap-cli")
	vulnmapCacheDir := path.Join(baseDirectory, subDir)

	err2 := os.MkdirAll(vulnmapCacheDir, FILEPERM_755)
	if err2 != nil {
		// Returning "vulnmap-cli" to be used as the cache directory name later.
		return subDir, err2
	}

	return vulnmapCacheDir, err
}

func FullPathInVulnmapCacheDir(cacheDir string, filename string) (string, error) {
	return path.Join(cacheDir, filename), nil
}
