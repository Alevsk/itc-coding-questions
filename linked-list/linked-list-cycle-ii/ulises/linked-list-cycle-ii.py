# Definition for singly-linked list.
class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution:
    def detectCycle(self, head: ListNode) -> ListNode:
        k = 1
        s = head
        while s != None and k <= 10 ** 4:
            s = s.next
            k += 1
        if s == None:
            return None

        max_index = (10 ** 4) - 1
        min_index = 0
        while True:
            m = (max_index + min_index) // 2
            s = head
            for i in range(m):
                s = s.next
            t = s.next
            for i in range(10 ** 4):
                if t == s:
                    break
                t = t.next
            if i < 10 ** 4:
                if min_index == max_index:
                    return s
                max_index = m
            else:
                if min_index == max_index:
                    break
                min_index = m + 1
        return None

h = ListNode(3)
h.next = ListNode(2)
h.next.next = ListNode(0)
h.next.next.next = ListNode(-4)
h.next.next.next.next = h.next

s = Solution()
s.detectCycle(h)