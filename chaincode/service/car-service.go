package service

import (
	"github.com/shijinshi/mfbobc/chaincode/dao"
	"github.com/shijinshi/mfbobc/chaincode/bean"
)

type CarService interface {
	CreateCar(car bean.Car)
	ChangeCarOwner(id, newOwner string)
}

type CarServiceImpl struct {
	Dao dao.CarDao`mytag:"CarDao"`
}

func (this CarServiceImpl)CreateCar(car bean.Car)  {
	this.Dao.Add(car)
}

func (this CarServiceImpl)ChangeCarOwner(id, newOwner string){
	car := this.Dao.Inquire(id)
	car.Owner = newOwner
	this.Dao.Update(car)
}