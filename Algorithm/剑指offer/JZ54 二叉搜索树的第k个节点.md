

# 中序遍历

知识点：二叉树递归

递归是一个过程或函数在其定义或说明中有直接或间接调用自身的一种方法，它通常把一个大型复杂的问题层层转化为一个与原问题相似的规模较小的问题来求解。因此递归过程，最重要的就是查看能不能讲原本的问题分解为更小的子问题，这是使用递归的关键。

思路：

根据二叉搜索树的性质，左子树的元素都小于根节点，右子树的元素都大于根节点。因此它的中序遍历（左中右）序列正好是由小到大的次序，因此我们可以尝试递归中序遍历，也就是从最小的一个节点开始，找到k个就是我们要找的目标。

具体做法：

step 1：设置全局变量count记录遍历了多少个节点，res记录第k个节点。
step 2：另写一函数进行递归中序遍历，当节点为空或者超过k时，结束递归，返回。
step 3：优先访问左子树，再访问根节点，访问时统计数字，等于k则找到。
step 4：最后访问右子树。

go
```
var index int

func KthNode(proot *TreeNode, k int) int {
    res := midTravel(proot,k)

    if res == nil{
        return -1
    }
    return res.Val
}


func midTravel(node *TreeNode, targetPos int) *TreeNode{
    if node == nil{
        return nil
    }

    if res:= midTravel(node.Left, targetPos);res != nil{
        return res
    }
    index += 1
    if targetPos == index{
        return node
    }

    if res:= midTravel(node.Right, targetPos);res!=nil{
        return res
    }

    return nil

}

```