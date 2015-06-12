package main

import (
	"errors"
	"fmt"
	"github.com/docker/docker/pkg/progressreader"
	"github.com/docker/docker/pkg/streamformatter"
	"io"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
	"log"
	"net/http"
	"os"
	osUser "os/user"
	"path"
	"runtime"
	"strings"
)

var client *http.Client

func init() {
	user, err := osUser.Current()
	defaultOutput := ""

	if err == nil {
		if runtime.GOOS == "darwin" {
			defaultOutput = path.Join(user.HomeDir, "Documents", "Apple Events")
		} else {
			defaultOutput = path.Join(user.HomeDir, "Apple Events")
		}
	} else if cwd, err := os.Getwd(); err == nil {
		defaultOutput = cwd
	}

	output = kingpin.Flag("output", "Location to store output").Short('o').Default(defaultOutput).String()
}

var (
	output *string
	preferHd = kingpin.Flag("hd", "Prefer videos in HD quality").Bool()
	dryRun = kingpin.Flag("dry", "Dry run (don't download anything)").Short('n').Bool()
	noFolders = kingpin.Flag("no-folders", "Don't create a separate folder for each event").Bool()
)

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

			log.Println("Server requested authentication. Starting a browser...")
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

	if *dryRun {
		return errors.New(fmt.Sprintf("[DRY RUN]"))
	}

	reader := progressreader.New(progressreader.Config {
		In:        resp.Body,
		Out:       os.Stdout,
		Formatter: streamformatter.NewStreamFormatter(),
		Size:      int(resp.ContentLength),
		NewLines:  true,
		ID:        "",
		Action:    "Downloading",
	})

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
	if !*dryRun && len(path) > 0 { // could be empty
		if err := os.MkdirAll(path, 0755); err != nil {
			log.Printf("Failed to create output directory: %s\nReason: %s\n", path, err)
			os.Exit(1) // no use to try other files
		}
	}
}

func DownloadFile(session *Session, source string, fileName string) error {
	destinationDirectory := *output

	if !*noFolders {
		destinationDirectory = path.Join(destinationDirectory, session.Event.Name)
	}

	assertDirectory(destinationDirectory)
	destination := path.Join(destinationDirectory, SafeFileName(fileName))
	log.Printf("\n%s\n", destination)

	if FileExists(destination) {
		log.Printf("Already downloaded. Skipping...\n")
		return nil
	}

	temporary := destination + ".wddownload"

	if err := download(source, temporary); err != nil {
		log.Printf("Failed to download: %s\n", err)
		os.Remove(temporary)
		return err
	}

	if err := os.Rename(temporary, destination); err != nil {
		log.Printf("Failed to move temporary file to final location: %s\n", err)
		os.Remove(temporary)
		return err
	}

	return nil
}

func DownloadVideo(session *Session) {
	if url, fileName, ok := session.Video(*preferHd); ok {
		DownloadFile(session, url, fileName)
	}
}

func DownloadSlides(session *Session) {
	if url, fileName, ok := session.Slides(); ok {
		DownloadFile(session, url, fileName)
	}
}
