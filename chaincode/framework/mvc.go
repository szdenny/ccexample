package framework

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"reflect"
	sc "github.com/hyperledger/fabric/protos/peer"
	"fmt"
	"errors"
	"strings"
)

type TypeRegister map[string]reflect.Type

func (t TypeRegister) Set(name string, value reflect.Type) {
	//name string, typ reflect.Type
	t[name] = value
}

func (t TypeRegister) GetBean(name string, APIstub shim.ChaincodeStubInterface) (reflect.Value, error) {
	if tt, ok := t[name]; ok {
		bean := reflect.New(tt)
		for i := 0; i < tt.NumField(); i++ {
			field := tt.Field(i)
			beanName := field.Tag.Get("bean")

			if beanName != "" {
				refBean, _ := TypeReg.GetBean(beanName, APIstub)
				if refBean != reflect.ValueOf(nil) {
					ff := bean.Elem().Field(i)

					ff.Set(refBean.Elem())
				}
			}

			value := field.Tag.Get("value")
			if value != "" {
				ff := bean.Elem().Field(i)
				ff.SetString(value)
			}

			stub := field.Tag.Get("stub")
			if stub != ""{
				ff := bean.Elem().Field(i)
				ff.Set(reflect.ValueOf(APIstub))
			}
		}
		return bean, nil
	}
	return reflect.ValueOf(nil), errors.New("no one")
}

var TypeReg = make(TypeRegister)

func BuildController(APIstub shim.ChaincodeStubInterface) reflect.Value {
	function, _ := APIstub.GetFunctionAndParameters()

	controller, e := TypeReg.GetBean(strings.Split(function, ".")[0], APIstub)
	if e != nil {
		return reflect.ValueOf(nil)
	}

	return controller
}

func Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {
	// Retrieve the requested Smart Contract function and arguments
	function, args := APIstub.GetFunctionAndParameters()

	fmt.Printf("%s %s\n", args[0], args)
	controller := BuildController(APIstub)

	method := controller.MethodByName(strings.Split(function, ".")[1])

	a := make([]reflect.Value, len(args))
	for i := 0; i < len(args); i++ {
		a[i] = reflect.ValueOf(args[i])
	}
	method.Call([]reflect.Value{reflect.ValueOf(args)})
	return shim.Success(nil)
}
