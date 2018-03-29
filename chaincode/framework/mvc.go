package framework

import (
	"github.com/szdenny/ccexample/chaincode/service"
	"github.com/szdenny/ccexample/chaincode/dao"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/szdenny/ccexample/chaincode/controller"
	"reflect"
	sc "github.com/hyperledger/fabric/protos/peer"
)

func BuildService(APIstub shim.ChaincodeStubInterface) service.CarService {
	carDao := dao.CarDaoImpl{
		TableName: "CAR",
		APIstub:   APIstub,
	}
	sc := service.CarServiceImpl{
		Dao: carDao,
	}

	return sc
}

func BuildController(APIstub shim.ChaincodeStubInterface) controller.CarController {
	return controller.CarController{
		S: BuildService(APIstub),
	}

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