package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"regexp"

	"github.com/disintegration/imaging"
)

var width = flag.Int("width", math.MaxInt16, "Output image width")
var height = flag.Int("height", math.MaxInt16, "Output image height")
var output = flag.String("type", "png", "Output image format")

func main() {
	flag.Parse()

	extMatch := regexp.MustCompile(`\.[a-zA-Z0-9]+$`)

	for _, srcFilename := range flag.Args() {
		fmt.Printf("Processing [%s]... ", srcFilename)
		src, err := imaging.Open(srcFilename)
		if err != nil {
			log.Printf("error: %s\n", err)
			continue
		}
		dest := imaging.Fit(src, *width, *height, imaging.MitchellNetravali)
		destFilename := extMatch.ReplaceAllString(srcFilename, "") + fmt.Sprintf("-%dx%d.%s", dest.Bounds().Dx(), dest.Bounds().Dy(), *output)
		err = imaging.Save(dest, destFilename)
		if err != nil {
			log.Printf("error: %s\n", err)
			continue
		}
		fmt.Printf("OK; written to %s\n", destFilename)
	}
}
