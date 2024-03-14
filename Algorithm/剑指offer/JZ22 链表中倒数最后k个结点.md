# 描述
输入一个长度为 n 的链表，设链表中的元素的值为 ai ，返回该链表中倒数第k个节点。
如果该链表长度小于k，请返回一个长度为 0 的链表。


# 解法 1 快慢指针

go
```
func FindKthToTail(pHead *ListNode, k int) *ListNode {
	// write code here
	fast := pHead
    slow := pHead

    for i:=0;i<k;i++{
        if fast == nil{
            return nil
        }

        fast = fast.Next
    }

    for fast !=nil{
        fast = fast.Next
        slow = slow.Next
    }

    return slow

}
```

c++
```
class Solution {
  public:
    /**
     * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
     *
     *
     * @param pHead ListNode类
     * @param k int整型
     * @return ListNode类
     */
    ListNode* FindKthToTail(ListNode* pHead, int k) {
        // write code here

        ListNode* fast = pHead;
        ListNode* slow = pHead;



        for (int i = 0; i < k; i++) {
            if (fast == NULL) {
                return NULL;
            }
            fast = fast->next;
        }

        while (fast != NULL) {
            fast = fast->next;
            slow = slow->next;
        }

        return slow;

    }
};
```