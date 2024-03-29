package logging

import (
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/jkstack/jkframe/utils"
)

// DateRotateConfig date logger config
type DateRotateConfig struct {
	Level       Level  // log level
	Dir         string // log file save directory, only used if WriteFile=true
	Name        string // log file name, only used if WriteFile=true
	Rotate      int    // save rotated file count, only used if WriteFile=true
	WriteStdout bool   // allow write to stdout
	WriteFile   bool   // allow write to file
}

type rotateDateLogger struct {
	sync.Mutex
	date string
	cfg  DateRotateConfig

	// runtime
	f *os.File
	w *writer
}

// NewRotateDateLogger create logger by date rotate
func NewRotateDateLogger(cfg DateRotateConfig) Logger {
	var ws []io.Writer
	if cfg.WriteStdout {
		ws = append(ws, os.Stdout)
	}
	var f *os.File
	if cfg.WriteFile {
		os.MkdirAll(cfg.Dir, 0755)
		var err error
		f, err = os.OpenFile(filepath.Join(cfg.Dir, cfg.Name+".log"), os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
		utils.Assert(err)
		ws = append(ws, f)
	}
	if len(ws) == 0 {
		panic(errors.New("no output"))
	}
	var w io.Writer
	if len(ws) == 1 {
		w = ws[0]
	} else {
		w = io.MultiWriter(ws[0], ws[1])
	}
	return Logger{
		logger: &rotateDateLogger{
			date: time.Now().Format("20060102"),
			cfg:  cfg,
			f:    f,
			w:    newWriter(w),
		},
		lastCheck: time.Now(),
	}
}

// SetDateRotate set log rotate by date
func SetDateRotate(cfg DateRotateConfig) {
	DefaultLogger = NewRotateDateLogger(cfg)
}

func (l *rotateDateLogger) setLevel(level Level) {
	l.cfg.Level = level
}

func (l *rotateDateLogger) currentLevel() Level {
	return l.cfg.Level
}

func (l *rotateDateLogger) rotate() {
	now := time.Now().Format("20060102")
	if l.date == now {
		return
	}
	if !l.cfg.WriteFile {
		return
	}
	l.Lock()
	defer l.Unlock()
	files, _ := filepath.Glob(filepath.Join(l.cfg.Dir, l.cfg.Name+"_*.log"))
	for _, file := range files {
		date := strings.TrimPrefix(filepath.Base(file), l.cfg.Name+"_")
		date = strings.TrimSuffix(date, ".log")
		t, _ := time.Parse("20060102", date)
		if time.Since(t).Hours() > float64(24*l.cfg.Rotate) {
			os.Remove(file)
		}
	}
	l.f.Close()
	os.Rename(filepath.Join(l.cfg.Dir, l.cfg.Name+".log"),
		filepath.Join(l.cfg.Dir, l.cfg.Name+"_"+l.date+".log"))
	l.f, _ = os.OpenFile(filepath.Join(l.cfg.Dir, l.cfg.Name+".log"), os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	var w io.Writer
	if l.cfg.WriteStdout {
		w = io.MultiWriter(os.Stdout, l.f)
	} else {
		w = l.f
	}
	l.w = newWriter(w)
	l.date = now
}

func (l *rotateDateLogger) printf(fmt string, a ...interface{}) string {
	return l.w.Printf(fmt, a...)
}

func (l *rotateDateLogger) write(str string) {
	l.w.Write(str)
}

func (l *rotateDateLogger) flush() {
	f := l.f
	if f != nil {
		f.Sync()
	}
}

func (l *rotateDateLogger) files() []string {
	if !l.cfg.WriteFile {
		return nil
	}
	var ret []string
	ret = append(ret, filepath.Join(l.cfg.Dir, l.cfg.Name+".log"))
	files, err := filepath.Glob(filepath.Join(l.cfg.Dir, l.cfg.Name+"_*.log"))
	if err == nil {
		ret = append(ret, files...)
	}
	return ret
}
