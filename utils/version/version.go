package version

import (
	"errors"
	"os"
	"os/exec"
	"strings"

	"gopkg.in/ini.v1"

	"github.com/BeesNestInc/CassetteOS-Common/utils/file"
	_ "github.com/mattn/go-sqlite3" // nolint
)

const (
	LegacyCassetteOSServiceName = "cassetteos.service"
	configKeyUniqueToZero3x = "USBAutoMount"
	configKeyDBPath         = "DBPath"
)

var (
	// this value will be updated at init() to actual config file path.
	LegacyCassetteOSConfigFilePath = "/etc/cassetteos.conf"

	_configFile        *ini.File
	_cassetteOSBinFilePath string
)

var (
	ErrLegacyVersionNotFound = errors.New("legacy version not found")
	ErrVersionNotFound       = errors.New("version (non-legacy) not found")
)

func init() {
	serviceFilePath := file.FindFirstFile("/etc/systemd", LegacyCassetteOSServiceName)
	if serviceFilePath == "" {
		return
	}

	serviceFile, err := ini.Load(serviceFilePath)
	if err != nil {
		return
	}

	section, err := serviceFile.GetSection("Service")
	if err != nil {
		return
	}

	key, err := section.GetKey("ExecStart")
	if err != nil {
		return
	}

	execStart := key.Value()
	texts := strings.Split(execStart, " ")

	// locaste cassetteos binary.
	_cassetteOSBinFilePath = texts[0]

	if _, err := os.Stat(_cassetteOSBinFilePath); os.IsNotExist(err) {
		_cassetteOSBinFilePath, err = exec.LookPath("cassetteos")

		if err != nil {
			return
		}
	}

	// locate the config file
	if len(texts) > 2 {
		for i, text := range texts {
			if text == "-c" {
				LegacyCassetteOSConfigFilePath = texts[i+1]
				break
			}
		}
	}

	if _, err := os.Stat(LegacyCassetteOSConfigFilePath); os.IsNotExist(err) {
		return
	}

	_configFile, _ = ini.Load(LegacyCassetteOSConfigFilePath)
}

func DetectLegacyVersion() (int, int, int, error) {
	if _, _, _, err := DetectVersion(); err == nil {
		return -1, -1, -1, ErrLegacyVersionNotFound
	}

	if _configFile == nil {
		return -1, -1, -1, ErrLegacyVersionNotFound
	}

	minorVersion, err := DetectMinorVersion()
	if err != nil {
		return -1, -1, -1, err
	}

	if minorVersion == 2 {
		return 0, 2, 99, nil // 99 means we don't know the patch version.
	}

	configKeyDBPathExist, err := IsConfigKeyDBPathExist()
	if err != nil {
		return -1, -1, -1, err
	}

	if !configKeyDBPathExist {
		return 0, 3, 0, nil // it could be 0.3.0, 0.3.1 or 0.3.2 but only one version can be returned.
	}

	return 0, 3, 3, nil // it could be 0.3.3 or 0.3.4 but only one version can be returned.
}

func DetectVersion() (int, int, int, error) {
	cmd := exec.Command(_cassetteOSBinFilePath, "-v")
	versionBytes, err := cmd.Output()
	if err != nil {
		return -1, -1, -1, ErrVersionNotFound
	}

	major, minor, patch, _, _, err := ParseVersion(string(versionBytes))
	if err != nil {
		return -1, -1, -1, ErrVersionNotFound
	}

	return major, minor, patch, nil
}

// Detect minor version of CassetteOS. It returns 2 for "0.2.x" or 3 for "0.3.x"
//
// (This is often useful when failing to get version from API because CassetteOS is not running.)
func DetectMinorVersion() (int, error) {
	if _configFile == nil {
		return -1, ErrLegacyVersionNotFound
	}

	if _configFile.Section("server").HasKey(configKeyUniqueToZero3x) {
		return 3, nil
	}

	return 2, nil
}

// Check if user data is stored in database (0.3.3+)
func IsConfigKeyDBPathExist() (bool, error) {
	if _configFile == nil {
		return false, ErrLegacyVersionNotFound
	}

	if !_configFile.Section("app").HasKey(configKeyDBPath) {
		return false, nil
	}

	return true, nil
}
