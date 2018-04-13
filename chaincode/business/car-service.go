package business

import "github.com/szdenny/ccexample/chaincode/framework"

type CarService interface {
	CreateCar(car Car)
	ChangeCarOwner(id, newOwner string)
}

type CarServiceImpl struct {
	Dao CarDao `bean:"carDao"`
}

func (this *CarServiceImpl) CreateCar(car Car) {
	this.Dao.Add(car)
}

func (this *CarServiceImpl) ChangeCarOwner(id, newOwner string) {
	car := this.Dao.Inquire(id)
	car.Owner = newOwner
	this.Dao.Update(car)
}

func init() {
	framework.TypeReg.Set("carService", &CarServiceImpl{})
}
