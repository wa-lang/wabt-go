// 版权 @2022 凹语言 作者。保留所有权利。

// wabt 可执行程序包装
package wabt

import _ "embed"

// Wabt 版本号
const Version = "1.0.29"

//go:embed internal/wabt-1.0.29-macos/bin/wat2wasm
var wat2wasm_macos string

//go:embed internal/wabt-1.0.29-ubuntu/bin/wat2wasm
var wat2wasm_ubuntu string

//go:embed internal/wabt-1.0.29-windows/bin/wat2wasm.exe
var wat2wasm_windows string

// 读取 wat2wasm 命令
func LoadWat2Wasm_macos() []byte {
	return []byte(wat2wasm_macos)
}

// 读取 wat2wasm 命令
func LoadWat2Wasm_ubuntu() []byte {
	return []byte(wat2wasm_ubuntu)
}

// 读取 wat2wasm 命令数据
func LoadWat2Wasm_windows() []byte {
	return []byte(wat2wasm_windows)
}
