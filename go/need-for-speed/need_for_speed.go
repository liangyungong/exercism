package speed

type Car struct {
	speed        int
	battery      int
	batteryDrain int
	distance     int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	var c Car
	c.speed = speed
	c.batteryDrain = batteryDrain
	c.battery = 100
	return c
}

type Track struct {
	distance int
}

// NewTrack creates a new track
func NewTrack(distance int) Track {
	var t Track
	t.distance = distance
	return t
}

// Drive drives the car one time. If there is not enough battery to drive one more time,
// the car will not move.
func Drive(car Car) Car {
	if car.battery >= car.batteryDrain {
		car.battery -= car.batteryDrain
		car.distance += car.speed
	}
	return car
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	for car.battery >= car.batteryDrain {
		car = Drive(car)
		track.distance -= car.speed
	}

	return track.distance <= 0
}
