// Basic Fatal, Info, and Debug logging

package logger

import(
	"os"
	"fmt"
	"time"
)

type LogLevel int
const (
        FATAL LogLevel = iota
        INFO
        DEBUG
)

var LOG_LEVEL LogLevel

func init() {
   LOG_LEVEL = INFO
}

type LogMessage struct {
	location		string	`json:"location"`
	level				string	`json:"level"`
	message			string	`json:"message"`
	timestamp		string	`json:"timestamp"`
}

func SetLogLevel(newLevel LogLevel) LogLevel {
	if(newLevel <= DEBUG && newLevel >= FATAL) {
		LOG_LEVEL = newLevel
	}
	return LOG_LEVEL
}

func now() string {
	return time.Now().Format(time.RFC3339)
}

func log(mes *LogMessage, stderr bool) {
	if(stderr) {
		fmt.Fprintf(os.Stderr, "[%s] %s (%s) %s\n", mes.timestamp, mes.level, mes.location, mes.message)
	} else {
		fmt.Fprintf(os.Stdout, "[%s] %s (%s) %s\n", mes.timestamp, mes.level, mes.location, mes.message)
	}
}

//Call only at top level after everything has been cleared up
func Fatal(m ...interface{}) {
	ci := retrieveCallInfo()
	log(&LogMessage{fmt.Sprint(ci.packageName, "/", ci.fileName, ":", ci.line), "FATAL", fmt.Sprint(m...), now()}, true)
	os.Exit(1)
}

//Messages directed at users
func Info(m ...interface{}) {
	if(LOG_LEVEL >= INFO) {
		ci := retrieveCallInfo()
		log(&LogMessage{ci.packageName, "INFO", fmt.Sprint(m...), now()}, false)
	}
}

//Messages directed at developers
func Debug(m ...interface{}) {
	if(LOG_LEVEL >= DEBUG) {
		ci := retrieveCallInfo()
		log(&LogMessage{fmt.Sprint(ci.packageName, "/", ci.fileName, ":", ci.line), "DEBUG", fmt.Sprint(m...), now()}, false)
	}
}
