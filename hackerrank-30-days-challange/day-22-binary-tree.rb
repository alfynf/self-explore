# Day 22: Binary Search Tree

# The height of a binary search tree is the number of edges between the tree's root and its furthest leaf. 
# You are given a pointer, root, pointing to the root of a binary search tree. 
# Complete the getHeight function provided in your editor so that it returns the height of the binary search tree.

class Node
    attr_accessor :left,:right,:data
  def initialize data
      @left=@right=nil
      @data=data
    end
end
class Solution
    def insert (root,data)
        if root==nil
            return Node.new(data)
        else
            if data<=root.data
                cur=self.insert(root.left,data)
                root.left=cur
            else
                cur=self.insert(root.right,data)
                root.right=cur
            end
        end
        return root
    end

    def getHeight(root)
      #Write your code here
        if root == nil
            return 0
        end
        if root != nil && root.left == nil && root.right == nil
            return 0
        end
        
        len_right = self.getHeight(root.left)
        len_left = self.getHeight(root.right)
        
        if len_right > len_left
            len = 1 + len_right
        else
            len = 1 + len_left
        end
        return len
    end
end

myTree=Solution.new
root=nil
T=gets.to_i
for i in 1..T
    data=gets.to_i
    root=myTree.insert(root,data)
end
height=myTree.getHeight(root)
print height