# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution:
    def addTwoNumbers(self, l1, l2):
        """
        :type l1: ListNode
        :type l2: ListNode
        :rtype: ListNode
        """
        result = ListNode(0)
        result_iter = result
        carry = False

        l1_iter = l1
        l2_iter = l2

        while l1_iter or l2_iter:
            a = l1_iter.val if l1_iter else 0
            b = l2_iter.val if l2_iter else 0

            operand_sum = a + b + int(carry)
            carry = operand_sum > 9

            result_iter.next = ListNode(operand_sum % 10)
            result_iter = result_iter.next

            l1_iter = l1_iter.next if l1_iter else None
            l2_iter = l2_iter.next if l2_iter else None

        if carry:
            result_iter.next = ListNode(1)
            result_iter = result_iter.next

        return result.next
