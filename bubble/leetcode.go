package main

import (

	"fmt"
	"strconv"
)

type Server struct{
	worktime int
	Quests int
	Tag bool
}

func initServer() *Server {
	return &Server{worktime: 0,Quests: 0,Tag: true}
}
func selfDividingNumbers(left int, right int) (ans []int) {
	for left<=right{
		if judje(left)==true{
			ans = append(ans,left)
		}
		left++
	}
	return
}
func judje(num int) bool {
	str_num := strconv.Itoa(num)
	for _,j:=range str_num{
		a := string(j)
		tmp,_ := strconv.Atoi(a)
		if tmp==0{
			return false
		}
		if num%tmp!=0{
			return false
		}
	}
	return true
}
func main() {
	fmt.Println(selfDividingNumbers(47,85))
}

func busiestServer(k int, arrival []int, load []int) (ans []int) {
	//服务集群
	servers_group := make([]Server,k)
	for i:=0;i<k;i++{
		servers_group[i] = *initServer()
	}
	for i:=0;i<len(arrival);i++{
		tmp := i%k
		for {
			//服务器请求完成
			if servers_group[tmp].Tag == false{
				if arrival[i]>=servers_group[tmp].worktime{
					servers_group[tmp].Tag = true
				}
			}
			if servers_group[tmp].Tag==true{
				servers_group[tmp].Tag = false
				servers_group[tmp].worktime = load[i]+arrival[i]
				servers_group[tmp].Quests++
				break
			}else{
				tmp++
				if tmp==k{
					tmp = 0
				}
				if tmp == (i%k){
					break
				}
			}

		}
	}
	fmt.Println(servers_group)
	mmax :=0
	for i,j:=range servers_group{
		if j.Quests==mmax{
			ans = append(ans,i)
		}else if j.Quests>mmax{
			ans = []int{i}
			mmax = j.Quests
		}
	}
	return
}