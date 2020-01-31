package main

import (
	"gitlab.paradise-soft.com.tw/glob/tracer/logs"
	"testing"
)

func Benchmark_main(b *testing.B) {
	//logger := logs.NewConsoleLogger(logs.LevelDebug)
	zaplogger := logs.NewConsoleLogger(logs.LevelDebug)



	//var wg sync.WaitGroup
	for i := 0; i < b.N; i++ {
		zaplogger.Warn("asvmw")
		zaplogger.Info("sldkcm")
		zaplogger.Error("sdcmls")
		zaplogger.Debug("abc")
	//	wg.Add(12)
	//	{
	//		go func() {
	//			logger.Info("Info")
	//			wg.Done()
	//		}()
	//		go func() {
	//			logger.Infof("Info")
	//			wg.Done()
	//		}()
	//		go func() {
	//			logger.Infow("Info")
	//			wg.Done()
	//		}()
	//	}
	//	{
	//		go func() {
	//			logger.Trace("Trace")
	//			wg.Done()
	//		}()
	//		go func() {
	//			logger.Tracef("Trace")
	//			wg.Done()
	//		}()
	//		go func() {
	//			logger.Tracew("Trace")
	//			wg.Done()
	//		}()
	//	}
	//	{
	//		go func() {
	//			logger.Warn("Warn")
	//			wg.Done()
	//		}()
	//		go func() {
	//			logger.Warnf("Warn")
	//			wg.Done()
	//		}()
	//		go func() {
	//			logger.Warnw("Warn")
	//			wg.Done()
	//		}()
	//	}
	//	{
	//		go func() {
	//			logger.Error("Error")
	//			wg.Done()
	//		}()
	//		go func() {
	//			logger.Errorf("Error")
	//			wg.Done()
	//		}()
	//		go func() {
	//			logger.Errorw("Error")
	//			wg.Done()
	//		}()
	//	}
	//	wg.Wait()
	}
}
