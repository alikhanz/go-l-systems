package sierpinski_triangle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSierpinskiTriangle_ApplyRules(t *testing.T) {
	alg := NewSierpinskiTriangle()
	out, err := alg.ApplyRules("F-G-G")
	assert.Equal(t, "F-G+F+G-F-GG-GG", out)
	assert.NoError(t, err)

	out, err = alg.ApplyRules("A")
	assert.Equal(t, "", out)
	assert.Error(t, err, "unknown character A")
}

func TestSierpinskiTriangle_ApplyRulesRecursively(t *testing.T) {
	alg := NewSierpinskiTriangle()

	out, err := alg.ApplyRulesRecursively(alg.axiom, 0)
	assert.Equal(t, alg.axiom, out)
	assert.NoError(t, err)

	out, err = alg.ApplyRulesRecursively(alg.axiom, 1)
	assert.Equal(t, "F-G+F+G-F-GG-GG", out)
	assert.NoError(t, err)

	out, err = alg.ApplyRulesRecursively(alg.axiom, 2)
	assert.Equal(t, "F-G+F+G-F-GG+F-G+F+G-F+GG-F-G+F+G-F-GGGG-GGGG", out)
	assert.NoError(t, err)

	out, err = alg.ApplyRulesRecursively("A", 3)
	assert.Equal(t, "", out)
	assert.Error(t, err, "Failed rules apply")
}