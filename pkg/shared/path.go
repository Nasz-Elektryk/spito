package shared

import (
	"os"
	"os/user"
	"strings"
)

const (
	DirectoryPermissions = 0755
	FilePermissions      = 0644
)

func DoesPathExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func ExpandTilde(path *string) error {
	usr, err := user.Current()
	if err != nil {
		return err
	}

	if *path == "~" {
		*path = usr.HomeDir
	}
	if strings.HasPrefix(*path, "~/") {
		*path = strings.Replace(*path, "~", usr.HomeDir, 1)
	}
	return nil
}

func GetEnvWithDefaultValue(environmentVariable string, defaultValue string) string {
	if val := os.Getenv(environmentVariable); val != "" {
		return val
	}
	return defaultValue
}