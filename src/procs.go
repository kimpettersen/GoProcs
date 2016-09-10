package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(listProcesses())
}

type Process struct {
	Pid        string
	Executable string
}

func listProcesses() []Process {
	var processes []Process
	files, err := ioutil.ReadDir("/proc")

	if err != nil {
		log.Fatal("Could not read dir /proc :)")
	}
	var proc Process

	for _, file := range files {
		if _, err := strconv.Atoi(file.Name()); err == nil {

			cmd, err := ioutil.ReadFile("/proc/" + file.Name() + "/cmdline")

			cmdString := strings.Join(strings.Split(string(cmd), "\x00"), " ")

			if err != nil {
				log.Fatal("Can't read file:", err)
			}

			proc = Process{
				Pid:        file.Name(),
				Executable: cmdString,
			}
			processes = append(processes, proc)
		}
	}
	return processes
}
