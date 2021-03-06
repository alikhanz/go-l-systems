package main

import (
	"fmt"
	"image/png"
	"log"
	"net/http"
	"strconv"

	"github.com/alikhanz/go-l-systems/pkg/dragon_curve"
	"github.com/alikhanz/go-l-systems/pkg/fractal_plant"
	"github.com/alikhanz/go-l-systems/pkg/fractal_tree"
	"github.com/alikhanz/go-l-systems/pkg/koch_curve"
	"github.com/alikhanz/go-l-systems/pkg/sierpinski_triangle"
	"github.com/holizz/terrapin"
)

func main() {
	ft := fractal_tree.NewFractalTree()
	koch := koch_curve.NewKochCurve()
	sier := sierpinski_triangle.NewSierpinskiTriangle()
	dragon := dragon_curve.NewDragonCurve()
	fp := fractal_plant.NewFractalPlant()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		gen, err := strconv.Atoi(r.URL.Query().Get("gen"))
		if err != nil {
			gen = 2
		}

		alg := r.URL.Query().Get("alg")
		var t *terrapin.Terrapin

		switch alg {
			case "dragon":
				t, err = dragon.Render(gen)
				break
			case "koch":
				t, err = koch.Render(gen)
				break
			case "sier":
				t, err = sier.Render(gen)
				break
			case "ftree":
				t, err = ft.Render(gen)
				break
			case "fplant":
				t, err = fp.Render(gen)
				break
			default:
				t, err = ft.Render(gen)
				break
		}

		if err != nil {
			w.WriteHeader(500)
			return
		}

		_ = png.Encode(w, t.Image)
	})

	fmt.Println("Listening on http://localhost:3000")
	log.Fatalln(http.ListenAndServe(":3000", nil))
}