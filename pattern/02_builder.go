package pattern

import (
	"fmt"
	"strings"
)

// интерфейс для абстракции от катапульты (когда нибудь все ровно прийдется делать еще балисту)
type CatapultManager interface {
	GetCatapultInfo() string
}

// тип снаряда катапульты
type _CatapultPackage struct {
	typ     string
	mass    float64
	fireamo bool
}

// некая катапульта
type _Catapult struct {
	pack            *_CatapultPackage
	typematerial    string
	country         string
	power           float64
	angletrajectory float64
}

// надо реализоать для катапульты наш интерфейс CatapultManager
func (c *_Catapult) GetCatapultInfo() string {
	sb := strings.Builder{}

	if c.pack.fireamo {
		sb.WriteString(c.country + " catapult shoots fire " + c.pack.typ + " mass " + fmt.Sprintf("%f", c.pack.mass))
	} else {
		sb.WriteString(c.country + " catapult shoots " + c.pack.typ + "s mass " + fmt.Sprintf("%f", c.pack.mass))
	}
	sb.WriteString(" material " + c.typematerial + " power:" + fmt.Sprintf("%f", c.power) + " agline:" + fmt.Sprintf("%f", c.angletrajectory))
	return sb.String()
}

// некий абстрактный конструктор катапульт
type _CatapultBuilder struct {
}

// методы создания снаряда и катапульты
func (cb *_CatapultBuilder) createPack(typ string, mass float64, fireamo bool) *_CatapultPackage {
	return &_CatapultPackage{
		typ:     typ,
		mass:    mass,
		fireamo: fireamo,
	}
}
func (cb *_CatapultBuilder) createCatapult(typemat, country string, power, angletrajectory float64, catpack *_CatapultPackage) *_Catapult {
	return &_Catapult{pack: catpack,
		typematerial:    typemat,
		country:         country,
		power:           power,
		angletrajectory: angletrajectory,
	}
}

// абстрактный интерфейс для билдеров катапульт
type _CatapultAbsBuilder interface {
	Create() CatapultManager
}

// римская катапульта (билдер)
type RomanCatapultBuilder struct {
	_CatapultBuilder
}

func (rcb *RomanCatapultBuilder) Create() CatapultManager {
	return rcb.createCatapult("Wood", "Roman", 10.0, 45.0, rcb.createPack("Rock", 10.0, false))
}

// византийская катапульта (билдер)
type ByzantiumCatapultBuilder struct {
	_CatapultBuilder
}

func (bcb *ByzantiumCatapultBuilder) Create() CatapultManager {
	return bcb.createCatapult("Wood", "Byzantium", 12.0, 10.0, bcb.createPack("Arrow", 5.0, true))
}

// Китайская катапульта (билдер) (Китай все еще существует)
type ChinaCatapultBuilder struct {
	_CatapultBuilder
}

func (ccb *ChinaCatapultBuilder) Create() CatapultManager {
	return ccb.createCatapult("Wood", "China", 14.0, 15.0, ccb.createPack("Fireworks", 2.0, true))
}

// универсальный руководитель катапульт по интерфейсу неизвестных билдеров катапульт
type CatapultDirector struct {
	builder _CatapultAbsBuilder
}

func (cd *CatapultDirector) Setbuilder(newbuilder _CatapultAbsBuilder) {
	cd.builder = newbuilder
}

func (cd *CatapultDirector) Build() CatapultManager {
	return cd.builder.Create()
}

func NewCatapultDirector(builder _CatapultAbsBuilder) *CatapultDirector {
	return &CatapultDirector{builder: builder}
}

// Паттерн приколен когда нужно создать единный make (руководитель) сборщик из некоторых (рецептов) билдеров
// скрывая как именно мы там создаем тот или иной TYPE следовательно мы точно знаем что вышло
// минусы - описывать процесс сборки для каждого нового TYPE,
// переписывать много при изменениях в базовой логике

// catb := &pattern.ByzantiumCatapultBuilder{}
// catc := &pattern.ChinaCatapultBuilder{}
// catr := &pattern.RomanCatapultBuilder{}
// catdir := pattern.NewCatapultDirector(catb)
// fmt.Println(catdir.Build().GetCatapultInfo())
// catdir.Setbuilder(catc)
// fmt.Println(catdir.Build().GetCatapultInfo())
// catdir.Setbuilder(catr)
// fmt.Println(catdir.Build().GetCatapultInfo())
// OUT:
// Byzantium catapult shoots fire Arrow mass 5.000000 material Wood power:12.000000 agline:10.000000
// China catapult shoots fire Fireworks mass 2.000000 material Wood power:14.000000 agline:15.000000
// Roman catapult shoots Rocks mass 10.000000 material Wood power:10.000000 agline:45.000000
