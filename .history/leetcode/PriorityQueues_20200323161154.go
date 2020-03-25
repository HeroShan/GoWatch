package leetcode

type Queue struct{
	Index		int
	Value		int
	Priority	int
}

func Initqueue() *Queue {
	 var q *Queue
	 q = new(Queue)
	 return q
}