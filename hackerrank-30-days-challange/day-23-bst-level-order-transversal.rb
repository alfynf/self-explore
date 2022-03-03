# Day 23: BST Level-Order Transversal

# Task: A level-order traversal, also known as a breadth-first search, visits each level of a tree's nodes from left to right, 
# top to bottom. You are given a pointer, root, pointing to the root of a binary search tree. 
# Complete the levelOrder function provided in your editor so that it prints the level-order traversal of the binary search tree.

# Hint: You'll find a queue helpful in completing this challenge. 

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

    def levelOrder(root)
        #Write your code here
        queue = Array[root]
        
        while queue.length != 0
            cur = queue[0]
            queue = queue.slice(1,queue.length)
            print "#{cur.data} "
        
            if cur.left != nil
                queue.append(cur.left)
            end
            if cur.right != nil
                queue.append(cur.right)
            end
        end
    
    end
end

myTree=Solution.new
root=nil
T=gets.to_i
for i in 1..T
    data=gets.to_i
    root=myTree.insert(root,data)
end
myTree.levelOrder(root)
