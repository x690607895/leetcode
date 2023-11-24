package main

import "log"

// 给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。

// 给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。

// 示例 1：

// 输入：digits = "23"
// 输出：["ad","ae","af","bd","be","bf","cd","ce","cf"]
// 示例 2：

// 输入：digits = ""
// 输出：[]
// 示例 3：

// 输入：digits = "2"
// 输出：["a","b","c"]

// 作者：力扣 (LeetCode)
// 链接：https://leetcode.cn/leetbook/read/top-interview-questions-medium/xv8ka1/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
func main() {
	digits := "23"
	log.Println(letterCombinations(digits))
}
func letterCombinations(digits string) []string {
	result := make([]string, 0)
	temp := ""
	var dfs func(idx int)
	dfs = func(idx int) {
		if idx == len(digits) {
			if len(temp) == len(digits) {
				result = append(result, temp)
			}
			return
		}
		staticdata := getStaticData(rune(digits[idx]))
		for _, v := range staticdata {
			oldTemp := temp
			temp += v
			dfs(idx + 1)
			temp = oldTemp
		}
	}
	for k := range digits {
		dfs(k)
	}
	return result
}

func getStaticData(src rune) []string {
	switch src {
	case '2':
		return []string{"a", "b", "c"}
	case '3':
		return []string{"d", "e", "f"}
	case '4':
		return []string{"g", "h", "i"}
	case '5':
		return []string{"j", "k", "l"}
	case '6':
		return []string{"m", "n", "o"}
	case '7':
		return []string{"p", "q", "r", "s"}
	case '8':
		return []string{"t", "u", "v"}
	case '9':
		return []string{"w", "x", "y", "z"}
	default:
		return []string{}
	}
}
