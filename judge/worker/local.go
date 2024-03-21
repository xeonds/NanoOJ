package worker

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"xyz.xeonds/nano-oj/model"
)

func CPP(t *Task) (model.Status, string, int) {
	compileCommand := "g++ -o " + t.Workdir + "/program " + t.SourceFile
	compileResult, _ := exec.Command("bash", "-c", compileCommand).Output()
	if len(compileResult) != 0 {
		return model.CompilationError, string(compileResult), 0
	}
	var result string
	var passed = true
	rank := 100
	for i := 0; i < len(t.InputFiles); i++ {
		// Run the program with the given input
		runCommand := fmt.Sprintf("timeout %d %s/program < %s | diff - %s", t.TimeLimit, t.Workdir, t.InputFiles[i], t.ExpectFiles[i])
		programOutput, _ := exec.Command("bash", "-c", runCommand).Output()

		// Compare the program output with the expected output
		if string(programOutput) == "" {
			result += fmt.Sprintf("Test case %d passed\n", i+1)
		} else {
			result += fmt.Sprintf("Test case %d failed\n", i+1)
			rank -= 100 / len(t.InputFiles)
			passed = false
		}
	}
	if passed {
		return model.Accepted, result, rank
	} else {
		return model.WrongAnswer, result, rank
	}
}

func Python(t *Task) (model.Status, string, int) {
	var result string
	rank := 100
	status := model.Accepted
	for i := 0; i < len(t.InputFiles); i++ {
		expectOutput, _ := os.ReadFile(t.ExpectFiles[i])
		// Run the program with the given input
		runCommand := "python3 " + t.Workdir + "/" + t.SourceFile + " < " + t.InputFiles[i]
		programOutput, _ := exec.Command("bash", "-c", runCommand).Output()

		// Compare the program output with the expected output
		programOutputStr := strings.ReplaceAll(string(programOutput), "\r", "")
		expectOutputStr := strings.ReplaceAll(string(expectOutput), "\r", "")
		programOutputStr = strings.ReplaceAll(programOutputStr, "\n", "")
		expectOutputStr = strings.ReplaceAll(expectOutputStr, "\n", "")
		if programOutputStr == expectOutputStr {
			result += fmt.Sprintf("Test case %d passed\n", i+1)
		} else {
			status = model.WrongAnswer
			result += fmt.Sprintf("Test case %d failed\n", i+1)
			rank -= 100 / len(t.InputFiles)
		}
	}
	return status, result, rank
}
