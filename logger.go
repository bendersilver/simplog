package simplog

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"runtime"
	"runtime/debug"
	"strings"
	"time"
)

type lvl int

// Log levels.
const (
	LvlCRITICAL lvl = iota
	LvlERROR
	LvlWARNING
	LvlNOTICE
	LvlINFO
	LvlDEBUG
)

var maxlvl = LvlDEBUG
var file = os.Stderr

// auto init path
func init() {
	var f *os.File
	var err error
	, _ := os.Executable()
	var pth = []string{"./", path.Dir(os.Args[0]), exe}
	for _, p := range pth {
		f, err = os.OpenFile(path.Join(p, ".env"), os.O_RDONLY, 0644)
		if err != nil {
			continue
		} else {
			break
		}
	}
	if err != nil {
		return
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		spl := strings.Split(scanner.Text(), "=")
		if len(spl) > 1 && spl[0] == "LOG_PATH" {
			SetPath(spl[1])
		}
	}
}

func write(lv lvl, s string, skip int) {
	if maxlvl < lv {
		return
	}
	switch lv {
	case LvlCRITICAL:
		file.WriteString("\033[35mC ")
	case LvlERROR:
		file.WriteString("\033[31mE ")
	case LvlWARNING:
		file.WriteString("\033[33mW ")
	case LvlNOTICE:
		file.WriteString("\033[32mN ")
	case LvlINFO:
		file.WriteString("\033[37mI ")
	case LvlDEBUG:
		file.WriteString("\033[36mD ")
	}
	file.WriteString(time.Now().Format("2006-01-02 15:04:05.999"))
	_, fn, line, _ := runtime.Caller(2)

	file.WriteString(fmt.Sprintf(" %s:%d ▶ \033[0m %s", path.Base(fn), line, s))
}

// Debug -
func Debug(v ...interface{}) {
	write(LvlDEBUG, fmt.Sprintln(v...), 2)
}

// Debugf -
func Debugf(format string, a ...interface{}) {
	write(LvlDEBUG, fmt.Sprintf(format, a...), 2)
}

// Info -
func Info(v ...interface{}) {
	write(LvlINFO, fmt.Sprintln(v...), 2)
}

// Infof -
func Infof(format string, a ...interface{}) {
	write(LvlINFO, fmt.Sprintf(format, a...), 2)
}

// Notice -
func Notice(v ...interface{}) {
	write(LvlNOTICE, fmt.Sprintln(v...), 2)
}

// Noticef -
func Noticef(format string, a ...interface{}) {
	write(LvlNOTICE, fmt.Sprintf(format, a...), 2)
}

// Warning -
func Warning(v ...interface{}) {
	write(LvlWARNING, fmt.Sprintln(v...), 2)
}

// Warningf -
func Warningf(format string, a ...interface{}) {
	write(LvlWARNING, fmt.Sprintf(format, a...), 2)
}

// Error -
func Error(v ...interface{}) {
	write(LvlERROR, fmt.Sprintln(v...), 2)
}

// Errorf -
func Errorf(format string, a ...interface{}) {
	write(LvlERROR, fmt.Sprintf(format, a...), 2)
}

// Fatal -
func Fatal(v ...interface{}) {
	write(LvlCRITICAL, fmt.Sprintln(v...), 2)
	file.Close()
	os.Exit(1)
}

// Fatalf -
func Fatalf(format string, a ...interface{}) {
	write(LvlCRITICAL, fmt.Sprintf(format, a...), 2)
	file.Close()
	os.Exit(1)
}

// Recover -
func Recover(e *error) {
	if err := recover(); err != nil {
		*e = fmt.Errorf("%v\n%s", err, debug.Stack())
	}
}

// Close -
func Close() {
	file.Close()
}

// SetMaxLevel - set loggin max level
func SetMaxLevel(l lvl) {
	maxlvl = l
}

// SetPath - set path loller file
func SetPath(p string) {
	os.Mkdir(p, os.ModePerm)
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	flName := strings.Join([]string{path.Base(ex), "log"}, ".")
	file, err = os.OpenFile(path.Join(p, flName), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
}
