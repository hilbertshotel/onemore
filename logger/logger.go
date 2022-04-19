package logger

import (
    "log"
    "os"
    "runtime"
    "fmt"
    "path/filepath"
)

type Logger struct {
    Error func(error)
    Ok func(string)
}

func Init() *Logger {
    okLog := log.New(os.Stdout, "OK ", log.Ldate|log.Ltime)
    errLog := log.New(os.Stderr, "ERROR ", log.Ldate|log.Ltime)
    
    return &Logger{

        Error: func(err error) {
            _, file, line, _ := runtime.Caller(1)
            msg := fmt.Sprintf("(%v %v) %v", filepath.Base(file), line, err)
            errLog.Println(msg)
        },

        Ok: func(msg string) {
            okLog.Println(msg)
        },

    }
}

