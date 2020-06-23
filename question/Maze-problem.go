package question

import "fmt"

var m = make(map[string]bool)

// 未通过
func MazeProblem(maze [][]int, i, j int) bool {
	defer func() {
		m[fmt.Sprintf("%d@%d", i, j)] = true
	}()

	if m[fmt.Sprintf("%d@%d", i, j)] {
		return false
	}

	// 如果已经在终点, 那么问题以解决
	if i == len(maze)-1 && j == len(maze[0])-1 {
		fmt.Printf("(%d,%d)\n", i, j)
		return true
	}

	var up, down, left, right bool
	// 下
	if i < len(maze)-1 && maze[i+1][j] == 0 {
		down = MazeProblem(maze, i+1, j)
	}
	// 右
	if !down && j < len(maze[0])-1 && maze[i][j+1] == 0 {
		right = MazeProblem(maze, i, j+1)
	}
	// 上
	if !right && i > 0 && maze[i-1][j] == 0 {
		up = MazeProblem(maze, i-1, j)
	}
	// 左
	if !right && j > 0 && maze[i][j-1] == 0 {
		left = MazeProblem(maze, i, j+1)
	}

	// 如果能在该位置走通, 那么返回路径
	if up || down || left || right {
		fmt.Printf("(%d,%d)\n", i, j)
		return true
	}
	return false
}

//func main() {
//	var row, column int
//	_, err := fmt.Scanln(&row, &column)
//	if err != nil {
//		panic(err)
//	}
//
//	maze := make([][]int, row)
//	for i := 0; i < row; i++ {
//		for j := 0; j < column; j++ {
//			var info int
//			if _, err := fmt.Scanf("%d", &info); err != nil {
//				panic(err)
//			}
//			maze[i] = append(maze[i], info)
//		}
//	}
//
//	pass, path := question.MazeProblem(maze, 0, 0)
//	if pass {
//		for _, position := range path {
//			fmt.Printf("%s\n", position)
//		}
//	}
//}
