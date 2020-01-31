package main

import "gitlab.paradise-soft.com.tw/glob/tracer/logs"

func main() {
	console := logs.Console{
		Colors:logs.Colors{
		LevelDebug: logs.ColorGray,
		LevelInfo:  logs.ColorGreen,
		LevelWarn:  logs.ColorRed,
		LevelError: logs.ColorRed,
		LevelPanic: logs.ColorRed,
		},
	}

	file := logs.File{
		FileConfig:logs.FileConfig{
				FileName:   "",
				MaxSize:    0,
				MaxBackups: 0,
				MaxAge:     0,
				Compress:   false,
				},
	}
	writer := logs.Writer{
		Console: &console,
		File:    &file,
	}

	// New Logger
	// 設定Error Level
	logger := logs.New(logs.WithLevel(logs.LevelDebug), logs.WithWriter(writer))

	// Debug level
	logger.Debug( "aaa")
	logger.Debugf( "val: %s", "aaa")

	// Info level
	logger.Info( "bbb")
	logger.Infof("val: %s", "bbb")

	// Error level
	logger.Error("ddd")
	logger.Errorf( "val: %s", "ddd")

	// Panic level
	logger.Panic( "ddd")
	logger.Panicf("val: %s", "ddd")
}

