package fractal_plant

import (
	"github.com/holizz/terrapin"
	"github.com/pkg/errors"
	"image"
	"math"
	"strings"
)

const Step = 2
const AngleOfRotation = 25 * math.Pi / 180

type stackElement struct {
	orientation float64
	position terrapin.Position
}

type FractalPlant struct {
	axiom string
	rules map[rune]string
	stack []stackElement
}

func NewFractalPlant() *FractalPlant {
	return &FractalPlant{
		axiom: "X",
		rules: map[rune]string{
			'X': "F+[[X]-X]-F[-FX]+X",
			'F': "FF",
		},
		stack: make([]stackElement, 0),
	}
}

func (p *FractalPlant) Render(gen int) (*terrapin.Terrapin, error) {
	out, err := p.ApplyRulesRecursively(p.axiom, gen)

	if err != nil {
		return nil, errors.Wrap(err, "Failed render caused by rules apply error")
	}

	i := image.NewRGBA(image.Rect(0, 0, 1200, 800))
	t := terrapin.NewTerrapin(i, terrapin.Position{X: 350.0, Y: 800})
	t.Right(AngleOfRotation)

	for _, char := range out {
		switch char {
		case 'F':
			// @todo: сделать отображение листа
			t.Forward(Step)
			break
		case '-':
			t.Right(AngleOfRotation)
			break
		case '+':
			t.Left(AngleOfRotation)
			break
		case '[':
			p.stack = append(p.stack, stackElement{
				orientation: t.Orientation,
				position:    t.Pos,
			})
			break
		case ']':
			el := p.stack[len(p.stack)-1]
			p.stack = p.stack[:len(p.stack)-1]
			t.Pos = el.position
			t.Orientation = el.orientation
		}
	}

	p.stack = make([]stackElement, 0)
	return t, nil
}

func (p *FractalPlant) ApplyRulesRecursively(input string, iterationsLeft int) (string, error) {
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

func (p *FractalPlant) ApplyRules(input string) (string, error) {
	var output strings.Builder

	for _, char := range input {
		if char == '[' || char == ']' || char == '+' || char == '-' {
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