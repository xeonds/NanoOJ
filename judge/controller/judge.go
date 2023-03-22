package controller

import (
	"net/http"
	"os"
	"os/exec"
	"strconv"

	"github.com/gin-gonic/gin"
	"xyz.xeonds/nano-oj/model"
)

var problemSet map[int]model.Problem

func init() {
	problemSet = map[int]model.Problem{
		1: {
			ProblemID:      1,
			ProblemInput:   "input1",
			ExpectedOutput: "output1",
		},
		2: {
			ProblemID:      2,
			ProblemInput:   "input2",
			ExpectedOutput: "output2",
		},
	}
}

func JudgeProblemByID(c *gin.Context) {
	problemID := c.PostForm("problemID")
	sourceCode := c.PostForm("sourceCode")
	id, err := strconv.Atoi(problemID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "problemID must be an integer"})
		return
	}

	expectedOutput := problemSet[id].ExpectedOutput
	sourceCodeFile, err := os.Create("program.cc")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer sourceCodeFile.Close()
	_, err = sourceCodeFile.WriteString(sourceCode)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	expectedOutputFile, err := os.Create("program.out")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer expectedOutputFile.Close()
	_, err = expectedOutputFile.WriteString(expectedOutput)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	cmd := exec.Command("judge", "-s", "program.cc", "-e", "program.out")
	result, err := cmd.Output()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": string(result)})
}
