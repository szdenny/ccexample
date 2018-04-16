package business

import (
	"fmt"
	"github.com/szdenny/ccexample/chaincode/framework"
	"reflect"
)

type CarController struct {
	S CarServiceImpl `bean:"carService"`
}

func (this *CarController) ChangeOwner(args []string) {
	fmt.Printf("%s %s\n", args[0], args[0])
	this.S.ChangeCarOwner(args[0], args[1])
}

func init(){
	framework.TypeReg.Set("Car", reflect.TypeOf(CarController{}))
}