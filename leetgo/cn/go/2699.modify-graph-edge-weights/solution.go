// Created by zrcoder at 2023/06/09 07:43
// https://leetcode.cn/problems/modify-graph-edge-weights/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func modifiedGraphEdges(n int, edges [][]int, source int, destination int, target int) (ans [][]int) {

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	n := Deserialize[int](ReadLine(stdin))
	edges := Deserialize[[][]int](ReadLine(stdin))
	source := Deserialize[int](ReadLine(stdin))
	destination := Deserialize[int](ReadLine(stdin))
	target := Deserialize[int](ReadLine(stdin))
	ans := modifiedGraphEdges(n, edges, source, destination, target)

	fmt.Println("\noutput:", Serialize(ans))
}
