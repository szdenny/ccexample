package controller

import (
	"fmt"
	"github.com/szdenny/ccexample/chaincode/service"
)

type CarController struct {
	S service.CarService
}

func (this *CarController) ChangeOwner(args []string) {
	this.S.ChangeCarOwner(args[0], args[1])
	fmt.Println("i am here")
}
