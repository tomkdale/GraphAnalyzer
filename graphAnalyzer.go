// Golang Graph Analyzer
//Tom Dale
//
//Graph Analyzer takes in adjacency matrix and
//Checks to see if it is connected
//finds a euler cycle or a euler trail
//or notifies that no such paths exist

package main

import (
	"fmt"
)

//function called by findEulerPath, returns true if edge is a bridge, false if not
func isBridge(matrix [][]int, numVert, x, y int) bool {
	//edge is a bridge if the number of connected vertices is decreased when edge is removed

	//first check current number of connected edges
	before := numConnected(matrix, numVert)
	//now delte edge and check connected edges again
	matrix[x][y] = 0
	matrix[y][x] = 0
	after := numConnected(matrix, numVert) //check new number of connected edges
	matrix[x][y] = 1                       //add the edge back in because were dealing with 2D slices which are passed by reference in GoLang
	matrix[y][x] = 1
	if after < before { //if deleting that edge changed the number of connected points then it was a bridge
		return true
	}
	return false
}

//finds euler path given a proper vertex to start at
//function will print a cycle if one exists, regardless of where it starts, if only a path exists then first will be the proper vertex to start at
func findEulerPath(matrix [][]int, numVert, first int) {
	//start at first vertex passed into function (will be one of two odd values)
	fmt.Printf("%d", first+1)
	x := first //x and y are coordinates on the matrix
	y := 0
	storeBridge := 0 //this integer stores a bridge, when possible take any other edge except the bridge
	for true {       // keep chugging through matrix until all paths have been used
		if matrix[x][y] != 0 { //for each edge check if it is a bridge if so save it, if not take it
			if isBridge(matrix, numVert, x, y) {
				storeBridge = y
				y++
			} else { //if not bridge take travel across the edge then delete it from matrix list
				fmt.Printf("-> %d", y+1)
				matrix[x][y] = 0 // delete edge just taken
				matrix[y][x] = 0
				x = y
				y = 0
				storeBridge = -1
			}
		} else { //increment matrix iterator to next possible edge and repeat
			y++
		}

		//if weve gotten through all of potential edges of a vertex either take the saved bridge, or weve complete the path so get out of function
		if y == numVert { //if
			if storeBridge == -1 {
				return
			}
			fmt.Printf("-> %d", storeBridge+1) //move across bridge then delete edge from matrix
			matrix[x][storeBridge] = 0
			matrix[storeBridge][x] = 0
			x = storeBridge
			y = 0
			storeBridge = -1
		}

	}
}

//checks if euler path exists, only called when no euler cycle exists
func checkEulerPath(matrix [][]int, numVert int) bool {
	listOfDegrees := make([]int, numVert) //list the degrees of each edge
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
			//keep track of how many Odd degree vertices there are and label the first two values to be used by find EulerPath()
			if numOddDeg == 1 {
				firstOdd = i
			}
		}
	}
	//if there are in fact only too odd degreed vertices that are of degree 3 then pass that info to findEulerPath()
	if numOddDeg == 2 {
		fmt.Println("Euler path exists")
		findEulerPath(matrix, numVert, firstOdd) //no that euler path exists, find one
		return true
	}
	return false
}

//check to see if there is an euler cycle. Eulers cycles are easy, if there are all even degree vertices then one exists
func checkEulerCycle(matrix [][]int, numVert int) bool {
	for i := 0; i < numVert; i++ { //check each vertex, if one has odd degree return false
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
	return true //no odd degree vertex exist so return true
}

//recursive function checks all vertices in matrix to see if graph is connected
func connectedRecursive(matrix [][]int, numVert int, iter int, connected []int) []int {
	//iter is the iterator showing which column function is checking
	connected[iter] = 1 // add self to list of connected vertices
	for i := 0; i < numVert; i++ {
		if matrix[iter][i] != 0 && connected[i] == 0 { //if value hasnt been visited call recursive function to check it and all its child vertices
			connected = connectedRecursive(matrix, numVert, i, connected)
		}
	}
	return connected // return full list of connected verteces
}

//check connected returns number of connected vertices in matrix, if numVert = numConnected then graph is fully connected
func numConnected(matrix [][]int, numVert int) int {
	connected := make([]int, numVert)
	//first run through matrix to find first connected vertex, vertex 1 is not necissarily connected to anything else
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
	//no get list of connected points using connectRecursive function
	connected = connectedRecursive(matrix, numVert, first, connected)
	count := 0
	//add up all vertices present in list
	for i := 0; i < numVert; i++ {
		if connected[i] != 0 {
			count++
		}
	}
	return count
}

//Returns true if matrix is not symmetrical across the diagonal
func checkBadMatrix(matrix [][]int, numVert int) bool {
	for i := 0; i < numVert; i++ {
		for j := i + 1; j < numVert; j++ {
			if matrix[i][j] != matrix[j][i] { //just check if values are mirrored across diagonal
				return true
			}
		}
	}
	return false
}

//function prints matrix passed in
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

//function reads adjacency matrix from user input
func readAdjacencyMatrix(matrix [][]int, numVert int) [][]int {

	for i := 0; i < numVert; i++ {
		//fmt.Println("Enter list of 1's and 0's representing the vertices adjacent to vertex #", i+1, " (space separated)")
		for j := 0; j < numVert; j++ {
			fmt.Scan(&matrix[i][j])
		}
	}
	for i := 0; i < numVert; i++ { //set diagonal to be 0, since no vertex should be self connected
		matrix[i][i] = 0
	}
	return matrix
}

//main function takes in adjacency matrix then says if its connected, then gives and euler cycle or path
func main() {
	fmt.Println("Starting graph analyzer...")
	fmt.Println("Enter number of vertices:")
	var numVert int //number of vertices in adjacency matrix
	fmt.Scan(&numVert)

	//Create adjacency matrix
	matrix := make([][]int, numVert)
	for i := range matrix {
		matrix[i] = make([]int, numVert)
	}
	fmt.Println("Enter adjancy matrix, 1 = adjacent, 0 = not adjacent")
	//read from stdin the values for the matrix
	readAdjacencyMatrix(matrix, numVert)

	//cleanly print out matrix (will have corrected all digonal values to be zero)
	printMatrix(matrix, numVert)
	//check to see if matrix was not diagonally symmetrical, if so stop program
	if checkBadMatrix(matrix, numVert) {
		fmt.Println("Bad adjacency matrix, not diagonally symmetrical. Please restart and try again.")
		return
	}
	//check if graph is connected
	if numConnected(matrix, numVert) == numVert {
		fmt.Println("Graph is connected.")
		//check if euler cycle exists, if so find one
		if checkEulerCycle(matrix, numVert) {
			fmt.Println("Euler cycle exists!")
			findEulerPath(matrix, numVert, 0)
		} else {
			//check if euler path exists if so find one
			fmt.Println("Graph has no euler Cycle")
			if !checkEulerPath(matrix, numVert) {
				fmt.Println("and graph has no euler path.")
			}
		}
	} else {
		fmt.Println("Graph is not connected!")
	}

	return
}
