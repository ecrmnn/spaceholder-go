package main

import (
	"flag"
	"github.com/ecrmnn/spaceholder-go/providers"
	"log"
	"os"
	"net/http"
	"io"
	"math/rand"
	"time"
	"fmt"
	"sync"
	"strconv"
)

func main() {
	// Get flags and parse
	provider := flag.String("provider", "random", "Set the image provider")
	number := flag.Int("number", 1, "Number of images to download")
	size := flag.String("size", "1024x768", "Set the image size")
	flag.Parse()

	var wg sync.WaitGroup
	wg.Add(*number)

	downloaded := 0

	for i := 1; i <= *number; i++ {
		go func() {
			defer wg.Done()

			time.Sleep(2 * time.Second)

			// Get image URL
			url := providers.GetImageFromProvider(*provider, *size)

			// Download image
			response, e := http.Get(url)
			if e != nil {
				log.Fatal(e)
			}

			defer response.Body.Close()

			// Open file for writing
			file, err := os.Create(currentDirectory() + "/" + generateFilename(*size))
			if err != nil {
				log.Fatal(err)
			}

			_, err = io.Copy(file, response.Body)
			if err != nil {
				log.Fatal(err)
			}
			file.Close()

			downloaded++
			fmt.Println("Downloaded " + strconv.Itoa(downloaded) + " of " + strconv.Itoa(*number))
		}()
	}

	fmt.Println("Downloading image(s) ...")
	wg.Wait()
}

func currentDirectory() string {
	directory, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

	return directory;
}

func generateFilename(size string) string {
	return "spaceholder_" + size + "_" + randomString() + ".jpg"
}

func randomString() string {
	rand.Seed(time.Now().UTC().UnixNano())
	const chars = "abcdefghijklmnopqrstuvwxyz0123456789"
	result := make([]byte, 6)
	for i := 0; i < 6; i++ {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result)
}
