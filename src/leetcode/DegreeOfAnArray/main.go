package main

/**
  题目链接 https://leetcode.com/problems/degree-of-an-array/description/
 */
import "fmt"

func main(){
	fmt.Print("hello")
	nums:=[]int{1,2,2,3,1}
	println(findShortestSubArray(nums))
}


func findShortestSubArray(nums []int) int {

	maxCount:=1;
	maxCountValue:=nums[0];
	myMap:=make(map[int]*Node)
	for i,num:=range nums{
		val,ok:=myMap[num]
		if(ok){
			myMap[num].counts++
			myMap[num].end=i;
			if val.counts>maxCount{
				maxCount=val.counts;
				maxCountValue=num;
			}

			if(val.counts==maxCount){
				if val.end-val.begin<myMap[maxCountValue].end-myMap[maxCountValue].begin {
					maxCountValue=num;
				}
			}
		}else{
          node:=Node{1,i,i};
          myMap[num]=&node;
		}




	}

	return myMap[maxCountValue].end-myMap[maxCountValue].begin+1;

}

type  Node struct {
	counts int; //节点出现次数
    begin int   //节点首次开始的位置
    end  int    //节点结束的位置
}