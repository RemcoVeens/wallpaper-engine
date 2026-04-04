package main

import (
	"bytes"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"strings"
)

func main() {
	uhdir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("could not find user home directory")
	}
	FOLDER := uhdir + "/Pictures/bobby"
	if len(os.Args) < 2 {
		log.Fatal("Usage: wallpaper-engine <argument>")
	}
	switch os.Args[1] {
	case "random":
		// get all .jpg files from FOLDER
		files, err := os.ReadDir(FOLDER)
		if err != nil {
			log.Fatal(err)
		}
		var jpgFiles []string
		for _, file := range files {
			if strings.HasSuffix(file.Name(), ".jpg") {
				jpgFiles = append(jpgFiles, fmt.Sprintf("%s/%s", FOLDER, file.Name()))
			}
		}
		if len(jpgFiles) == 0 {
			log.Fatal("No .jpg files found in FOLDER")
		}
		randomFile := jpgFiles[rand.Intn(len(jpgFiles))]
		log.Printf("Setting wallpaper to %s", randomFile)
		cmd := exec.Command("awww", "img", randomFile)
		var out bytes.Buffer
		cmd.Stdout = &out
		err = cmd.Run()
		if err != nil {
			log.Printf("Error running command: %v", err)
		}
	case "list":
		files, err := os.ReadDir(FOLDER)
		if err != nil {
			log.Fatal(err)
		}
		for _, file := range files {
			if strings.HasSuffix(file.Name(), ".jpg") {
				fmt.Println(file.Name())
			}
		}
	case "set":
		if len(os.Args) != 3 {
			log.Fatal("Usage: wallpaper-engine set <file>")
		}
		file := os.Args[2]
		parced_file := fmt.Sprintf("%s/%s", FOLDER, file)
		log.Printf("Setting wallpaper to %s", parced_file)
		cmd := exec.Command("awww", "img", parced_file)
		var out bytes.Buffer
		cmd.Stdout = &out
		err = cmd.Run()
		if err != nil {
			log.Printf("Error running command: %v", err)
		}
	default:
		log.Fatal("Invalid argument")
	}
}
