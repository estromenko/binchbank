package parsers_test

import (
	"os"
	"testing"

	"github.com/estromenko/binchbank/internal/models"
	"github.com/estromenko/binchbank/parsers"
)

func TestParsers(t *testing.T) {
	file, err := os.Open("../data/test.csv")

	if err != nil {
		t.Error(err)
		return
	}

	var ops []*models.Operation

	parsers.Parse(file, &ops)
}
