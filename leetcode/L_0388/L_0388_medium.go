package L_0388

import (
	"fmt"
	"strings"
)

// https://leetcode-cn.com/problems/longest-absolute-file-path/

// 388. 文件的最长绝对路径

// 假设有一个同时存储文件和目录的文件系统。下图展示了文件系统的一个示例：

// 这里将 dir 作为根目录中的唯一目录。dir 包含两个子目录 subdir1 和 subdir2 。subdir1 包含文件 file1.ext 和子目录 subsubdir1；subdir2 包含子目录 subsubdir2，该子目录下包含文件 file2.ext 。

// 在文本格式中，如下所示(⟶表示制表符)：

// dir
// ⟶ subdir1
// ⟶ ⟶ file1.ext
// ⟶ ⟶ subsubdir1
// ⟶ subdir2
// ⟶ ⟶ subsubdir2
// ⟶ ⟶ ⟶ file2.ext

// 如果是代码表示，上面的文件系统可以写为 "dir\n\tsubdir1\n\t\tfile1.ext\n\t\tsubsubdir1\n\tsubdir2\n\t\tsubsubdir2\n\t\t\tfile2.ext" 。'\n' 和 '\t' 分别是换行符和制表符。

// 文件系统中的每个文件和文件夹都有一个唯一的 绝对路径 ，即必须打开才能到达文件/目录所在位置的目录顺序，所有路径用 '/' 连接。上面例子中，指向 file2.ext 的 绝对路径 是 "dir/subdir2/subsubdir2/file2.ext" 。每个目录名由字母、数字和/或空格组成，每个文件名遵循 name.extension 的格式，其中 name 和 extension由字母、数字和/或空格组成。

// 给定一个以上述格式表示文件系统的字符串 input ，返回文件系统中 指向 文件 的 最长绝对路径 的长度 。 如果系统中没有文件，返回 0。

// 执行用时：0 ms, 在所有 Go 提交中击败了100.00% 的用户
// 内存消耗：2.2 MB, 在所有 Go 提交中击败了12.77% 的用户

func lengthLongestPath(input string) int {
	if !strings.ContainsRune(input, '.') { // 路径里没 . 说明没有文件，长度即为0
		return 0
	}
	input = strings.TrimLeft(input, "\n")             // 去除前面的换行符
	maxLength := 0                                    // 记录最长路径
	curFolderLen := 0                                 // 当前文件夹名称长度
	childFolders := strings.Builder{}                 // 构建子目录
	for _, line := range strings.Split(input, "\n") { // 按行遍历
		if strings.HasPrefix(line, "\t") { // 以 制表符 开头的 是 子目录/子文件
			childFolders.WriteRune('\n')       // 这里可能产生冗余的前缀换行符
			childFolders.WriteString(line[1:]) // 去掉开头的一个制表符，将其加入当前子目录中
		} else { // 其他的 是当前目录内的子目录或文件
			if childFolders.Len() > 0 { // 结算已经记录的子目录
				// 递归计算最长文件路径（子目录中的最长路径 + 当前目录名长度 + 1  加上一个目录分隔符）
				l := lengthLongestPath(childFolders.String()) + curFolderLen + 1
				if maxLength < l {
					maxLength = l
				}
				childFolders.Reset() // 重置子目录记录
			}
			if strings.ContainsRune(line, '.') { // 当前行若是文件名。则按文件名计算长度
				l := len(line)
				if maxLength < l {
					maxLength = l
				}
			} else { // 重置当前目录的目录名长度
				curFolderLen = len(line)
			}
		}
	}
	// 若遍历完，仍有子目录未计算，则进行递归计算，与上面对子目录的操作相同
	if childFolders.Len() > 0 {
		l := lengthLongestPath(childFolders.String()) + curFolderLen + 1
		if maxLength < l {
			maxLength = l
		}
	}
	// fmt.Printf("%d %s\n", maxLength, input)
	return maxLength
}

func Test388() {
	fmt.Println(lengthLongestPath("dir\n        file.txt"))                                                                     // 16
	fmt.Println(lengthLongestPath("a\n\tb\n\t\tc"))                                                                             // 0
	fmt.Println(lengthLongestPath("dir\n\tsubdir1\n\tsubdir2\n\t\tfile.ext"))                                                   // 20
	fmt.Println(lengthLongestPath("a"))                                                                                         // 0
	fmt.Println(lengthLongestPath("dir\n\tsubdir1\n\t\tfile1.ext\n\t\tsubsubdir1\n\tsubdir2\n\t\tsubsubdir2\n\t\t\tfile2.ext")) // 32
}
