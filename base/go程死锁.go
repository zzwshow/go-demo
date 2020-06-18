package main

// 死锁常见的有
/*
	1,单go程自己死锁
		channel 至少要在两个go程中读写,否则造成死锁(添加缓冲的channel 除外)
	2,go程channel访问顺序导致死锁
	3,多go程,多channel 交叉死锁
*/

//1,单go程自己死锁
//func main(){
//	ch := make(chan int) // 注,若给管道添加了缓冲区就不会死锁了
//	ch <- 123 //向管道写数据,写完后没人读,死锁,下面读的代码没机会执行了.
//
//	num := <- ch //从通道内读
//	fmt.Println("num = ",num)
//
//}


//2,go程channel访问顺序导致死锁
//func main(){
//	ch := make(chan int)
//	num := <- ch             // 这行是个读管道操作.主go程运行到这里会直接阻塞 下面的代码没机会执行了,(把这两行拿到下面去)
//	fmt.Println("num = ",num)
//
//	go func() {
//		ch <- 123
//	}()
//}


//3,多go程,多channel 交叉死锁
func main(){

}