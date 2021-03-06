package koch_curve

import (
	"image"
	"math"
	"strings"

	"github.com/holizz/terrapin"
	"github.com/pkg/errors"
)

const Step = 5
const AngleOfRotation = 90 * math.Pi / 180

type KochCurve struct {
	axiom string
	rules map[rune]string
}

func NewKochCurve() *KochCurve {
	return &KochCurve{
		axiom: "F",
		rules: map[rune]string{
			'F': "F+F-F-F+F",
		},
	}
}

func (p *KochCurve) Render(gen int) (*terrapin.Terrapin, error) {
	out, err := p.ApplyRulesRecursively(p.axiom, gen)

	if err != nil {
		return nil, errors.Wrap(err, "Failed render caused by rules apply error")
	}

	i := image.NewRGBA(image.Rect(0, 0, 1300, 700))
	t := terrapin.NewTerrapin(i, terrapin.Position{X: 0, Y: 698})
	t.Right(AngleOfRotation)

	for _, char := range out {
		switch char {
		case 'F':
			t.Forward(Step)
			break
		case '-':
			t.Right(AngleOfRotation)
			break
		case '+':
			t.Left(AngleOfRotation)
		}
	}
	return t, nil
}

func (p *KochCurve) ApplyRulesRecursively(input string, iterationsLeft int) (string, error) {
	if iterationsLeft < 0 {
		return "", errors.New("Iterations count cannot be less than 0")
	}

	if iterationsLeft == 0 {
		return input, nil
	}

	result, err := p.ApplyRules(input)

	if err != nil {
		return "", errors.Wrap(err, "Failed rules apply")
	}

	iterationsLeft--
	return p.ApplyRulesRecursively(result, iterationsLeft)
}

func (p *KochCurve) ApplyRules(input string) (string, error) {
	var output strings.Builder

	for _, char := range input {
		if char == '+' || char == '-' {
			output.WriteRune(char)
			continue
		}

		converted, ok := p.rules[char]

		if !ok {
			return "", errors.New("unknown character " + string(char))
		}

		output.WriteString(converted)
	}

	return output.String(), nil
}