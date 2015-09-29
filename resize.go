package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"path/filepath"
	"regexp"

	"github.com/disintegration/imaging"
)

var width = flag.Int("width", math.MaxInt16, "Output image width")
var height = flag.Int("height", math.MaxInt16, "Output image height")
var output = flag.String("type", "png", "Output image format")
var destination = flag.String("destination", "", "Output image folder")

func main() {
	flag.Parse()

	if len(*destination) > 0 {
		err := os.MkdirAll(*destination, os.ModePerm)
		if err != nil {
			fmt.Printf("Unable to create destination: %s\n", err)
			os.Exit(1)
		}
	}

	extMatch := regexp.MustCompile(`\.[a-zA-Z0-9]+$`)

	filelist := flag.Args()
	// if len(filelist) == 0 {
	// 	filelist = []string{"./*.*"}
	// }

	for _, srcArg := range filelist {
		srcFilenames, err := filepath.Glob(srcArg)
		if err != nil {
			log.Printf("Glob error: %s\n", err)
			continue
		}
		for _, srcFilename := range srcFilenames {
			fmt.Printf("Processing [%s]... ", srcFilename)
			src, err := imaging.Open(srcFilename)
			if err != nil {
				log.Printf("error: %s\n", err)
				continue
			}
			dest := imaging.Fit(src, *width, *height, imaging.MitchellNetravali)
			destFilename := extMatch.ReplaceAllString(srcFilename, "") + fmt.Sprintf("-%dx%d.%s", dest.Bounds().Dx(), dest.Bounds().Dy(), *output)
			if len(*destination) > 0 {
				_, file := filepath.Split(destFilename)
				destFilename = filepath.Join(*destination, file)
			}
			err = imaging.Save(dest, destFilename)
			if err != nil {
				log.Printf("error: %s\n", err)
				continue
			}
			fmt.Printf("OK; written to %s\n", destFilename)
		}
	}
}
