package main

import (
	"testing"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"fmt"
)

var cc = new(SmartContract)
var stub = shim.NewMockStub("model_cc", cc)
var dictArg = `"id","name"`

func TestChainCode(t *testing.T) {
	args := make([][]byte, 0)
	args = append(args, []byte("Car.ChangeOwner"))
	args = append(args, []byte("id"))
	args = append(args, []byte("new owner"))

	response := stub.MockInvoke("1", args)

	fmt.Printf("%v", response)
}