// Teste de unidade

package enderecos

import (
	"testing"
)

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestTipoDeEndereco(t *testing.T) {
	// Test
	// enderecoParaTeste := "Avenida Paulista"

	// tipoDeEnderecoEsperado := "Avenida"

	// tipoDeEnderecoRecebido := TipoDeEndereco(enderecoParaTeste)
	// if tipoDeEnderecoRecebido != tipoDeEnderecoEsperado {
	// 	t.Error("O tipo recebido eh diferente do esperado")
	// }

	cenarioDeTeste := []cenarioDeTeste{
		{"Rua ABC", "Rua"},
		{"Avenida Paulista", "Avenida"},
		{"Rodovia do imigrantes", "Rodovia"},
		{"Estrada Qualquer", "Estrada"},
	}

	for _, cenario := range cenarioDeTeste {
		tipoDeEnderecoRecebido := TipoDeEndereco(cenario.enderecoInserido)
		if tipoDeEnderecoRecebido != cenario.retornoEsperado {
			t.Errorf("O tipo de recebido %s é diferente do esperado %s",
				tipoDeEnderecoRecebido,
				cenario.retornoEsperado,
			)
		}
	}
}
