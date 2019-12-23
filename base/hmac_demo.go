package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
)

/*
hmac算法是加密的hash算法,它需要一个hash算法（比如sha256获取md5等）和一个密匙key，在hash计算的过程中将密匙key混入，产生一个和原来hash算法相同位数的hash值。
在大多数情况下，我们甚至可以将hamc算法看做是加盐的hash算法（加盐是将一个随机字符串放在需要加密的密文前面或者后面，然后对这个拼接后的密文进行加密得到hash值）。
但它们的加密原理肯定不一样，虽然达到的效果是一样的，都是对密文混入一个第三方值，然后得到一个hash值。
*/

func CheckMAC(msg, msgMAC, key []byte) bool {
	mac := hmac.New(sha256.New,key) // 创建hash加密算法
	mac.Write(msg) // 写入数据
	expectedMAC := mac.Sum(nil) //获取加密后的hash
	fmt.Println("new: ",expectedMAC)
	fmt.Println("old: ",msgMAC)
	return hmac.Equal(expectedMAC,msgMAC) // 比较预期的hash和实际的hash
}



func main(){
	// 对sha256算法进行hash加密，key随便设置
	hash := hmac.New(sha256.New,[]byte("abc124")) // 创建对应的sha256哈希加密算法
	hash.Write([]byte("abc124"))	 // 写入加密数据
	fmt.Printf("%x\n",hash.Sum(nil))
	
	// b := CheckMAC([]byte("abc124"),[]byte("20d9f32758340ab6d82f54d83fdd66810824c743393a513f6b7ea58ef222dbf8"),[]byte("abc124"))
	
	// fmt.Println(b)
	
	
	// // 对md5算法进行hash加密，key随便设置
	// hash = hmac.New(md5.New,[]byte("abc123666")) // 创建对应的md5哈希加密算法
	// hash.Write([]byte("abc123")) // 写入加密数据
	// fmt.Printf("%x\n",hash.Sum(nil)) // 0eee86e484505ec4ab48c18095e6a8ac
}
