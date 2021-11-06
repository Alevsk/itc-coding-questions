/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode() : val(0), next(nullptr) {}
 *     ListNode(int x) : val(x), next(nullptr) {}
 *     ListNode(int x, ListNode *next) : val(x), next(next) {}
 * };
 */
class Solution {
public:
    ListNode* addTwoNumbers(ListNode* l1, ListNode* l2) {
        ListNode *r = new ListNode();
        ListNode *p;
        p = r;
        int l = 0;        
        while(true)
        {
            int v1 = l1 == nullptr ? 0 : l1->val;
            int v2 = l2 == nullptr ? 0 : l2->val;
            int s = l + v1 + v2;
            p->val = s % 10;
            l = s > 9 ? 1 : 0;
            l1 = l1 == nullptr ? nullptr : l1->next;
            l2 = l2 == nullptr ? nullptr : l2->next;
            if(l1 != nullptr || l2 != nullptr || l != 0)
            {
                p->next = new ListNode();
                p = p->next;
            }
            else
            {
                break;
            }
        }
        return r;
    }
};