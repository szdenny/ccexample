package business

import (
	"fmt"
	"github.com/szdenny/ccexample/chaincode/framework"
)

type CarController struct {
	S CarService `bean:"carService"`
}

func (this *CarController) ChangeOwner(args []string) {
	this.S.ChangeCarOwner(args[0], args[1])
	fmt.Println("i am here")
}

func init(){
	framework.TypeReg.Set("carControl", &CarController{})
}