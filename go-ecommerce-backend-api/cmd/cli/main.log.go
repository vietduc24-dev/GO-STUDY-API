package main

import (
	"os"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap"
)

func main() {
	// sugar := zap.NewExample().Sugar()
	// sugar.Infof("Hello name : %s , age : %d", "Duc", 20)

	// //logger, _ := zap.NewProduction()
    // logger := zap.NewExample()
	// logger.Info("Hello name : ", zap.String("name", "Duc"), zap.Int("age", 20))
     
	// logger := zap.NewExample()
	// logger.Info("Hello name Example")

	// logger, _ = zap.NewDevelopment()
	// logger.Info("Hello name Development")

    // logger, _= zap.NewProduction()
	// logger.Info("Hello name Production")

    //custom 
	encoder := getEncoderLog()
	sync := getWriterSync()
	core := zapcore.NewCore(encoder, sync, zapcore.InfoLevel)
	logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
    logger.Info("Info log", zap.Int("line",1))
	logger.Info("Info log", zap.Int("line",2))

}

func getEncoderLog() zapcore.Encoder {
	encodeConfig := zap.NewProductionEncoderConfig()
	//1791466738.846522 => 2026-10-08T22:38:58.845+0900
	encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	//ts -> time
	encodeConfig.TimeKey = "time"
	//from info Info

	encodeConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	//"caller":cli/main/log.go:10
	encodeConfig.EncodeCaller = zapcore.ShortCallerEncoder //
	
	return zapcore.NewConsoleEncoder(encodeConfig)

}

func getWriterSync() zapcore.WriteSyncer {
	// file, _ := os.OpenFile(name, flag, perm)
	file, _ := os.OpenFile("./log/log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	syncFile := zapcore.AddSync(file)
    syncConsole := zapcore.AddSync(os.Stderr) 
	return zapcore.NewMultiWriteSyncer(syncConsole, syncFile)

}	