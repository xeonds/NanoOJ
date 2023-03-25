package worker

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func (s *task) pyJudger() string {
	var result string
	for i := 0; i < len(s.InputFiles); i++ {
		expectOutput, _ := os.ReadFile(s.ExpectFiles[i])
		// Run the program with the given input
		runCommand := "python3 " + s.Workdir + "/" + s.SourceFile + " < " + s.InputFiles[i]
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
		}
	}
	return result
}
