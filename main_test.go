package dbsimpl

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSanityTest(t *testing.T) {
	b := NewRaw()
	b.Select("*").
		From("table").
		Where("TRUE")

	t.Log(b.Build())
	assert.Equal(t, "SELECT * FROM table WHERE TRUE", b.Build())
}

func TestRawStringAddition(t *testing.T) {

	b := NewRaw()
	b.Select("*").
		From("table1").
		RawS("INNER JOIN table2 ON table1.id = table2.id")

	assert.Equal(t, "SELECT * FROM table1 INNER JOIN table2 ON table1.id = table2.id", b.Build())
}
