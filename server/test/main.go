package main

import (
	"fmt"
	"runtime"
)

func level3() {
	pc := make([]uintptr, 10)               // 10個分の情報を格納できるスライスを確保
	n := runtime.Callers(0, pc)             // 現在のスタックフレームから情報を取得
	frames := runtime.CallersFrames(pc[:n]) // pcからフレーム情報を取得

	fmt.Printf("スタックトレース:\n")
	for {
		frame, more := frames.Next()
		fmt.Printf("- %s:%d %s\n", frame.File, frame.Line, frame.Function)
		if !more {
			break
		}
	}
}

func level2() {
	level3()
}

func level1() {
	level2()
}

func main() {
	level1()
}
