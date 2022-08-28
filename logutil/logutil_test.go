package logutil

import (
	"fmt"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	log := New(LevelDebug)
	{
		log.Debugf("hello %v", time.Now())
		log.Infof("hello %v", time.Now())
		log.Warnf("hello %v", time.Now())
		log.Errorf("hello %v", time.Now())
		log.Fatalf("hello %v", time.Now())
		fmt.Println()
	}
	{
		log.SetLevel(LevelInfo)
		log.Debugf("hello %v", time.Now())
		log.Infof("hello %v", time.Now())
		log.Warnf("hello %v", time.Now())
		log.Errorf("hello %v", time.Now())
		log.Fatalf("hello %v", time.Now())
		fmt.Println()
	}
	{
		log.SetLevel(LevelWarn)
		log.Debugf("hello %v", time.Now())
		log.Infof("hello %v", time.Now())
		log.Warnf("hello %v", time.Now())
		log.Errorf("hello %v", time.Now())
		log.Fatalf("hello %v", time.Now())
		fmt.Println()
	}
	{
		log.SetLevel(LevelError)
		log.Debugf("hello %v", time.Now())
		log.Infof("hello %v", time.Now())
		log.Warnf("hello %v", time.Now())
		log.Errorf("hello %v", time.Now())
		log.Fatalf("hello %v", time.Now())
		fmt.Println()
	}
	{
		log.SetLevel(LevelFatal)
		log.Debugf("hello %v", time.Now())
		log.Infof("hello %v", time.Now())
		log.Warnf("hello %v", time.Now())
		log.Errorf("hello %v", time.Now())
		log.Fatalf("hello %v", time.Now())
		fmt.Println()
	}
}

func TestDefaultLogger(t *testing.T) {
	SetLevel(LevelDebug)
	Debugf("hello %v", time.Now())
	Infof("hello %v", time.Now())
	Warnf("hello %v", time.Now())
	Errorf("hello %v", time.Now())
	Fatalf("hello %v", time.Now())
	fmt.Println()

	SetLevel(LevelInfo)
	Debugf("hello %v", time.Now())
	Infof("hello %v", time.Now())
	Warnf("hello %v", time.Now())
	Errorf("hello %v", time.Now())
	Fatalf("hello %v", time.Now())
	fmt.Println()

	SetLevel(LevelWarn)
	Debugf("hello %v", time.Now())
	Infof("hello %v", time.Now())
	Warnf("hello %v", time.Now())
	Errorf("hello %v", time.Now())
	Fatalf("hello %v", time.Now())
	fmt.Println()

	SetLevel(LevelError)
	Debugf("hello %v", time.Now())
	Infof("hello %v", time.Now())
	Warnf("hello %v", time.Now())
	Errorf("hello %v", time.Now())
	Fatalf("hello %v", time.Now())
	fmt.Println()

	SetLevel(LevelFatal)
	Debugf("hello %v", time.Now())
	Infof("hello %v", time.Now())
	Warnf("hello %v", time.Now())
	Errorf("hello %v", time.Now())
	Fatalf("hello %v", time.Now())
	fmt.Println()
}
