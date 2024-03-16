# 描述
定义栈的数据结构，请在该类型中实现一个能够得到栈中所含最小元素的 min 函数，输入操作时保证 pop、top 和 min 函数操作时，栈中一定有元素。

此栈包含的方法有：
push(value):将value压入栈中
pop():弹出栈顶元素
top():获取栈顶元素
min():获取栈中最小元素


# 解法 双栈

step 1：使用一个栈记录进入栈的元素，正常进行push、pop、top操作。
step 2：使用另一个栈记录每次push进入的最小值。
step 3：每次push元素的时候与第二个栈的栈顶元素比较，若是较小，则进入第二个栈，若是较大，则第二个栈的栈顶元素再次入栈，因为即便加了一个元素，它依然是最小值。于是，每次访问最小值即访问第二个栈的栈顶。

go
```
package main

var stack1 []int
var stack2 []int

func Push(node int) {
	// write code here
	stack1 = append(stack1, node)
	if len(stack2) == 0 {
		stack2 = append(stack2, node)
		return
	}
	if node > stack2[len(stack2)-1] {
		stack2 = append(stack2, stack2[len(stack2)-1])
	} else {
		stack2 = append(stack2, node)
	}
	return

}
func Pop() {
	// write code here
    stack1=stack1[:len(stack1)-1]
    stack2=stack2[:len(stack2)-1]

}
func Top() int {
	// write code here
      return stack1[len(stack1)-1]

}
func Min() int {
	// write code here
    return stack2[len(stack2)-1]
}

```

c++
```
class Solution {
public:
stack<int> s1;
stack<int> s2;
    void push(int value) {
        s1.push(value);

        if(s2.empty()|| s2.top() >value){
            s2.push(value);
        }else{
            s2.push(s2.top());
        }

    }
    void pop() {
        s1.pop();
        s2.pop();
    }
    int top() {
        return s1.top();
    }
    int min() {
        return s2.top();
    }
};
```