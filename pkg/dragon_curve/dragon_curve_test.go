package dragon_curve

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDragonCurve_ApplyRules(t *testing.T) {
	alg := NewDragonCurve()
	out, err := alg.ApplyRules("F")
	assert.Equal(t, "F+G", out)
	assert.NoError(t, err)

	out, err = alg.ApplyRules("A")
	assert.Equal(t, "", out)
	assert.Error(t, err, "unknown character A")
}

func TestDragonCurve_ApplyRulesRecursively(t *testing.T) {
	alg := NewDragonCurve()

	out, err := alg.ApplyRulesRecursively(alg.axiom, 0)
	assert.Equal(t, alg.axiom, out)
	assert.NoError(t, err)

	out, err = alg.ApplyRulesRecursively(alg.axiom, 1)
	assert.Equal(t, "F+G", out)
	assert.NoError(t, err)

	out, err = alg.ApplyRulesRecursively(alg.axiom, 2)
	assert.Equal(t, "F+G+F-G", out)
	assert.NoError(t, err)

	out, err = alg.ApplyRulesRecursively("A", 3)
	assert.Equal(t, "", out)
	assert.Error(t, err, "Failed rules apply")
}