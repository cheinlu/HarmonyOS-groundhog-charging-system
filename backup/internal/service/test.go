package main

import (
	"fmt"
	"sync"
	"time"
)

func init() {
	fmt.Println("========")
}

func main() {
	// 生成两个长度为 Length 的随机矩阵
	const Length = 100
	matrix1 := generateMatrix(Length)
	matrix2 := generateMatrix(Length)

	// 计算结果矩阵并打印执行时间
	startTime := time.Now()
	multiply(matrix1, matrix2)
	// resultMatrix := multiply(matrix1, matrix2)
	endTime := time.Now()
	elapsedTime := endTime.Sub(startTime).Milliseconds()
	fmt.Printf("矩阵相乘用时：%d 毫秒\n", elapsedTime)

	// atomic.AddInt32()

	// 输出结果矩阵
	// for _, row := range resultMatrix {
	// 	for _, value := range row {
	// 		fmt.Printf("%d ", value)
	// 	}
	// 	fmt.Println()
	// }
}

// 生成一个长度为 length 的随机矩阵
func generateMatrix(length int) [][]int {
	matrix := make([][]int, length)
	for i := 0; i < length; i++ {
		row := make([]int, length)
		for j := 0; j < length; j++ {
			row[j] = j
		}
		matrix[i] = row
	}
	return matrix
}

// 并发计算矩阵相乘
func multiply(matrix1 [][]int, matrix2 [][]int) [][]int {
	// 检查矩阵是否可以相乘
	if len(matrix1[0]) != len(matrix2) {
		panic("矩阵不能相乘！")
	}

	// 创建结果矩阵
	resultMatrix := make([][]int, len(matrix1))
	for i := 0; i < len(matrix1); i++ {
		row := make([]int, len(matrix2[0]))
		resultMatrix[i] = row
	}

	// 创建锁和等待组
	var lock sync.Mutex
	var waitGroup sync.WaitGroup

	// 并发计算结果矩阵的每个元素
	for i := 0; i < len(resultMatrix); i++ {
		for j := 0; j < len(resultMatrix[0]); j++ {
			waitGroup.Add(1)
			go func(row int, col int) {
				defer waitGroup.Done()
				// 计算结果矩阵的对应元素
				sum := 0
				for k := 0; k < len(matrix1[0]); k++ {
					sum += matrix1[row][k] * matrix2[k][col]
				}
				// 将结果矩阵的对应元素写入
				lock.Lock()
				resultMatrix[row][col] = sum
				lock.Unlock()
			}(i, j)
		}
	}

	// 等待所有任务完成
	waitGroup.Wait()

	return resultMatrix
}
