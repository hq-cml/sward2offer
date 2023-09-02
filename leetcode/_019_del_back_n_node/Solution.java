package org.example;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode() {}
 *     ListNode(int val) { this.val = val; }
 *     ListNode(int val, ListNode next) { this.val = val; this.next = next; }
 * }
 */

class ListNode {
     int val;
    ListNode next;
    ListNode() {}
    ListNode(int val) { this.val = val; }
    ListNode(int val, ListNode next) { this.val = val; this.next = next; }
 }

public class Solution {
    public ListNode removeNthFromEnd(ListNode head, int n) {
        if (head==null) {
            return null;
        }
        if (n <= 0 ) {
            return null;
        }

        n = n+1;
        ListNode p1 = head;
        ListNode p2 = head;
        while (n>0 && p2 != null) {
            p2 = p2.next;
            n--;
        }

        if (n == 0) {
            while (p2!=null) {
                p1 = p1.next;
                p2 = p2.next;
            }
            p1.next = p1.next.next;
            return head;
        } else if (n == 1) {
            return head.next;
        } else {
            return null;
        }
    }

    public static void main(String[] args) {
        Solution s = new Solution();
    }
}
