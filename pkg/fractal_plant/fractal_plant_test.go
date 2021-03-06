package fractal_plant

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFractalPlant_ApplyRules(t *testing.T) {
	alg := NewFractalPlant()
	out, err := alg.ApplyRules(alg.axiom)
	assert.Equal(t, "F+[[X]-X]-F[-FX]+X", out)
	assert.NoError(t, err)

	out, err = alg.ApplyRules("A")
	assert.Equal(t, "", out)
	assert.Error(t, err, "unknown character A")
}

func TestFractalPlant_ApplyRulesRecursively(t *testing.T) {
	alg := NewFractalPlant()

	out, err := alg.ApplyRulesRecursively(alg.axiom, 0)
	assert.Equal(t, alg.axiom, out)
	assert.NoError(t, err)

	out, err = alg.ApplyRulesRecursively(alg.axiom, 1)
	assert.Equal(t, "F+[[X]-X]-F[-FX]+X", out)
	assert.NoError(t, err)

	out, err = alg.ApplyRulesRecursively(alg.axiom, 2)
	assert.Equal(t, "FF+[[F+[[X]-X]-F[-FX]+X]-F+[[X]-X]-F[-FX]+X]-FF[-FFF+[[X]-X]-F[-FX]+X]+F+[[X]-X]-F[-FX]+X", out)
	assert.NoError(t, err)

	out, err = alg.ApplyRulesRecursively("A", 3)
	assert.Equal(t, "", out)
	assert.Error(t, err, "Failed rules apply")
}