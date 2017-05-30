package ice

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/mholt/archiver"
)

const releaseAPI = "https://api.github.com/repos/electron/electron/releases/latest"

// ElectronExists checks if an electron prebuilt binary exists
func ElectronExists() bool {
	_, err := exec.LookPath(filepath.Join("electron", "electron"))
	return err == nil
}

// GetElectron gets electron pre-built binaries for packaging and distribution
func GetElectron() error {
	if ElectronExists() {
		if Verbose {
			log.Print("electron already exists in directory, skipping fetch")
		}
		return nil
	}
	url, err := releaseURL()
	if err != nil {
		return fmt.Errorf("unable to get latest Electron release: %v", err)
	}
	fmt.Printf("Now fetching latest stable Electron prebuilt: %s\n", url)
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer func() { _ = resp.Body.Close() }()

	zipFile, err := ioutil.TempFile("", "")
	if err != nil {
		return fmt.Errorf("unable to create temp file for use: %v", err)
	}
	defer func() { _ = os.Remove(zipFile.Name()) }()

	_, err = io.Copy(zipFile, resp.Body)
	if err != nil {
		return fmt.Errorf("unable to copy response body to file: %v", err)
	}
	_ = resp.Body.Close()
	if err := zipFile.Close(); err != nil {
		return err
	}
	basedir := "electron"
	if err := os.MkdirAll(filepath.Join("electron", "resources", "app"), 0755); err != nil {
		return err
	}
	return archiver.Zip.Open(zipFile.Name(), basedir)
}

type apiResp struct {
	Name       string
	PreRelease bool

	Assets []struct {
		ID          int
		Name        string
		ContentType string
		Created     string `json:"created_at"`
		Updated     string `json:"updated_at"`
		DownloadURL string `json:"browser_download_url"`
	}
}

func releaseURL() (string, error) {
	cl := &http.Client{
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest("GET", releaseAPI, nil)
	if err != nil {
		return "", err
	}
	resp, err := cl.Do(req)
	if err != nil {
		return "", err
	}
	defer func() { _ = resp.Body.Close() }()
	dat, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	_ = resp.Body.Close()
	var ar apiResp
	if err := json.Unmarshal(dat, &ar); err != nil {
		return "", err
	}
	if ar.Name == "" {
		return "", errors.New("invalid response from Github API")
	}

	prf := runtime.GOOS
	suf := runtime.GOARCH
	if runtime.GOOS == "windows" {
		prf = "win32"
	}
	if runtime.GOARCH == "amd64" {
		suf = "x64"
	}
	want := fmt.Sprintf("%s-%s.zip", prf, suf)
	for _, asset := range ar.Assets {
		if strings.HasPrefix(asset.Name, "electron") && strings.HasSuffix(asset.Name, want) {
			return asset.DownloadURL, nil
		}
	}
	return "", fmt.Errorf("unable to find Electron binary for %s", want)
}
