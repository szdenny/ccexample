package framework

import (
	"testing"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func TestBuildController(t *testing.T) {
	BuildController(shim.NewMockStub("test", nil))
}
