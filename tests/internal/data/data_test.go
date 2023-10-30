package data_test

import (
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	log.Println("Fazer algo antes dos testes")

	result := m.Run() // roda todos os testes contidos no package

	log.Println("Finalizando todos tests")

	os.Exit(result)
}
