// Created by zrcoder at 2023/06/08 16:29
// https://leetcode.cn/problems/regular-expression-matching/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func isMatch(s string, p string) bool {

}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	s := Deserialize[string](ReadLine(stdin))
	p := Deserialize[string](ReadLine(stdin))
	ans := isMatch(s, p)

	fmt.Println("\noutput:", Serialize(ans))
}
