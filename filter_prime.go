package main

import "fmt"

/**
筛法，先将一定范围内的数据全部放入origin channel，Processor函数取出第一个素数并打印，
然后筛选不能整除它的，即剩下的素数，并放入新的channel，不断一层层的递归处理，每一层开一个channel，
当最终channel里取不出来时即已经完成，可以退出goroutine,同时关闭wait channel，使得主函数
可以结束阻塞并返回
 */

func Processor(seq chan int, wait chan struct{})  {
	go func() {
		prime, ok := <-seq
		if !ok {
			close(wait)
			return
		}
		fmt.Println(prime)
		out := make(chan int)
		Processor(out, wait)
		for num:= range seq {
			if num%prime!=0 {
				out <- num
			}
		}
		close(out)
	}()
}

func main()  {
	origin, wait := make(chan int), make(chan struct{})
	Processor(origin, wait)
	for num:=2;num<100000 ;num++  {
		origin <- num
	}
	close(origin)
	<-wait   //组赛主函数，直到Processor处理完关闭wait
}
