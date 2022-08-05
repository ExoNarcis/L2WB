package pattern

import (
	"fmt"
)

// интерфейс некой машины
type Mashine interface {
	On()
	Off()
	Restart()
}

// наша POWER машина
type PowerMashine struct {
	//...
}

// реализуем интерфейс
func (rw *PowerMashine) On() {
	fmt.Println("PowerMashine ON")
}

func (rw *PowerMashine) Off() {
	fmt.Println("PowerMashine Off")
}

func (rw *PowerMashine) Restart() {
	fmt.Println("PowerMashine Restart")
}

// интерфейс некекой команды
type _Command interface {
	Exec()
}

// менеджер который выполняет некую команду
type ComandManager struct {
	command _Command
}

// метод для смены команды
func (cm *ComandManager) SetCommand(newcommand _Command) {
	cm.command = newcommand
}

// запуск команды
func (cm *ComandManager) Start() {
	cm.command.Exec()
}

// структуры команд реазующий интерфейс
type On struct {
	Mashine Mashine
}

func (o *On) Exec() {
	o.Mashine.On()
}

type Off struct {
	Mashine Mashine
}

func (o *Off) Exec() {
	o.Mashine.Off()
}

type Restart struct {
	Mashine Mashine
}

func (r *Restart) Exec() {
	r.Mashine.Restart()
}

// Паттерн комманда строит представление методов как обстракцию объектов что позволяет манипулировать методами как объектами
// и при этом строить всякие истории коммант их отмены и. т.д.
// pw := pattern.PowerMashine{}
// on := pattern.On{&pw}
// off := pattern.Off{&pw}
// r := pattern.Restart{&pw}
// cm := pattern.ComandManager{}
// cm.SetCommand(&on)
// cm.Start()
// cm.SetCommand(&off)
// cm.Start()
// cm.SetCommand(&r)
// cm.Start()
// OUT:
// PowerMashine ON
// PowerMashine Off
// PowerMashine Restart
