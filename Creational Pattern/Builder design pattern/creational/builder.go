package creational



// BuildProcess định hướng các step để build 1 phương tiện
// Mỗi 1 builder phải triển khai cái interface này nếu nó muốn được dùng trong sản xuất. Trong mỗi bước set,
// chúng ta trả về các bước build giống nhau, nên chúng ta có thể xâu chuỗi các bước khác nhau lại với nhau trong
// 1 câu lệnh. Cuối cùng chúng ta cần có method GetVehicle để lấy biến `Vehicle` từ thằng builder ra  
type BuildProcess interface {
	SetWheels() BuildProcess
	SetSeats() BuildProcess
	SetStructure() BuildProcess
	GetVehicle() VehicleProduct
}

// ManufacturingDirector là thằng phụ trách việc chấp nhận các builders. Nó có method construct mà sẽ dùng thằng builder để lưu trữ
// trong Manufacturing, và sẽ tái sử dụng các step cần thiết.
type ManufacturingDirector struct {
	// ManufacturingDirector cần 1 builder để store cái builder đang được sử dụng
	builder BuildProcess
}

// Construct a
func (f *ManufacturingDirector) Construct() {
	f.builder.SetSeats().SetStructure().SetWheels()
 //Implementation goes here
}

// SetBuilder cho phép chúng ta thay đổi các builder mà đang được sử dụng bởi thằng ManufacturingDirector
func (f *ManufacturingDirector) SetBuilder(b BuildProcess) {
	f.builder = b
 //Implementation goes here
}

type VehicleProduct struct {
	Wheels int
	Seats int
	Structure string
}


// CarBuilder thằng builder đầu tiên là thằng CarBuilder. Nó phải triển khai mọi method mà được định nghĩa trong
// interface BuildProcess
type CarBuilder struct {
	v VehicleProduct
}

// SetWheels a
func (c *CarBuilder) SetWheels() BuildProcess {
	c.v.Wheels = 4
 	return c

}

// SetSeats a
func (c *CarBuilder) SetSeats() BuildProcess {
 	c.v.Seats = 5
 	return c

}

// SetStructure a
func (c *CarBuilder) SetStructure() BuildProcess {
 	c.v.Structure = "Car"
 	return c
}

// GetVehicle a
func (c *CarBuilder) GetVehicle() VehicleProduct {
 	return c.v
}


// BikeBuilder tạo 1 struct BikeBuilder
type BikeBuilder struct {
	v VehicleProduct
}

// SetWheels các method bình thường
func (b *BikeBuilder) SetWheels() BuildProcess {
	b.v.Wheels = 2
	return b
}

// SetSeats các method bình thường
func (b *BikeBuilder) SetSeats() BuildProcess {
	b.v.Seats = 2
	return b
}

// SetStructure các method bình thường
func (b *BikeBuilder) SetStructure() BuildProcess {
	b.v.Structure = "Motorbike"
	return b
}

// GetVehicle các method bình thường
func (b *BikeBuilder) GetVehicle() VehicleProduct {
	return b.v
}