package worker

import (
	"fmt"
	"os/exec"

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
		// Run the program with the given input
		runCommand := fmt.Sprintf("timeout %d %s/program < %s | diff - %s", s.TimeLimit, s.Workdir, s.InputFiles[i], s.ExpectFiles[i])
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
