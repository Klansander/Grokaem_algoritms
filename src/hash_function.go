package main
import (
    "fmt"
    "os"
    "bufio"
    )
 
func main() {
    n,count,count1:=0,1,1
	flag:=true
    in := bufio.NewReader(os.Stdin)
    fmt.Fscan(in,&n)
	Map :=make(map[int]int)
    for p:=0;p<n;p++ {
        fmt.Fscan(in,&count)
        mas :=make([]int,count)
        for i:=0;i<count;i++ {
            fmt.Fscan(in,&mas[i])
			Map[mas[i]]+=1
        }
        for j:=0;j<count-1;j++ {
			if mas[j]==mas[j+1] {
				count1++
			} else {
				if Map[mas[j]]!=count1 {
					count1=1
					flag=false
					break
				}
				count1=1
			}
            
        }
        if flag {
           fmt.Println("YES") 
        } else {
            fmt.Println("NO")
        }
        flag =true
        mas=nil
		Map=nil
    }
}