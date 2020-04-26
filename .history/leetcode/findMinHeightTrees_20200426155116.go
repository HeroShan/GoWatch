package leetcode
// 20  18  	17  16 		98  93			无向树  求最大深度根节点
//    9      	  8 		  7	
//    		5			4
//    			  3
//1.遍历节点，如果可以连续就深度加一   否则就新增一个链表 生成完链表层层递减，直到层数为1，如果是只剩一个链表，则返回链表，否则返回链表前指针的值
//2.去掉每一层周围的节点为二的节点   就找到最中心的根节点
// {20,9}	{18,9}	{17,8}	{16,8}	{98,7}	{93,7}
// {5,9}	{5,8}	{4,8}	{4,7}	{3,5}	{3,4}
import (
	"container/list"
)
func FindMinHeightTrees(n int, edges [][]int) []int {
    degree := make(map[int]int)
    graph := make(map[int]*list.List)
    
    Q := list.New()
    ans := make([]int,0)
    
    for _, v := range edges {
        if graph[v[0]] == nil {
            graph[v[0]] = list.New()
        }
        if graph[v[1]] == nil {
            graph[v[1]] = list.New()
        }
        graph[v[0]].PushBack(v[1])
        graph[v[1]].PushBack(v[0])
        
        degree[v[0]] += 1
        degree[v[1]] += 1
    }
    for k,v := range degree {
        if v == 1 {
            Q.PushBack(k)
        }
    }
    rst := n
    if rst == 1 {
        ans = append(ans,0)
    }
    
    for rst != 1 && rst != 2 {
        l := Q.Len()
        rst -= l
        
        for i:=0;i<l;i++ {
            cur := Q.Front()
            Q.Remove(cur)
            for t:=graph[cur.Value.(int)].Front(); t!=nil;t=t.Next() {
                temp := t.Value.(int)
                if degree[temp] > 0 {
                    degree[temp]--
                }
                if degree[temp] == 1 {
                    Q.PushBack(temp)
                }
            }
        }
    }
    for Q.Len() > 0 {
        ans = append(ans, Q.Front().Value.(int))
        Q.Remove(Q.Front())
    }
    return ans
}