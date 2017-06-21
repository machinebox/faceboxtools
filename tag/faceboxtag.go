package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"flag"

	"github.com/machinebox/sdk-go/facebox"
)

func main() {
	var (
		dir         = flag.String("dir", "./testdata", "source directory")
		faceboxAddr = flag.String("facebox", "http://localhost:8080", "facebox address")
		images      = flag.String("images", ".jpg", "image files extension")
	)
	flag.Parse()

	faceboxClient := facebox.New(*faceboxAddr)

	filepath.Walk(*dir, func(path string, info os.FileInfo, err error) error {
		if !strings.HasSuffix(info.Name(), *images) {
			return nil
		}
		parts := strings.Split(path, string(filepath.Separator))
		if len(parts) < 2 {
			return nil
		}
		log.Printf("+ %v\n", info.Name())
		tags, err := tag(faceboxClient, path)
		if err != nil {
			log.Println("[ERROR]: Tagging", err)
			return nil
		}
		for _, tag := range tags {
			log.Printf("\t%v\n", tag)
		}
		return nil
	})
}

func tag(faceboxClient *facebox.Client, path string) ([]string, error) {
	names := []string{}
	r, err := os.Open(path)
	if err != nil {
		return names, err
	}
	defer r.Close()
	faces, err := faceboxClient.Check(r)
	if err != nil {
		return names, err
	}
	for _, face := range faces {
		if face.Matched {
			names = append(names, face.Name)
		}
	}
	return names, nil
}
