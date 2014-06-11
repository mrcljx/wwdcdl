package main

import (
	"errors"
	"flag"
	"fmt"
	docker "github.com/dotcloud/docker/utils"
	"io"
	"net/http"
	"os"
	osUser "os/user"
	"path"
	"runtime"
	"strings"
)

var preferHd bool
var dryRun bool
var output string
var client *http.Client

func init() {
	user, err := osUser.Current()
	defaultOutput := ""

	if err == nil {
		if runtime.GOOS == "darwin" {
			defaultOutput = path.Join(user.HomeDir, "Documents", "WWDC")
		} else {
			defaultOutput = path.Join(user.HomeDir, "WWDC")
		}
	} else if cwd, err := os.Getwd(); err == nil {
		defaultOutput = cwd
	}

	flag.StringVar(&output, "output", defaultOutput, "Location to store output")
	flag.BoolVar(&preferHd, "hd", false, "Prefer videos in HD quality")
	flag.BoolVar(&dryRun, "n", false, "Dry run (don't download anything)")
}

func download(source string, destination string) (err error) {
	req, err := http.NewRequest("GET", source, nil)

	if err != nil {
		return
	}

	if err != nil {
		return
	}

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		// fmt.Printf("Body:\n")
		// io.Copy(os.Stderr, resp.Body)
		return errors.New(fmt.Sprintf("Server responded with unexpected status-code %d", resp.StatusCode))
	}

	if resp.Request.URL != nil {
		redirectedToLogin := strings.Index(resp.Request.URL.Host, "daw.apple.com") >= 0

		if redirectedToLogin {
			if authenticator.IsAuthenticated() {
				return errors.New("Server requested authentication but we authenticated already.")
			}

			fmt.Printf("Server requested authentication. Starting a browser...\n")
			err = authenticator.Authenticate()

			if err != nil {
				return
			}

			// retry download
			return download(source, destination)
		}
	}

	if contentType := resp.Header.Get("Content-Type"); strings.Index(contentType, "text/") == 0 {
		return errors.New(fmt.Sprintf("Server responded with unexpected content-type '%s'", contentType))
	}

	if dryRun {
		return errors.New(fmt.Sprintf("[DRY RUN]"))
	}

	formatter := docker.NewStreamFormatter(false)
	reader := docker.ProgressReader(resp.Body, int(resp.ContentLength), os.Stdout, formatter, true, source, "Downloading")

	file, err := os.Create(destination)

	if err != nil {
		return
	}

	defer file.Close()

	_, err = io.Copy(file, reader)
	return
}

func FileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func assertDirectory(path string) {
	if !dryRun && len(path) > 0 { // could be empty
		if err := os.MkdirAll(path, 0755); err != nil {
			fmt.Printf("Failed to create output directory: %s\nReason: %s\n", path, err)
			os.Exit(1) // no use to try other files
		}
	}
}

func DownloadFile(source string, fileName string) error {

	assertDirectory(output)
	destination := path.Join(output, fileName)
	fmt.Printf("\n%s\n", destination)

	if FileExists(destination) {
		fmt.Printf("Already downloaded. Skipping...\n")
		return nil
	}

	temporary := destination + ".wddownload"

	if err := download(source, temporary); err != nil {
		fmt.Printf("Failed to download: %s\n", err)
		os.Remove(temporary)
		return err
	}

	if err := os.Rename(temporary, destination); err != nil {
		fmt.Printf("Failed to move temporary file to final location: %s\n", err)
		os.Remove(temporary)
		return err
	}

	return nil
}

func DownloadVideo(session *Session) {
	if url, fileName, ok := session.Video(preferHd); ok {
		DownloadFile(url, fileName)
	}
}

func DownloadSlides(session *Session) {
	if url, fileName, ok := session.Slides(); ok {
		DownloadFile(url, fileName)
	}
}
