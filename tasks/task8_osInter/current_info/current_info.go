package currentinfo

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/inancgumus/screen"
	ps "github.com/mitchellh/go-ps"
)

// IsValid checks is path valid and writable
func IsValid(fp string) bool {
	// Check if file already exists
	if _, err := os.Stat(fp); err == nil {
		return true
	}

	// Attempt to create it
	var d []byte
	if err := os.WriteFile(fp, d, 0644); err == nil {
		os.Remove(fp) // And delete it
		return true
	}

	return false
}

// CurInfo struct about current info we have
type CurInfo struct {
	_addStr  string
	_hasStr  bool
	_curPath string
}

// ConstructorCurInfo returns CurInfo
func ConstructorCurInfo() CurInfo {
	var toRet CurInfo
	toRet._curPath = "D:/"
	toRet._hasStr = false
	return toRet
}

// PrintNewScreen prints new screen
func (ci CurInfo) PrintNewScreen() {
	screen.Clear()
	if ci._hasStr {
		fmt.Println(ci._addStr)
	}
	fmt.Print(ci._curPath)
}

// GetPath get path from currentInfo
func (ci CurInfo) GetPath() string {
	return ci._curPath
}

// PS command for shell PS
func (ci *CurInfo) PS() {
	var sb strings.Builder

	processList, err := ps.Processes()
	if err != nil {
		ci._addStr = err.Error()
		ci._hasStr = true
		return
	}
	for x := range processList {
		var process ps.Process
		process = processList[x]
		//log.Printf("%d\t%s\n",process.Pid(),process.Executable())
		sb.WriteString(fmt.Sprintf("%d\t%s\n", process.Pid(), process.Executable()))
	}
	ci._addStr = sb.String()
	ci._hasStr = true
}

// TryAddDir trying add subdirectory to current directory
func (ci *CurInfo) TryAddDir(addDir string) {
	if IsValid(ci._curPath + addDir + "/") {
		ci._curPath += (addDir + "/")
		ci._hasStr = false
	} else if IsValid(addDir) {
		ci._curPath = addDir
		ci._hasStr = false
	} else {
		ci._addStr = "incorrect directory"
		ci._hasStr = true
	}
}

// TryRemoveDir trying remove subdirectory to current directory
func (ci *CurInfo) TryRemoveDir() {
	var i int = len(ci._curPath) - 2
	for i != -1 && ci._curPath[i] != '/' {
		i--
	}
	if i == -1 {
		ci._addStr = "there is no subcatallog"
		ci._hasStr = true
	} else {
		ci._curPath = ci._curPath[:i+1]
		ci._hasStr = false
	}
}

// KillProcessByName kill process bu name (all)
func (ci *CurInfo) KillProcessByName(procname string) {
	kill := exec.Command("taskkill", "/im", procname, "/T", "/F")
	err := kill.Run()
	if err != nil {
		ci._addStr = err.Error()
		ci._hasStr = true
	} else {
		ci._hasStr = false
	}
}

// Echo is like cmd echo
func (ci *CurInfo) Echo(newAddStr string) {
	ci._addStr = newAddStr
	ci._hasStr = true
}
