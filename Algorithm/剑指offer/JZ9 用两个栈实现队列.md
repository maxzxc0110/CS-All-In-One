# 描述
用两个栈来实现一个队列，使用n个元素来完成 n 次在队列尾部插入整数(push)和n次在队列头部删除整数(pop)的功能。 队列中的元素为int类型。保证操作合法，即保证pop操作时队列内已有元素。

数据范围： 

n≤1000
要求：存储n个元素的空间复杂度为 

O(n) ，

插入与删除的时间复杂度都是 

O(1)

go
```
package main
 
var stack1 [] int
var stack2 [] int
 
func Push(node int) {
    stack1 = append(stack1, node)                //直接压栈，stack1无需判断是否栈满了，无长度限制
}
 
func Pop() int{
    if len(stack2) == 0 {                        //stack2为空时，移动stack1的内容到stack2
        for i := 0; i < len(stack1); i++ {        //将stack1的数据，从后往前，全部移到stack2中
            stack2 = append(stack2, stack1[len(stack1)-1-i])
        }
        stack1 = []int{}                        //stack1置空
    }
 
    res := stack2[len(stack2)-1]                 //获取最后一个位置作为栈顶元素
    stack2 = stack2[:len(stack2)-1]             //左闭右开，取得0~n-2的元素作为新的切片
    return res                                    //stack2弹出结果
}
```

c++
```
class Solution {
  public:
    void push(int node) {
        stack1.push(node);
    }

    int pop() {

        if (stack2.empty()) {
            while (!stack1.empty()) {
                stack2.push(stack1.top());
                stack1.pop();
            }
        }


        int res = stack2.top();
        stack2.pop();
        return res;
    }

  private:
    stack<int> stack1;
    stack<int> stack2;
};
```