package framework

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"reflect"
	sc "github.com/hyperledger/fabric/protos/peer"
	"fmt"
	"errors"
)

type BeanRegistry map[string]interface{}
type TypeRegister map[string]interface{}

func (t TypeRegister) Set(name string, value interface{}) {
	//name string, typ reflect.Type
	t[name] = value
}

func (t TypeRegister) GetBean(name string) (interface{}, error) {
	if bean, ok := BeanReg[name]; ok {
		return bean, nil
	}

	if bean, ok := t[name]; ok {
		tt := reflect.TypeOf(bean).Elem()
		vv := reflect.ValueOf(bean).Elem()
		for i := 0; i < tt.NumField(); i++ {
			field := tt.Field(i)
			beanName := field.Tag.Get("bean")

			if beanName != "" {
				refBean, _ := TypeReg.GetBean(beanName)
				if refBean != nil {
					field := vv.Field(i)

					field.Set(reflect.ValueOf(refBean))
				}
			}
		}
		BeanReg[name] = bean
		return bean, nil
	}
	return nil, errors.New("no one")
}

var TypeReg = make(TypeRegister)
var BeanReg = make(BeanRegistry)

func BuildController(APIstub shim.ChaincodeStubInterface) interface{} {
	function, _ := APIstub.GetFunctionAndParameters()

	controller, e := TypeReg.GetBean(function)
	if e != nil {
		fmt.Println(e)
		return nil
	}

	return controller
}

func Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {
	// Retrieve the requested Smart Contract function and arguments
	function, args := APIstub.GetFunctionAndParameters()

	c := BuildController(APIstub)

	r := reflect.ValueOf(&c)
	m := r.MethodByName(function)

	m.Call([]reflect.Value{reflect.ValueOf(args)})
	return shim.Error("Invalid Smart Contract function name.")
}
