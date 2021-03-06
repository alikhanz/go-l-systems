package fractal_tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFractalTree_ApplyRules(t *testing.T) {
	alg := NewFractalTree()
	out, err := alg.ApplyRules("0")
	assert.Equal(t, "1[0]0", out)
	assert.NoError(t, err)

	out, err = alg.ApplyRules("A")
	assert.Equal(t, "", out)
	assert.Error(t, err, "unknown character A")
}

func TestFractalTree_ApplyRulesRecursively(t *testing.T) {
	alg := NewFractalTree()

	out, err := alg.ApplyRulesRecursively(alg.axiom, 0)
	assert.Equal(t, alg.axiom, out)
	assert.NoError(t, err)

	out, err = alg.ApplyRulesRecursively(alg.axiom, 1)
	assert.Equal(t, "1[0]0", out)
	assert.NoError(t, err)

	out, err = alg.ApplyRulesRecursively(alg.axiom, 2)
	assert.Equal(t, "11[1[0]0]1[0]0", out)
	assert.NoError(t, err)

	out, err = alg.ApplyRulesRecursively(alg.axiom, 3)
	assert.Equal(t, "1111[11[1[0]0]1[0]0]11[1[0]0]1[0]0", out)
	assert.NoError(t, err)

	out, err = alg.ApplyRulesRecursively("A", 3)
	assert.Equal(t, "", out)
	assert.Error(t, err, "Failed rules apply")
}