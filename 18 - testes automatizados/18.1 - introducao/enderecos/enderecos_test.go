// Teste de unidade

package enderecos

import "testing"

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestTipoDeEndereco(t *testing.T) {
	t.Parallel()
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

// go test --coverprofile cobertura.txt = gera um arquivo.txt com informacoes do resultado do teste
// go tool cover --func=cobertura.txt = deixa mais facil de ler o arquivo.txt no terminal
// go tool cover --html=cobertura.txt = gera um arquivo HTML

func TestQualquer(t *testing.T) {
	t.Parallel()
	if 1 > 2 {
		t.Errorf("Teste quebrou!")
	}
}
