package worker

import "fmt"

type SourceCodeJudgement struct {
	Workdir     string
	Lang        string
	SourceFile  string
	InputFiles  []string
	ExpectFiles []string
	TimeLimit   int
}

func (s *SourceCodeJudgement) Judge() (string, error) {
	switch s.Lang {
	case "c", "cpp":
		return s.cJudger(), nil
	// case "python":
	// return judgement.pyJudger(), nil
	default:
		return "", fmt.Errorf("unsupported language: %s", s.Lang)
	}
}
