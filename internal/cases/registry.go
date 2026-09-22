package cases

import (
	"errors"
	"fmt"
)

var ErrCaseNotFound = errors.New("case not found")

type CaseRegistry struct {
	cases map[string]string
}

func NewCaseRegistry() *CaseRegistry {
	return &CaseRegistry{
		cases: map[string]CaseEntry{
			"case_017": {
				ID:    "case_017",
				Title: "Последний экземпляр",
				Path:  "cases/case_017",
			},
		},
	}
}

type CaseEntry struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Path  string `json:"-"`
}

func NewCaseRegistry() *CaseRegistry {
	return &CaseRegistry{
		cases: map[string]string{
			"case_017": "cases/case_017",
		},
	}
}

func (r *CaseRegistry) GetPath(caseID string) (string, error) {
	path, ok := r.cases[caseID]
	if !ok {
		return "", fmt.Errorf("%w: %s", ErrCaseNotFound, caseID)
	}

	return path, nil
}
