package org.example;

import java.util.*;

class ListNode {
     int val;
    ListNode next;
    ListNode() {}
    ListNode(int val) { this.val = val; }
    ListNode(int val, ListNode next) { this.val = val; this.next = next; }
 }

public class Solution {
    public boolean canJump(int[] nums) {
        return help(nums, 0);
    }

    private boolean help(int[] nums, int currIdx) {
        if (currIdx == nums.length-1) {
            return true;
        }

        for (int i=1; i<= nums[currIdx]; i++) {
            if (help(nums, currIdx + i)) {
                return true;
            }
        }
        return false;
    }

    public static void main(String[] args) {
        Solution s = new Solution();
        int[] arr = {3,2,1,0,4};
        s.canJump(arr);
        //System.out.println(s.combinationSum(3));
    }
}
