package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func main() {
	procs := listProcesses()
	reader := bufio.NewReader(os.Stdin)

	for idx, proc := range procs {
		fmt.Printf("(%v) %v - %v\n", idx+1, proc.Executable, proc.Pid)
	}
	fmt.Println("\nEnter number of process to kill")

	input, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal("Error reading input", err)
	}

	input = strings.TrimSpace(input)
	idx, err := strconv.Atoi(input)

	if err != nil {
		log.Fatalf("%v is not an option", input)
	}
	idx = idx - 1
	proc := procs[idx]
	kill(proc, "9")
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

func kill(proc Process, signal string) {
	if err := exec.Command("kill", "-"+signal, string(proc.Pid)).Start(); err != nil {
		log.Fatal("Couldn't kill ", err)
	}
	fmt.Printf("Killed %v\n", proc.Executable)

}
