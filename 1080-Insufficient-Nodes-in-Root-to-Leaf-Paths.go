
/*

		  1
	2	  	 	-3

-5		   4

[1,2,-3,-5,nil,4,nil] ,limit=-1

预期结果是
		  1
		  	 	-3

		   4
[1,nil,-3,nil,nil,4,nil] 

题目的意思是 根-叶子 整条路径要么保留，要么删除

*/


func InnersufficientSubset(root *TreeNode, limit int, parent *TreeNode) bool{
	if nil == root{
		return false
	}

    bLeftRmvChild := false
    bRightRmvChild := false
    bLeftRmvChild = InnersufficientSubset(root.Left, limit-root.Val, root)
    bRightRmvChild = InnersufficientSubset(root.Right, limit-root.Val, root)
    if (bLeftRmvChild || bRightRmvChild) && root.Left == root.Right{
        if root == parent.Left {
            parent.Left = nil
        } else {
            parent.Right = nil
        }
        return true
    }

	if root.Val < limit && root.Left == root.Right {
		if root == parent.Left {
			parent.Left = nil
		} else {
			parent.Right = nil
		}

		return true
	}

	return false
}

func sufficientSubset(root *TreeNode, limit int) *TreeNode {
    if root == nil{
        return nil
    }
	InnersufficientSubset(root.Left, limit-root.Val, root)
	InnersufficientSubset(root.Right, limit-root.Val, root)

	if root.Right == root.Left && root.Val < limit{
	    return nil
    }
	return root
}