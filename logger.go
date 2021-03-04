package simplog

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"
	"time"
)

type lvl int

// Log levels.
const (
	CRITICAL lvl = iota
	ERROR
	WARNING
	NOTICE
	INFO
	DEBUG
)

var file = os.Stdout

// auto init path
func init() {
	f, err := os.OpenFile("./.env", os.O_RDONLY, 0644)
	if err != nil {
		return
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		spl := strings.Split(scanner.Text(), "=")
		if len(spl) > 1 && spl[0] == "LOG_PATH" {
			SetPath(spl[1])
		}
	}
}

func write(lv lvl, s string) {
	switch lv {
	case CRITICAL:
		file.WriteString("\033[35mC ")
	case ERROR:
		file.WriteString("\033[31mE ")
	case WARNING:
		file.WriteString("\033[33mW ")
	case NOTICE:
		file.WriteString("\033[32mN ")
	case INFO:
		file.WriteString("\033[37mI ")
	case DEBUG:
		file.WriteString("\033[36mD ")
	}
	file.WriteString(time.Now().Format("2006-01-02 15:04:05.999"))
	_, fn, line, _ := runtime.Caller(2)

	file.WriteString(fmt.Sprintf(" %s:%d ▶ \033[0m %s", path.Base(fn), line, s))
}

// Debug -
func Debug(v ...interface{}) {
	write(DEBUG, fmt.Sprintln(v...))
}

// Debugf -
func Debugf(format string, a ...interface{}) {
	write(DEBUG, fmt.Sprintf(format, a...))
}

// Info -
func Info(v ...interface{}) {
	write(INFO, fmt.Sprintln(v...))
}

// Infof -
func Infof(format string, a ...interface{}) {
	write(INFO, fmt.Sprintf(format, a...))
}

// Notice -
func Notice(v ...interface{}) {
	write(NOTICE, fmt.Sprintln(v...))
}

// Noticef -
func Noticef(format string, a ...interface{}) {
	write(NOTICE, fmt.Sprintf(format, a...))
}

// Warning -
func Warning(v ...interface{}) {
	write(WARNING, fmt.Sprintln(v...))
}

// Warningf -
func Warningf(format string, a ...interface{}) {
	write(WARNING, fmt.Sprintf(format, a...))
}

// Error -
func Error(v ...interface{}) {
	write(ERROR, fmt.Sprintln(v...))
}

// Errorf -
func Errorf(format string, a ...interface{}) {
	write(ERROR, fmt.Sprintf(format, a...))
}

// Fatal -
func Fatal(v ...interface{}) {
	write(CRITICAL, fmt.Sprintln(v...))
	file.Close()
	os.Exit(1)
}

// Fatalf -
func Fatalf(format string, a ...interface{}) {
	write(CRITICAL, fmt.Sprintf(format, a...))
	file.Close()
	os.Exit(1)
}

// Close -
func Close() {
	file.Close()
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
