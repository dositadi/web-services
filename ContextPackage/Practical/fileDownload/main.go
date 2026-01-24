package main

import (
	"fmt"
	"time"
)

type Image struct {
	Url   string
	Image string
}

// Demonstrating Pipeline!

func Download(images []Image) <-chan string {
	out := make(chan string)

	go func() {
		for i, image := range images {
			fmt.Printf("Downloading %s from %s....\n", image.Image, image.Url)
			time.Sleep(time.Duration(10*i) * 500 * time.Millisecond)
			out <- image.Image
			fmt.Printf("%s Download has been completed.\n", image.Image)
		}
		close(out)
	}()
	return out
}

func Resize(in <-chan string) <-chan string {
	out := make(chan string)

	go func() {
		for image := range in {
			fmt.Printf("Resizing image: %s....\n", image)
			time.Sleep(time.Duration(2) * 500 * time.Millisecond)
			out <- image
			fmt.Printf("Image Resize completed for: %s\n", image)
		}
		close(out)
	}()
	return out
}

func Upload(in <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		for image := range in {
			fmt.Printf("Uploading image: %s....\n", image)
			time.Sleep(time.Duration(4) * 500 * time.Millisecond)
			out <- image
			fmt.Printf("Image Upload successful for: %s\n", image)
		}
		close(out)
	}()
	return out
}

func main() {
	images := []Image{
		Image{
			Url:   "google.com",
			Image: "google.jpg",
		},
		Image{
			Url:   "facebook.com",
			Image: "facebook.jpg",
		},
		Image{
			Url:   "amazon.com",
			Image: "amazon.jpg",
		},
		Image{
			Url:   "x.com",
			Image: "x.png",
		},
		Image{
			Url:   "instagram.com",
			Image: "instagram.jpg",
		},
	}

	resizeImageChan := Download(images)

	uploadImageChan := Resize(resizeImageChan)

	uploadedChan := Upload(uploadImageChan)

	for image := range uploadedChan {
		fmt.Printf("Uploaded '%s' to database successfully!\n\n", image)
	}
}
