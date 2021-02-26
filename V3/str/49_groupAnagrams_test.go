package str

import "testing"

func TestGroupAnagrams(t *testing.T) {

}

/*
	这里会有两种方式：
	1. 相同字母的单词 排序后一定相同，所以可以 对strs中的每个字符串排序，若是相同则统计起来
	2. 相同字母的单词，单词出现的频率一定是一样的，所以我们可以统计每个单词中字母出现的频率 若出现相同的则 合并到一起
*/
func groupAnagrams(strs []string) [][]string {

	if len(strs) < 1 {
		return nil
	}

	resMap := make(map[[26]int][]string, len(strs))

	for i := 0; i < len(strs); i++ {

		str := strs[i]
		count := [26]int{}

		for j := 0; j < len(str); j++ {
			count[str[j]-'a']++
		}
		resMap[count] = append(resMap[count], str)
	}

	res := make([][]string, 0, len(strs))
	for _, v := range resMap {
		res = append(res, v)
	}

	return res

}
