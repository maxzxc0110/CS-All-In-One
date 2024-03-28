# 描述
给定一个二叉树其中的一个结点，请找出中序遍历顺序的下一个结点并且返回。注意，树中的结点不仅包含左右子结点，同时包含指向父结点的next指针。下图为一棵有9个节点的二叉树。树中从父节点指向子节点的指针用实线表示，从子节点指向父节点的用虚线表示


示例:
输入:{8,6,10,5,7,9,11},8
返回:9
解析:这个组装传入的子树根节点，其实就是整颗树，中序遍历{5,6,7,8,9,10,11}，根节点8的下一个节点就是9，应该返回{9,10,11}，后台只打印子树的下一个节点，所以只会打印9，如下图，其实都有指向左右孩子的指针，还有指向父节点的指针，下图没有画出来

数据范围：节点数满足 
1
≤
�
≤
50
 
1≤n≤50  ，节点上的值满足 
1
≤
�
�
�
≤
100
 
1≤val≤100 

要求：空间复杂度 
�
(
1
)
 
O(1)  ，时间复杂度 
�
(
�
)
 
O(n) 
输入描述：
输入分为2段，第一段是整体的二叉树，第二段是给定二叉树节点的值，后台会将这2个参数组装为一个二叉树局部的子树传入到函数GetNext里面，用户得到的输入只有一个子树根节点
返回值描述：
返回传入的子树根节点的下一个节点，后台会打印输出这个节点

# 解法

思路：

我们首先要根据给定输入中的结点指针向父级进行迭代，直到找到该树的根节点；然后根据根节点进行中序遍历，当遍历到和给定树节点相同的节点时，下一个节点就是我们的目标返回节点

具体做法：

step 1：首先先根据当前给出的结点找到根节点
step 2：然后根节点调用中序遍历
step 3：将中序遍历结果存储下来
step 4：最终拿当前结点匹配是否有符合要求的下一个结点


```
var ans []*TreeLinkNode

func InTraversal(pNode *TreeLinkNode) {
	if pNode == nil {
		return
	}
	InTraversal(pNode.Left)
	ans = append(ans, pNode)
	InTraversal(pNode.Right)
}

func GetNext(pNode *TreeLinkNode) *TreeLinkNode {
	
	var root *TreeLinkNode = pNode

	for root.Next != nil {
		root = root.Next
	}
	InTraversal(root)
	for i := 0; i < len(ans)-1; i++ {
		cur := ans[i]
		if pNode == cur {
			return ans[i+1]
		}
	}

	return nil

}
```