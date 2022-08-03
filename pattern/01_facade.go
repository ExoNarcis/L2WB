package pattern

import (
	"fmt"
)

// допустим микроподсистема которая делает n вещи в с файлами
type _FileIODirector struct {
	filefullpath string
}

func (f *_FileIODirector) WriteActionEnd() bool {
	fmt.Println("Writed in File")
	return true
}

// псевдоконструктор
func newFileIODirector(path string) *_FileIODirector {
	return &_FileIODirector{filefullpath: path}
}

// допустим микроподсистема которая делает n вещи в с базой данных
type _DBDirector struct {
	DBType, DBName, username, password string
}

func (d *_DBDirector) WriteAction() bool {
	fmt.Println("Writed in DB")
	return true
}

func (d *_DBDirector) DoSome() {
	fmt.Println("dosomedb")
}

// псевдоконструктор
func newDBDirector(Dbtype, Dbname, UserName, Password string) *_DBDirector {
	return &_DBDirector{DBType: Dbtype, DBName: Dbname, username: UserName, password: Password}
}

// допустим микроподсистема которая делает n вещи в с чем либо еще
type _DoSomeAction struct {
	//...
}

func (d *_DoSomeAction) DoSome() {
	fmt.Println("Do Some Action")
}

// псевдоконструктор
func newDoSomeAction() *_DoSomeAction {
	return &_DoSomeAction{}
}

// все эти подсистемы допустим требуется для общей работы с n процессом допустим записать в файл бд и еще куда либо
// считать из бд или файла и еще там логика возможно будет если то иначе и т.д.
// но работая с ними напрямую скорее всего мы упремся в бесмысленное повторение кода так еще и логику раскроем внешней среде
// для этого есть фасад - это обертка под конкретные реализации структур в гошке для создания единого объекта в идеале универсального
// даже если обращение не требуется ко всем подсистемам пусть оно будет в нашей обертке что бы не порезатся скажем так.
type ActionManager struct {
	f  *_FileIODirector
	d  *_DBDirector
	ds *_DoSomeAction
}

// тут нет особой логики но все же она возможна...
func (am *ActionManager) WriteAction() bool {
	am.ds.DoSome()
	return am.f.WriteActionEnd() && am.d.WriteAction()
}

// псевдоконструктор
func NewActionManager(filepath, Dbtype, Dbname, UserName, Password string) *ActionManager {
	return &ActionManager{
		f:  newFileIODirector(filepath),
		d:  newDBDirector(Dbtype, Dbname, UserName, Password),
		ds: newDoSomeAction(),
	}
}

// ac := pattern.NewActionManager("", "", "", "", "")
// ac.WriteAction()
// OUT:
// Do Some Action
// Writed in File
// Writed in DB
