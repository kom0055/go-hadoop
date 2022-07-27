package log

import "log"

var (
	Tracef = func(format string, args ...any) {
		if len(args) == 0 {
			log.Println(format)
			return
		}
		log.Printf(format, args)
	}
	Debugf = func(format string, args ...any) {
		if len(args) == 0 {
			log.Println(format)
			return
		}
		log.Printf(format, args)
	}
	Infof = func(format string, args ...any) {
		if len(args) == 0 {
			log.Println(format)
			return
		}
		log.Printf(format, args)
	}
	Warnf = func(format string, args ...any) {
		if len(args) == 0 {
			log.Println(format)
			return
		}
		log.Printf(format, args)
	}
	Errorf = func(format string, args ...any) {
		if len(args) == 0 {
			log.Println(format)
			return
		}
		log.Printf(format, args)
	}
	Fatalf = func(format string, args ...any) {
		if len(args) == 0 {
			log.Fatalln(format)
			return
		}
		log.Fatalf(format, args)
	}
)

func SetTracef(f func(format string, args ...any)) {
	Tracef = f
}

func SetDebugf(f func(format string, args ...any)) {
	Debugf = f
}

func SetInfof(f func(format string, args ...any)) {
	Infof = f
}

func SetWarnf(f func(format string, args ...any)) {
	Warnf = f
}

func SetErrorf(f func(format string, args ...any)) {
	Errorf = f
}

func SetFatalf(f func(format string, args ...any)) {
	Fatalf = f
}
