package shapes

import (
	"math"
	"testing"
)

func TestArea(t *testing.T)  {
	t.Run("Retangulo", func(t *testing.T) {
		ret := Retangulo{ 10, 12 }
		areaEsperada := float64(120)
		areaRecebida := ret.Area()

		if areaEsperada != areaRecebida {
			t.Fatalf("A área esperada é diferente da aŕea recebida.\n Esperado: %f\n Recebido:%f",
				areaEsperada,
				areaRecebida,
			)
		}
	})

	t.Run("Circulo", func(t *testing.T) {
		circ := Circulo{ 10 }
		areaEsperada := float64(math.Pi * 100)
		areaRecebida := circ.Area()

		if areaEsperada != areaRecebida {
			t.Fatalf("A área esperada é diferente da aŕea recebida.\n Esperado: %f\n Recebido:%f",
				areaEsperada,
				areaRecebida,
			)
		}
	})
}