package worker

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"xyz.xeonds/nano-oj/model"
)

func CPP(t *Task) (model.Status, string) {
	compileCommand := "g++ -o " + t.Workdir + "/program " + t.SourceFile
	compileResult, _ := exec.Command("bash", "-c", compileCommand).Output()
	if len(compileResult) != 0 {
		// TODO: return detailed info
		return model.CompilationError, "Compilation failed"
	}
	var result string
	var passed = true
	for i := 0; i < len(t.InputFiles); i++ {
		// Run the program with the given input
		runCommand := fmt.Sprintf("timeout %d %s/program < %s | diff - %s", t.TimeLimit, t.Workdir, t.InputFiles[i], t.ExpectFiles[i])
		programOutput, _ := exec.Command("bash", "-c", runCommand).Output()

		// Compare the program output with the expected output
		if string(programOutput) == "" {
			result += fmt.Sprintf("Test case %d passed\n", i+1)
		} else {
			result += fmt.Sprintf("Test case %d failed\n", i+1)
			passed = false
		}
	}
	if passed {
		return model.Accepted, result
	} else {
		return model.WrongAnswer, result
	}
}

func Python(t *Task) (model.Status, string) {
	var result string
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
		}
	}
	return status, result
}
