package main

import "fmt"

func main() {
	GenerateRsaKey(4096)
	src := []byte("我是小崔, 如果我死了, 肯定不是自杀...我是小崔, 如果我死了, 肯定不是自杀...我是小崔, 如果我死了, 肯定不是自杀...我是小崔, 如果我死了, 肯定不是自杀...我是小崔, 如果我死了, 肯定不是自杀...")
	cipherText := RSAEncrypt(src, "public.pem")
	//fmt.Println(string(cipherText))
	plainText := RSADecrypt(cipherText, "private.pem")
	fmt.Println(string(plainText))
}
