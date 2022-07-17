package test

import (
	"mall/internal/service"
	"testing"
)

func TestEs(t *testing.T) {
	var s = service.EsProductService
	s.ImportAll()
}
