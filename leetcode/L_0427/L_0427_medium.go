package L_0427

import "fmt"

// https://leetcode-cn.com/problems/construct-quad-tree/

// 427. 建立四叉树

// 给你一个 n * n 矩阵 grid ，矩阵由若干 0 和 1 组成。请你用四叉树表示该矩阵 grid 。

// 你需要返回能表示矩阵的 四叉树 的根结点。

// 注意，当 isLeaf 为 False 时，你可以把 True 或者 False 赋值给节点，两种值都会被判题机制 接受 。

// 四叉树数据结构中，每个内部节点只有四个子节点。此外，每个节点都有两个属性：

//     val：储存叶子结点所代表的区域的值。1 对应 True，0 对应 False；
//     isLeaf: 当这个节点是一个叶子结点时为 True，如果它有 4 个子节点则为 False 。

// class Node {
//     public boolean val;
//     public boolean isLeaf;
//     public Node topLeft;
//     public Node topRight;
//     public Node bottomLeft;
//     public Node bottomRight;
// }

// 我们可以按以下步骤为二维区域构建四叉树：

//     如果当前网格的值相同（即，全为 0 或者全为 1），将 isLeaf 设为 True ，将 val 设为网格相应的值，并将四个子节点都设为 Null 然后停止。
//     如果当前网格的值不同，将 isLeaf 设为 False， 将 val 设为任意值，然后如下图所示，将当前网格划分为四个子网格。
//     使用适当的子网格递归每个子节点。

// 四叉树格式：

// 输出为使用层序遍历后四叉树的序列化形式，其中 null 表示路径终止符，其下面不存在节点。

// 它与二叉树的序列化非常相似。唯一的区别是节点以列表形式表示 [isLeaf, val] 。

// 如果 isLeaf 或者 val 的值为 True ，则表示它在列表 [isLeaf, val] 中的值为 1 ；如果 isLeaf 或者 val 的值为 False ，则表示值为 0 。

type Node struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *Node
	TopRight    *Node
	BottomLeft  *Node
	BottomRight *Node
}

func gridValue(grid [][]int, n int) int {
	v := grid[0][0]
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] != v {
				return -1
			}
		}
	}
	return v
}

func construct(grid [][]int) *Node {
	n := len(grid)
	if n == 1 {
		return &Node{
			Val:    grid[0][0] == 1,
			IsLeaf: true,
		}
	}
	curNode := &Node{}
	val := gridValue(grid, n)
	if val != -1 {
		curNode.Val = val == 1
		curNode.IsLeaf = true
		return curNode
	}
	halfN := n / 2
	var tempGrid [][]int

	// TopLeft
	tempGrid = make([][]int, halfN)
	for i := 0; i < halfN; i++ {
		tempGrid[i] = grid[i][:halfN]
	}
	curNode.TopLeft = construct(tempGrid)

	// TopRight
	tempGrid = make([][]int, halfN)
	for i := 0; i < halfN; i++ {
		tempGrid[i] = grid[i][halfN:]
	}
	curNode.TopRight = construct(tempGrid)

	// BottomLeft
	tempGrid = make([][]int, halfN)
	for i := halfN; i < n; i++ {
		tempGrid[i-halfN] = grid[i][:halfN]
	}
	curNode.BottomLeft = construct(tempGrid)

	// BottomRight
	tempGrid = make([][]int, halfN)
	for i := halfN; i < n; i++ {
		tempGrid[i-halfN] = grid[i][halfN:]
	}
	curNode.BottomRight = construct(tempGrid)

	return curNode
}

func Test427() {
	n := construct([][]int{{0, 1}, {1, 0}})
	fmt.Println(n) // [[0,1],[1,0],[1,1],[1,1],[1,0]
}
