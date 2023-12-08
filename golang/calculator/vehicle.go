package calculator

import (
	"errors"
)

type Vehicle interface {
	getVehicleType() VehicleType
}

type VehicleType string

const (
	Basic      VehicleType = ""
	Motorcycle VehicleType = "Motorcycle"
	Bus        VehicleType = "Bus"
	Tractor    VehicleType = "Tractor"
	Emergency  VehicleType = "Emergency"
	Diplomat   VehicleType = "Diplomat"
	Foreign    VehicleType = "Foreign"
	Military   VehicleType = "Military"
)

var ErrWrongTypeOfVehicle = errors.New("The supplied vehicle type isn't supported")

type OtherVehicle struct {
	Type VehicleType
}

func (ot OtherVehicle) getVehicleType() VehicleType {
	return ot.Type
}

func ParseVehicleType(vehtype VehicleType) (Vehicle, error) {
	var err error
	var veh Vehicle
	switch vehtype {
	case Basic:
		veh = Car{}
	case Bus:
		veh = OtherVehicle{vehtype}
	case Diplomat:
		veh = OtherVehicle{vehtype}
	case Emergency:
		veh = OtherVehicle{vehtype}
	case Foreign:
		veh = OtherVehicle{vehtype}
	case Military:
		veh = OtherVehicle{vehtype}
	case Motorcycle:
		veh = Motorbike{}
	default:
		err = ErrWrongTypeOfVehicle
	}

	return veh, err
}
