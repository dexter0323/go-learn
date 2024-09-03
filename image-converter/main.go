package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/chai2010/webp"
)

func main() {
    dir := "./path/to/your/images" // Replace with your directory path

    // Create a channel to send file paths to workers
    fileChan := make(chan string)
    var wg sync.WaitGroup

    // Start a fixed number of worker goroutines
    numWorkers := 8
    for i := 0; i < numWorkers; i++ {
        wg.Add(1)
        go worker(fileChan, &wg)
    }

    // Walk through the directory and send file paths to the channel
    err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }

        if !info.IsDir() {
            ext := strings.ToLower(filepath.Ext(info.Name()))
            if ext == ".jpg" || ext == ".jpeg" || ext == ".png" {
                fileChan <- path
            }
        }
        return nil
    })

    if err != nil {
        fmt.Printf("Error walking the path %q: %v\n", dir, err)
    }

    // Close the channel and wait for all workers to finish
    close(fileChan)
    wg.Wait()
}

func worker(fileChan <-chan string, wg *sync.WaitGroup) {
    defer wg.Done()
    for path := range fileChan {
        err := convertToWebP(path)
        if err != nil {
            fmt.Printf("Failed to convert %s: %v\n", path, err)
        } else {
            fmt.Printf("Converted %s to WebP\n", path)
        }
    }
}

func convertToWebP(path string) error {
    file, err := os.Open(path)
    if err != nil {
        return err
    }
    defer file.Close()

    var img image.Image
    ext := strings.ToLower(filepath.Ext(path))
    switch ext {
    case ".jpg", ".jpeg":
        img, err = jpeg.Decode(file)
    case ".png":
        img, err = png.Decode(file)
    default:
        return fmt.Errorf("unsupported file type: %s", ext)
    }
    if err != nil {
        return err
    }

    // Create the 'converted' directory if it doesn't exist
    dir := filepath.Dir(path)
    convertedDir := filepath.Join(dir, "converted")
    if _, err := os.Stat(convertedDir); os.IsNotExist(err) {
        err = os.Mkdir(convertedDir, 0755)
        if err != nil {
            return err
        }
    }

    // Construct the output path in the 'converted' directory
    filename := strings.TrimSuffix(filepath.Base(path), ext) + ".webp"
    outPath := filepath.Join(convertedDir, filename)
    outFile, err := os.Create(outPath)
    if err != nil {
        return err
    }
    defer outFile.Close()

    options := &webp.Options{
        Lossless: false,
        // Quality:  75,
    }
    err = webp.Encode(outFile, img, options)
    if err != nil {
        return err
    }

    return nil
}