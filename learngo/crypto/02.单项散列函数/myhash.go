package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// 使用sha256
func myHash() {
	// sha256.Sum256([]byte("hello, go"))

	// 1. 创建哈希接口对象
	myHash := sha256.New()
	// 2. 添加数据
	src := []byte("我是小崔, 如果我死了, 肯定不是自杀...我是小崔, 如果我死了, 肯定不是自杀...我是小崔, 如果我死了, 肯定不是自杀...我是小崔, 如果我死了, 肯定不是自杀...我是小崔, 如果我死了, 肯定不是自杀...")
	myHash.Write(src)
	myHash.Write(src)
	myHash.Write(src)
	// 3. 计算结果
	res := myHash.Sum(nil)
	// 4. 格式化为16进制形式
	myStr := hex.EncodeToString(res)
	fmt.Printf("%s\n", myStr)
}

/*
82161fa95a
92e15caed3
6850a5d28a
40bf069225
24fd209200
91f867da95
467f
*/
