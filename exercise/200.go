package exercise

// func gridDfs(grid [][]byte,i,j int){
//     row,col := len(grid),len(grid[0])
//     grid[i][j] = '0'
//     if i-1 >=0 && grid[i-1][j] == '1'{
//         gridDfs(grid,i-1,j)
//     }
//     if i+1 < row && grid[i+1][j] == '1'{
//         gridDfs(grid,i+1,j)
//     }
//     if j-1 >=0 && grid[i][j-1] == '1'{
//         gridDfs(grid,i,j-1)
//     }
//     if j+1 < col && grid[i][j+1] == '1'{
//         gridDfs(grid,i,j+1)
//     }
// }

// // 岛屿数量，深度搜索实现
// func numIslands(grid [][]byte) int {
//     if len(grid) == 0{
//         return 0
//     }
//     ans := 0
//     row,col := len(grid),len(grid[0])
//     for i:=0;i<row;i++{
//         for j:=0;j<col;j++{
// 			// 按行，按列进行遍历，当遍历到的元素是“1”时，
// 			// 进行深度搜索，将上、下、左、右所有元素都标记为0
// 			// 同时进行累数加1
//             if grid[i][j] == '1'{
//                 ans++
//                 gridDfs(grid,i,j)
//             }
//         }
//     }
//     return ans
// }