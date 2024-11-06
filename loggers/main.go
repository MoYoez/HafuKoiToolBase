package loggers

import (
	"fmt"
	"log"
	"os"
	"time"
)

/*

	Colorful Logger
	but you can use this https://github.com/charmbracelet/log, it may better.


*/

// Color codes
const (
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorPurple = "\033[35m"
	colorCyan   = "\033[36m"
	colorWhite  = "\033[37m"
	colorReset  = "\033[0m"
	colorGray   = "\033[90m"
)

// ColorLogger is a custom logger with color output
type ColorLogger struct {
	logger *log.Logger
	color  string
}

var (
	// SuccessfulLogger logger Successful is Successful logger
	SuccessfulLogger = NewColorLogger(os.Stdout, "INFO:", log.Ldate|log.Ltime, colorGreen)

	// StandardLogger is a basic logger that writes to stdout in White
	StandardLogger = NewColorLogger(os.Stdout, "INFO: ", log.Ldate|log.Ltime, colorGray)

	// ErrorLogger is a logger that writes to stderr in red
	ErrorLogger = NewColorLogger(os.Stderr, "ERROR: ", log.Ldate|log.Ltime, colorRed)

	// WarningLogger is a logger that writes to stdout in yellow
	WarningLogger = NewColorLogger(os.Stdout, "WARNING: ", log.Ldate|log.Ltime, colorYellow)
)

func init() {

	// Update existing loggers to use the multi-writer
	StandardLogger = &ColorLogger{
		logger: log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime),
		color:  colorGray,
	}
	WarningLogger = &ColorLogger{
		logger: log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime),
		color:  colorYellow,
	}
	ErrorLogger = &ColorLogger{
		logger: log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime),
		color:  colorRed,
	}
	SuccessfulLogger = &ColorLogger{
		logger: log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime),
		color:  colorGreen,
	}

	// Set up a goroutine to rotate logs daily
	go func() {
		for {
			nextDay := time.Now().Add(24 * time.Hour)
			nextDay = time.Date(nextDay.Year(), nextDay.Month(), nextDay.Day(), 0, 0, 0, 0, nextDay.Location())
			timer := time.NewTimer(time.Until(nextDay))
			<-timer.C
			// Update loggers with new multi-writers
			StandardLogger = &ColorLogger{
				logger: log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
				color:  colorGreen,
			}
			WarningLogger = &ColorLogger{
				logger: log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile),
				color:  colorYellow,
			}
			ErrorLogger = &ColorLogger{
				logger: log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
				color:  colorRed,
			}
		}
	}()
}

// NewColorLogger creates a new ColorLogger
func NewColorLogger(out *os.File, prefix string, flag int, color string) *ColorLogger {
	return &ColorLogger{
		logger: log.New(out, prefix, flag),
		color:  color,
	}
}

// Println prints a line with color
func (cl *ColorLogger) Println(v ...interface{}) {
	cl.logger.Println(cl.color + fmt.Sprint(v...) + colorReset)
}

func (cl *ColorLogger) Printf(format string, v ...interface{}) {
	cl.logger.Printf(cl.color + fmt.Sprintf(format, v...) + colorReset)
}
