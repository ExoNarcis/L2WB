package pattern

import (
	"fmt"
)

// абстракция
type _Aircraft struct {
	speed  float64
	mass   float64
	lenght float64
	width  float64
}

func (a *_Aircraft) Fly() {
	//..
}

// некая функциональность
func (a *_Aircraft) GetSpeed() float64 {
	return a.speed
}
func (a *_Aircraft) SetSpeed(newspeed float64) {
	a.speed = newspeed
}
func (a *_Aircraft) GetMass() float64 {
	return a.mass
}
func (a *_Aircraft) SetMass(newmass float64) {
	a.mass = newmass
}
func (a *_Aircraft) GetLenght() float64 {
	return a.lenght
}
func (a *_Aircraft) SetLenght(newlenght float64) {
	a.lenght = newlenght
}
func (a *_Aircraft) GetWidth() float64 {
	return a.width
}
func (a *_Aircraft) SetWidth(newwidth float64) {
	a.width = newwidth
}

// говорим что нас может посещать визитор с таким интерфейсом
func (a *_Aircraft) Accept(av AirABSVisitor) {
	av.Visit(a)
}

// наследник
type AirBomber struct {
	_Aircraft
	//...
}

func NewAirBomber(_speed, _mass, _lenght, _width float64) *AirBomber {
	a := AirBomber{}
	a.speed = _speed
	a.mass = _mass
	a.lenght = _lenght
	a.width = _width
	return &a
}

// наследник
type AirBus struct {
	_Aircraft
	//...
}

func NewAirBus(_speed, _mass, _lenght, _width float64) *AirBus {
	a := AirBus{}
	a.speed = _speed
	a.mass = _mass
	a.lenght = _lenght
	a.width = _width
	return &a
}

// Интерфейс для визитора
type AircraftGetter interface {
	GetSpeed() float64
	GetMass() float64
	GetLenght() float64
	GetWidth() float64
	Accept(AirABSVisitor)
}

// связываем визитор с интерфесом который реализован в _Aircraft
type AirABSVisitor interface {
	Visit(AircraftGetter)
}

// структура реализует AirABSVisitor
type AirVisitor struct {
	speed  float64
	mass   float64
	lenght float64
	width  float64
}

// метод посещения
func (av *AirVisitor) Visit(ag AircraftGetter) {
	av.speed = ag.GetSpeed()
	av.mass = ag.GetMass()
	av.lenght = ag.GetLenght()
	av.width = ag.GetWidth()
}

// просто печать
func (av *AirVisitor) PrintAir() {
	fmt.Println("speed:", av.speed, " mass:", av.mass, " lenght:", av.lenght, " width:", av.width)
}

// Визитор паттерн для работы с объектами разных типов, прекрасно когда стек типов уже определен так как позволит изменять сам визитор
// добавлять новые методы для работы с классами и т.д. беда только 1 при добавлении новых классов в визитор (допустимых для него)
// нужно будет всю иерархию классов тянуть
// bomb := pattern.NewAirBomber(1.0, 1.0, 1.0, 1.0)
// bus := pattern.NewAirBus(2.0, 2.0, 2.0, 2.0)
// vis := &pattern.AirVisitor{}
// bomb.Accept(vis)
// vis.PrintAir()
// bus.Accept(vis)
// vis.PrintAir()
// OUT:
// speed: 1  mass: 1  lenght: 1  width: 1
// speed: 2  mass: 2  lenght: 2  width: 2
