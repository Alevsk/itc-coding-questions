# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution:
    def mergeTwoLists(self, l1, l2):
        """
        :type l1: ListNode
        :type l2: ListNode
        :rtype: ListNode
        """
        result = ListNode(0)
        result_iter = result
        while l1 and l2:
            a = l1.val
            b = l2.val
            next_node = None
            if a <= b:
                next_node = ListNode(a)
                l1 = l1.next
            else:
                next_node = ListNode(b)
                l2 = l2.next

            result_iter.next = next_node
            result_iter = result_iter.next

        result_iter.next = l1 if l1 is not None else l2
        return result.next
