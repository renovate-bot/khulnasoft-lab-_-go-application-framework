package utils

import (
	"log"
	"path"
)

// Gets the system temp directory and, if it doesn't exist, attempts to create it.
func systemTempDirectory(debugLogger *log.Logger, osUtil VulnmapOSUtil) (string, error) {
	tempDir := osUtil.TempDir()
	// make sure this directory exists
	debugLogger.Println("system temp directory:", tempDir)
	_, err := osUtil.Stat(tempDir)
	if err != nil {
		debugLogger.Println("system temp directory does not exist... attempting to create it:", tempDir)
		err = osUtil.MkdirAll(tempDir, FILEPERM_755)
		if err != nil {
			debugLogger.Println("failed to create system temp directory:", tempDir)
			return "", err
		}
	}

	return tempDir, nil
}

func vulnmapTempDirectoryImpl(debugLogger *log.Logger, osUtil VulnmapOSUtil) (string, error) {
	tempDir, err := systemTempDirectory(debugLogger, osUtil)
	if err != nil {
		return "", err
	}

	vulnmapTempDir := path.Join(tempDir, "vulnmap")

	// make sure it exists
	_, err = osUtil.Stat(vulnmapTempDir)
	if err != nil {
		debugLogger.Println("vulnmap temp directory does not exist... attempting to create it:", vulnmapTempDir)
		err = osUtil.MkdirAll(vulnmapTempDir, FILEPERM_755)
		if err != nil {
			debugLogger.Println("failed to create vulnmap temp directory:", vulnmapTempDir)
			return "", err
		}
	}

	return vulnmapTempDir, nil
}

func VulnmapTempDirectory(debugLogger *log.Logger) (string, error) {
	osutil := NewVulnmapOSUtil()
	return vulnmapTempDirectoryImpl(debugLogger, osutil)
}
