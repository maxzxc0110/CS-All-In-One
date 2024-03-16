# 描述
请实现一个函数，将一个字符串s中的每个空格替换成“%20”。
例如，当字符串为We Are Happy.则经过替换之后的字符串为We%20Are%20Happy。

数据范围:

0≤len(s)≤1000 。保证字符串中的字符为大写英文字母、小写英文字母和空格中的一种。

# 解法

## 辅助栈
go
```
func replaceSpace(s string) string {

	var ans string

	temp := strings.Split(s, " ")

	for i := 0; i < len(temp); i++ {
		ans += temp[i]
		if i != len(temp)-1 {
			ans += "%20"
		}
	}

	return ans

}
```

## 字符截取相加
c++
```
class Solution {
public:
    /**
     * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
     *
     * 
     * @param s string字符串 
     * @return string字符串
     */
    string replaceSpace(string s) {
       
        string res = "";

        for(int i=0;i<s.length();i++){
            if (s[i] ==' '){
                res += "%20";
            }else{
                res += s[i];
            }
        }

        return res;


    }
};
```