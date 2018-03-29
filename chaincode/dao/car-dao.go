package dao

import (
	"fmt"
	"encoding/json"
	"github.com/shijinshi/mfbobc/chaincode/bean"
	"github.com/hyperledger/fabric/core/chaincode/shim"

)

type CarDao interface {
	Add(car bean.Car)
	Inquire(id string) bean.Car
	Update(car bean.Car)
}

type CarDaoImpl struct {
	APIstub shim.ChaincodeStubInterface
	TableName string
}

func (this CarDaoImpl)Add(car bean.Car){
	carAsBytes, _ := json.Marshal(car)
	this.APIstub.PutState(this.TableName+car.ID, carAsBytes)
	fmt.Println("Dao.Add called")
}

func (this CarDaoImpl)Inquire(id string) bean.Car{
	carAsBytes, _ := this.APIstub.GetState(this.TableName+id)
	car := bean.Car{}

	json.Unmarshal(carAsBytes, &car)
	return car
}

func (this CarDaoImpl)Update(car bean.Car){
	carAsBytes, _ := json.Marshal(car)
	this.APIstub.PutState(car.ID, carAsBytes)
}
