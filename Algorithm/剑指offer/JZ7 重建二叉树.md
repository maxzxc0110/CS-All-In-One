# 描述
给定节点数为 n 的二叉树的前序遍历和中序遍历结果，请重建出该二叉树并返回它的头结点。
例如输入前序遍历序列{1,2,4,7,3,5,6,8}和中序遍历序列{4,7,2,1,5,3,8,6}，则重建出如下图所示。


# 解法 递归建立二叉树

step 1：先根据前序遍历第一个点建立根节点。
step 2：然后遍历中序遍历找到根节点在数组中的位置。
step 3：再按照子树的节点数将两个遍历的序列分割成子数组，将子数组送入函数建立子树。
step 4：直到子树的序列长度为0，结束递归。

go
```
func reConstructBinaryTree( pre []int ,  vin []int ) *TreeNode {
    if len(pre) == 0 {
        return nil
    }
 
    root := &TreeNode{pre[0], nil, nil}
 
    i := 0
    for i = 0; i < len(vin); i++ {
        if vin[i] == pre[0] {
            break
        }
    }
 
    root.Left = reConstructBinaryTree(pre[1: len(vin[: i]) +1], vin[: i])
    root.Right = reConstructBinaryTree(pre[len(vin[: i]) +1: ], vin[i +1:])
 
    return root
}
```