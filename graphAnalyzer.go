package main

import (
	"fmt"
) //Graph Analyzer takes in adjacency matrix and
//Checks to see if it is connected
//finds a euler cycle or a euler trail

func isBridge(matrix [][]int, numVert, x, y int) bool { //function called by findEulerPath, returns true if edge is a bridge, false if not
	//edge is a bridge if the number of connected vertices is decreased when edge is removed

	//first check current number of connected edges
	before := numConnected(matrix, numVert)
	//now delte edge and check connected edges again
	matrix[x][y] = 0
	matrix[y][x] = 0
	after := numConnected(matrix, numVert)
	fmt.Printf("b %d , a %d \n", before, after)
	if after < before {
		return true
	}
	return false
}

//finds euler path given a proper vertex to start at
func findEulerPath(matrix [][]int, numVert, first int) {
	//start at first
	fmt.Printf("%d", first+1)
	x := first //x and y are coordinates on the matrix
	y := 0
	storeBridge := 0
	for true { // keep chugging through matrix until all paths have been used
		if x == 3 {
			fmt.Println(x, " ", y)
			printMatrix(matrix, numVert)
		}
		if matrix[x][y] != 0 { //for each edge check if it is a bridge if so save it, if not take it
			if isBridge(matrix, numVert, x, y) {
				storeBridge = y
				fmt.Printf("Bridge at %d %d ", x, y)
				y++
			} else {
				fmt.Printf("-> %d", y+1)
				matrix[x][y] = 0 // delete edge just taken
				matrix[y][x] = 0
				x = y
				y = 0
				storeBridge = -1
			}
		} else {
			y++
		}

		if y == numVert { //if
			if storeBridge == -1 {
				return
			}
			fmt.Printf("-> %d", storeBridge+1)
			matrix[x][storeBridge] = 0
			matrix[storeBridge][x] = 0
			x = storeBridge
			y = 0
			storeBridge = -1
		}

	}
}

func checkEulerPath(matrix [][]int, numVert int) bool {
	listOfDegrees := make([]int, numVert)
	for i := 0; i < numVert; i++ {
		for j := 0; j < numVert; j++ {
			if matrix[i][j] != 0 {
				listOfDegrees[i]++
			}
		}
	}
	//now search through list of degrees and find if there is all even values except 2 3 degree values
	numOddDeg := 0
	firstOdd := -1 //initalize to -1 since this negative values will never be used

	for i := 0; i < numVert; i++ {
		if listOfDegrees[i]%2 == 1 {
			numOddDeg++
			//keep track of how many Odd degree verteces there are and label the first two values to be used by find EulerPath()
			if numOddDeg == 1 {
				firstOdd = i
			}
		}
	}
	//if there are in fact only too odd degreed verteces that are of degree 3 then pass that info to findEulerPath()
	if numOddDeg == 2 {
		findEulerPath(matrix, numVert, firstOdd)
		return true
	}
	return false
}

func checkEulerCycle(matrix [][]int, numVert int) bool {
	for i := 0; i < numVert; i++ {
		degree := 0
		for j := 0; j < numVert; j++ {
			if matrix[i][j] != 0 { //add 1 to degree for every connected vertex
				degree++
			}
		}
		if degree%2 == 1 {
			return false
		}
	}
	return true
}

func printEulerCycle(matrix [][]int, numVert int) {
	i := 0
	fmt.Printf("%d", i)
	for j := 0; j < numVert; {
		if matrix[i][j] != 0 {
			fmt.Printf("->%d ", j)
			matrix[i][j] = 0
			matrix[j][i] = 0
			i = j
			j = 0
		} else {
			j++
		}
	}

}

//recursive function checks all verteces in matrix to see if graph is connected
func connectedRecursive(matrix [][]int, numVert int, iter int, connected []int) []int {
	//iter is the iterator showing which column function is checking
	connected[iter] = 1
	for i := 0; i < numVert; i++ {
		if matrix[iter][i] != 0 && connected[i] == 0 {
			connected = connectedRecursive(matrix, numVert, i, connected)
		}
	}
	return connected
}

//check connected returns number of connected verteces in matrix
func numConnected(matrix [][]int, numVert int) int {
	connected := make([]int, numVert)
	//first run through matrix to find first connected vertex
	first := 0
	for i := 0; i < numVert; i++ {
		for j := 0; j < numVert; j++ {
			if matrix[i][j] != 0 {
				first = i
				break
			}
		}
		if first != 0 {
			break
		}
	}
	connected = connectedRecursive(matrix, numVert, first, connected)
	count := 0
	for i := 0; i < numVert; i++ {
		if connected[i] != 0 {
			count++
		}
	}
	return count
}

func checkBadMatrix(matrix [][]int, numVert int) bool {
	for i := 0; i < numVert; i++ {
		for j := i + 1; j < numVert; j++ {
			if matrix[i][j] != matrix[j][i] {
				return true
			}
		}
	}
	return false
}

func printMatrix(matrix [][]int, numVert int) {
	fmt.Println("Printing adjacency list")
	for i := 0; i < numVert; i++ {
		for j := 0; j < numVert; j++ {
			fmt.Printf("%d ", matrix[i][j])
		}
		fmt.Println()
	}
	return
}

func readAdjacencyMatrix(matrix [][]int, numVert int) [][]int {

	for i := 0; i < numVert; i++ {
		//fmt.Println("Enter list of 1's and 0's representing the verteces adjacent to vertex #", i+1, " (space separated)")
		for j := 0; j < numVert; j++ {
			fmt.Scan(&matrix[i][j])
		}
	}
	for i := 0; i < numVert; i++ { //set diagonal to be 0, since no vertex should be self connected
		matrix[i][i] = 0
	}
	return matrix
}
func main() {
	fmt.Println("Starting graph analyzer...")
	fmt.Println("Enter number of verteces:")
	var numVert int
	fmt.Scan(&numVert)

	//Take in adjacency matrix
	matrix := make([][]int, numVert)
	for i := range matrix {
		matrix[i] = make([]int, numVert)
	}
	fmt.Println("Enter adjancy matrix, 1 = adjacent, 0 = not adjacent")
	readAdjacencyMatrix(matrix, numVert)

	printMatrix(matrix, numVert)
	if checkBadMatrix(matrix, numVert) {
		fmt.Println("Bad adjacency matrix, not diagonally symmetrical. Please restart and try again.")
		return
	}
	//check if graph is connected
	if numConnected(matrix, numVert) == numVert {
		fmt.Println("Graph is connected.")
		if checkEulerCycle(matrix, numVert) {
			fmt.Println("Euler cycle exists!")
			printEulerCycle(matrix, numVert)
		} else {
			fmt.Println("Graph has no euler Cycle")
			if checkEulerPath(matrix, numVert) {
				fmt.Println("Euler path exists")
			} else {
				fmt.Println("and graph has no euler path.")
			}
		}
	} else {
		fmt.Println("Graph is not connected!")
	}

	return
}
