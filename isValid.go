package main

import "log"

// 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。

// 有效字符串需满足：

// 左括号必须用相同类型的右括号闭合。
// 左括号必须以正确的顺序闭合。

// 示例 1：

// 输入：s = "()"
// 输出：true
// 示例 2：

// 输入：s = "()[]{}"
// 输出：true
// 示例 3：

// 输入：s = "(]"
// 输出：false
// 示例 4：

// 输入：s = "([)]"
// 输出：false
// 示例 5：

// 输入：s = "{[]}"
// 输出：true
//

// 作者：力扣 (LeetCode)
// 链接：https://leetcode.cn/leetbook/read/top-interview-questions-easy/xnbcaj/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

func main() {
	s := "([])"
	log.Println(isValid(s), '[', ']', '{', '}', '(', ')')
}

func isValid(s string) bool {
	if len(s)%2 == 1 {
		return false
	}

	struck := make([]rune, 0)
	for _, v := range s {
		switch v {
		case '[', '{', '(':
			struck = append(struck, v)
			break
		case ']', '}', ')':
			if len(struck) <= 0 {
				return false
			}
			end := struck[len(struck)-1]
			if end+2 == v || end+1 == v {
				struck = struck[:len(struck)-1]
			} else {
				return false
			}
			break
		default:
			break
		}
	}
	if len(struck) == 0 {
		return true
	}
	return false
}
