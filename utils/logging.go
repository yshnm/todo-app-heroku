package utils

import (
	"io"
	"log"
	"os"
)

func LoggingSettings(logFile string) {
	// ログファイルを読み込んで、読み書き、作成、追記を設定。パーミッションを0666に設定
	logfile, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln(err)
	}
	// 書き込み先を標準出力とログファイルに設定
	multiLogFile := io.MultiWriter(os.Stdout, logfile)
	log.SetFlags(log.Ldate|log.Ltime|log.Lshortfile)

	log.SetOutput(multiLogFile)
}