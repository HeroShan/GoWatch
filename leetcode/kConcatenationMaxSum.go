package leetcode
//https://leetcode-cn.com/problems/k-concatenation-maximum-sum/

func KConcatenationMaxSum(arr []int, k int) int {
	 if(arr == nil || len(arr)==0){
		 return 0
	 }
	 var(
		newarr []int
		length		int
	 )
	 length = len(arr)-1
	 if k > 1{
		if(arr[length]>0){
			newarr = append(newarr,arr[length])
		}
		for _,v := range arr{
			newarr = append(newarr,v)
		}
		if(arr[0]>0){
		   newarr = append(newarr,arr[0])
		}
	 }else{
		newarr = arr
	 }
	 return 1
	

	 
	 

	
}

func getMax(arr []int) {

}