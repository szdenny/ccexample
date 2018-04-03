package framework

import (
	"github.com/szdenny/ccexample/chaincode/service"
	"github.com/szdenny/ccexample/chaincode/dao"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/szdenny/ccexample/chaincode/controller"
	"reflect"
	sc "github.com/hyperledger/fabric/protos/peer"
	"errors"
	"fmt"
)

type TypeRegister map[string]reflect.Type

func (t TypeRegister) Set(beanName string,interfaceName interface{}) {
	//name string, typ reflect.Type
	t[beanName] = reflect.TypeOf(interfaceName)
}

func (t TypeRegister) Get(name string) (interface{}, error) {
	if typ, ok := t[name]; ok {
		return reflect.New(typ).Elem().Interface(), nil
	}
	return nil, errors.New("no one")
}

var TypeReg = make(TypeRegister)

func BuildController(APIstub shim.ChaincodeStubInterface) interface{}{
	controller, error :=  TypeReg.Get("")
	if error != nil {
		fmt.Println(error)
		return nil
	}

	return controller
}

func Invoke(APIstub shim.ChaincodeStubInterface) sc.Response{
	// Retrieve the requested Smart Contract function and arguments
	function, args := APIstub.GetFunctionAndParameters()

	c := BuildController(APIstub)

	r := reflect.ValueOf(&c)
	m := r.MethodByName(function)

	m.Call([]reflect.Value{reflect.ValueOf(args)})
	return shim.Error("Invalid Smart Contract function name.")
}