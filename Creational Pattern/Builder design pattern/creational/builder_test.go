package creational

import "testing"

func TestBuilderPattern(t *testing.T) {
	// tạo 1 ManufacturingDirector phụ trách việc tạo các vehicle trong suốt quá trình test
	manufacturingComplex := ManufacturingDirector{}

	// sau khi tạo thằng ManufacturingDirector chúng ta tạo thằng Carbuilder
	carBuilder := &CarBuilder{}

	// rồi chúng ta truyền vào dây chuyền chế tạo bằng các sử dụng SetBuilder method.
	manufacturingComplex.SetBuilder(carBuilder)
	
	// khi mà thằng ManufacturingDirector biết nó cần phải khởi tạo cái gì thì chúng ta gọi cái hàm Construct
	// để nó tạo VehicleProduct bằng cách sử dungj Carbuilder
	manufacturingComplex.Construct()
	
	
	car := carBuilder.GetVehicle()

	if car.Wheels != 4 {
		t.Errorf("Wheels on a car must be 4 and they were %d\n", car.Wheels)
	}
	if car.Structure != "Car" {
		t.Errorf("Structure on a car must be 'Car' and was %s\n",
		car.Structure)
	}
	if car.Seats != 5 {
		t.Errorf("Seats on a car must be 5 and they were %d\n", car.Seats)
	}

	// tạo cho bike
	bikeBuilder := &BikeBuilder{}
	manufacturingComplex.SetBuilder(bikeBuilder)
	manufacturingComplex.Construct()
	motorbike := bikeBuilder.GetVehicle()
	motorbike.Seats = 1

	if motorbike.Wheels != 2 {
		t.Errorf("Wheels on a motorbike must be 2 and they were %d\n", motorbike.Wheels)
	}

	if motorbike.Structure != "Motorbike" {
		t.Errorf("Structure on a motorbike must be 'Motorbike' and was %s\n", motorbike.Structure)
	}

}
