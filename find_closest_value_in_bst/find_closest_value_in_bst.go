package main

type BST struct {
	Value int

	Left  *BST
	Right *BST
}

// Recursive Solutiony1j
// Time: O(log n); Space: O(log n)
func (tree *BST) FindClosestValue(target int) int {
	return tree.FindClosestValueForNode(target, 1)
}

func (tree *BST) FindClosestValueForNode(target, closestSeen int) int {
	diffForNode := abs(tree.Value, target)
	closestDiff := abs(closestSeen, target)

	if diffForNode < closestDiff {
		closestSeen = tree.Value
	}

	if tree.Value < target && tree.Right != nil {
		return tree.Right.FindClosestValueForNode(target, closestSeen)
	} else if tree.Value > target && tree.Left != nil {
		return tree.Left.FindClosestValueForNode(target, closestSeen)
	}

	return closestSeen
}

func abs(val1, val2 int) int {
	if val2 > val1 {
		return val2 - val1
	}

	return val1 - val2
}

// Iterative Solution
// Time: O(log(n)) ; Space: O(1)
func (tree *BST) FindClosestValue(target int) int {
	curNode := tree
	closestSeen := tree.Value

	for curNode != nil {
		curDiff := abs(target, curNode.Value)
		closestDiff := abs(target, closestSeen)

		if curDiff < closestDiff {
			closestSeen = curNode.Value
		}

		if curNode.Value < target {
			curNode = curNode.Right
		} else if curNode.Value > target {
			curNode = curNode.Left
		} else {
			break
		}
	}

	return closestSeen
}

func abs(val1, val2 int) int {
	if val2 > val1 {
		return val2 - val1
	}

	return val1 - val2
}
