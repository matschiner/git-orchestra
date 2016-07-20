package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

type configstr struct {
	Path string `json:"path"`
}

var filename = ".git_orchestra.config"

func main() {
	// fmt.Println("test")
	var config []configstr
	var err error
	body, _ := ioutil.ReadFile(filename)
	json.Unmarshal(body, &config)
	if len(os.Args) < 2 {
		log.Fatal("Please input a command!")
	}
	action := os.Args[1]
	switch action {

	case "add":

		if len(os.Args) < 3 {
			log.Fatal("Please input a path!")
		}
		inputpath := os.Args[2]
		_, err3 := os.Stat(inputpath)
		if os.IsNotExist(err3) {
			log.Fatal(inputpath + " is not a valid path!")
		} else if err3 != nil {
			log.Fatal(err3)
		}

		config = append(config, configstr{os.Args[2]})

		writeconfig(config)

		fmt.Println("Added " + os.Args[2] + "!")

	case "rm":

		if len(os.Args) < 3 {
			log.Fatal("Please input a path!")
		}
		inputpath := os.Args[2]

		var confign []configstr
		for _, path := range config {
			if path.Path != inputpath {
				confign = append(confign, path)
			}
		}
		writeconfig(confign)

		fmt.Println("Deleted " + os.Args[2] + "!")

	case "pull":
		fmt.Println("Pulling ...")
		for _, path := range config {
			fmt.Println(" - " + path.Path)
			// cmd := exec.Command("bash", "-c", "cd ", path.Path, "&&", "ls")
			err = os.Chdir(path.Path)
			if err != nil {
				fmt.Println(err)
			}
			// cmd := exec.Command("ls", "&&", "git", "pull")
			cmd := exec.Command("ls")
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			cmd.Run()
			os.Chdir("../")
			// fmt.Println(out)
		}
	case "push":
		fmt.Println("Pushing ...")
		for _, path := range config {
			fmt.Println(" - " + path.Path)
			// cmd := exec.Command("bash", "-c", "cd ", path.Path, "&&", "ls")
			os.Chdir(path.Path)
			cmd := exec.Command("git", "push")
			//cmd := exec.Command("ls")
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			cmd.Run()
			os.Chdir("../")
			// fmt.Println(out)
		}
	case "commit":
		fmt.Println("Commiting ...")
		for _, path := range config {
			fmt.Println(" - " + path.Path)
			// cmd := exec.Command("bash", "-c", "cd ", path.Path, "&&", "ls")
			os.Chdir(path.Path)
			cmd := exec.Command("git", "commit", "-a")
			//cmd := exec.Command("ls")
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			cmd.Run()
			os.Chdir("../")
			// fmt.Println(out)
		}
	case "list":
		fmt.Println("Listing ...")
		for _, path := range config {
			fmt.Println(" - " + path.Path)
		}
	default:
		fmt.Println("Please input a valid command!")

	}
	// // config = append(config, configstr{"test"})
	// j, err := json.MarshalIndent(&config, " ", "   ")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	//
	// err2 := ioutil.WriteFile(".gitpullconfig", j, 0644)
	// if err2 != nil {
	// 	fmt.Println(err2)
	// }
}
func writeconfig(config []configstr) {
	j, err := json.MarshalIndent(&config, " ", "   ")
	if err != nil {
		fmt.Println(err)
	}
	err2 := ioutil.WriteFile(filename, j, 0644)
	if err2 != nil {
		fmt.Println(err2)
	}
}
