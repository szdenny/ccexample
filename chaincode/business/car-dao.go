package business

import (
	"fmt"
	"encoding/json"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/szdenny/ccexample/chaincode/framework"
)

type CarDao interface {
	Add(car Car)
	Inquire(id string) Car
	Update(car Car)
}

type CarDaoImpl struct {
	APIstub shim.ChaincodeStubInterface
	TableName string
}

func (this *CarDaoImpl)Add(car Car){
	carAsBytes, _ := json.Marshal(car)
	this.APIstub.PutState(this.TableName+car.ID, carAsBytes)
	fmt.Println("Dao.Add called")
}

func (this *CarDaoImpl)Inquire(id string) Car{
	carAsBytes, _ := this.APIstub.GetState(this.TableName+id)
	car := Car{}

	json.Unmarshal(carAsBytes, &car)
	return car
}

func (this *CarDaoImpl)Update(car Car){
	carAsBytes, _ := json.Marshal(car)
	this.APIstub.PutState(car.ID, carAsBytes)
}

func init(){
	framework.TypeReg.Set("carDao", &CarDaoImpl{})
}
