package worker

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"xyz.xeonds/nano-oj/database/model"
)

func (s *task) cJudger() (model.Status, string) {
	compileCommand := "g++ -o " + s.Workdir + "/program " + s.SourceFile
	compileResult, _ := exec.Command("bash", "-c", compileCommand).Output()
	if len(compileResult) != 0 {
		// TODO: return detailed info
		return model.CompilationError, "Compilation failed"
	}
	var result string
	var passed = true
	for i := 0; i < len(s.InputFiles); i++ {
		expectOutput, _ := os.ReadFile(s.ExpectFiles[i])
		// Run the program with the given input
		runCommand := s.Workdir + "/program < " + s.InputFiles[i]
		programOutput, _ := exec.Command("bash", "-c", runCommand).Output()

		// Compare the program output with the expected output
		programOutputStr := strings.ReplaceAll(string(programOutput), "\r", "")
		expectOutputStr := strings.ReplaceAll(string(expectOutput), "\r", "")
		programOutputStr = strings.ReplaceAll(programOutputStr, "\n", "")
		expectOutputStr = strings.ReplaceAll(expectOutputStr, "\n", "")
		if programOutputStr == expectOutputStr {
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
