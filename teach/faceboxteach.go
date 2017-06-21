package main

import (
	"context"
	"flag"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/machinebox/sdk-go/facebox"
	"github.com/machinebox/sdk-go/x/boxutil"
)

func main() {
	var (
		dir         = flag.String("dir", "./testdata", "source directory")
		faceboxAddr = flag.String("facebox", "http://localhost:8080", "facebox address")
		images      = flag.String("images", ".jpg", "image files extension")
	)
	flag.Parse()

	// make a new facebox client
	faceboxClient := facebox.New(*faceboxAddr)

	log.Println("waiting for box to be ready...")
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	err := boxutil.WaitForReady(ctx, faceboxClient)
	if err != nil {
		if err == boxutil.ErrCanceled {
			log.Fatalln("timed out waiting for box to be ready")
		}
		log.Fatalln(err)
	}
	log.Println("box ready")

	filepath.Walk(*dir, func(path string, info os.FileInfo, err error) error {
		if !strings.HasSuffix(info.Name(), *images) {
			return nil
		}
		parts := strings.Split(path, string(filepath.Separator))
		if len(parts) < 2 {
			return nil
		}
		name := parts[len(parts)-2]
		name = strings.Replace(name, "_", " ", -1)
		log.Printf("+ Teach: %v (%v)\n", name, info.Name())
		err = teachFromFile(faceboxClient, path, name, info.Name())
		if err != nil {
			log.Println("[ERROR]: Teaching", err)
			return nil
		}
		return nil
	})
}

func teachFromFile(faceboxClient *facebox.Client, path, name, filename string) error {
	r, err := os.Open(path)
	if err != nil {
		return err
	}
	defer r.Close()
	err = faceboxClient.Teach(r, filename, name)
	if err != nil {
		return err
	}
	return nil
}
