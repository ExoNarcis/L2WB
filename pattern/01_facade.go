package pattern

import (
	"fmt"
)

type _FileIODirector struct {
	filefullpath string
}

func (f *_FileIODirector) WriteActionEnd() bool {
	fmt.Println("Writed in File")
	return true
}

func newFileIODirector(path string) *_FileIODirector {
	return &_FileIODirector{filefullpath: path}
}

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

func newDBDirector(Dbtype, Dbname, UserName, Password string) *_DBDirector {
	return &_DBDirector{DBType: Dbtype, DBName: Dbname, username: UserName, password: Password}
}

type _DoSomeAction struct {
	//...
}

func (d *_DoSomeAction) DoSome() {
	fmt.Println("Do Some Action")
}

func newDoSomeAction() *_DoSomeAction {
	return &_DoSomeAction{}
}

type ActionManager struct {
	f  *_FileIODirector
	d  *_DBDirector
	ds *_DoSomeAction
}

func (am *ActionManager) WriteAction() bool {
	am.ds.DoSome()
	return am.f.WriteActionEnd() && am.d.WriteAction()
}

func NewActionManager(filepath, Dbtype, Dbname, UserName, Password string) *ActionManager {
	return &ActionManager{
		f:  newFileIODirector(filepath),
		d:  newDBDirector(Dbtype, Dbname, UserName, Password),
		ds: newDoSomeAction(),
	}
}
