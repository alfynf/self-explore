# Day 34: More Linked List

# Task: A Node class is provided for you in the editor. A Node object has an integer data field, data, and a Node instance pointer,
# next, pointing to another node (i.e.: the next node in a list).
# A removeDuplicates function is declared in your editor, which takes a pointer to the head node of a linked list as a parameter. 
# Complete removeDuplicates so that it deletes any duplicate nodes from the list and returns the head of the updated list.

# Note: The head pointer may be null, indicating that the list is empty. Be sure to reset your pointer when performing deletions to 
# avoid breaking the list.

class Node

    attr_accessor :data,:next

  def initialize data
        @data = data
        @next = nil
    end
end

class Solution

    def removeDuplicates(head)
      #write your code here
      cur = head
      while cur != nil && cur.next != nil
        while cur.next != nil && cur.data == cur.next.data
            cur.next = cur.next.next
        end
            cur = cur.next
      end
        return head
    end

    def insert(head,value)
        p=Node.new(value)
        if head==nil
             head=p
         elsif head.next==nil
             head.next=p
         else  
             current = head
         while current.next != nil
             current = current.next
       end 
       current.next = p
      end
      return head   
   end 	

	def display(head)
       current = head
       while current 
          print current.data," "
          current = current.next
       end
   end
end
       
mylist= Solution.new
head=nil
T=gets.to_i
for i in 1..T
    data=gets.to_i
    head=mylist.insert(head,data)
end
head=mylist.removeDuplicates(head)
mylist.display(head)