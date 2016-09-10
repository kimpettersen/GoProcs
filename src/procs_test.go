package main

import (
	"strconv"
	"testing"
)

func TestListProcessesAllPIDsAreNumbers(t *testing.T) {
	procs := listProcesses()
	for _, proc := range procs {
		if _, err := strconv.Atoi(proc.Pid); err != nil {
			t.Error("Found dir that is not a process")
		}
	}
}

func TestListProcessesProcessesHasNames(t *testing.T) {
	procs := listProcesses()
	for _, proc := range procs {
		if len(proc.Executable) < 1 {
			t.Error("Process does not have an executable" + proc.Pid)
		}
	}
}

func TestListProcessesMoreThanOneProcess(t *testing.T) {
	procs := listProcesses()
	if len(procs) < 2 {
		t.Errorf("should have more than one process running: %v", procs)
	}
}
