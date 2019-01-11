package main

import (
	"encoding/json"
	"fmt"
)

// 经常会使用结构体来转换成JSON。json包是通过反射机制来实现编解码的，
// 因此结构体必须导出所转换的字段，不导出的字段不会被json包解析

type DebugInfo struct {
	Level  string
	Msg    string
	author string // 未导出字段不会被json解析
}

func marshal2() {
	dbgInfos := []DebugInfo{
		DebugInfo{"debug", `File: "test.txt" not found,`, "ljs"},
		DebugInfo{"", "Logic error", "Gopher"},
	}

	if data, err := json.Marshal(dbgInfos); err == nil {
		fmt.Printf("%s\n", data)
	}
}
