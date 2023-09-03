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
    public List<List<Integer>> permute(int[] nums) {
        List<List<Integer>> ret = new ArrayList<>();
        List<Integer> org = new ArrayList<>();
        List<Integer> curr = new ArrayList<>();

        for (int v : nums) {
            org.add(v);
        }
        help(org, curr, ret);
        return ret;
    }

    private void help(List<Integer> nums, List<Integer> curr, List<List<Integer>> ret) {
        if (nums.size() == 0 ) {
            ret.add(curr);
            return;
        }


        for (int i=0; i<nums.size(); i++ ) {
            List<Integer> copyNums = new ArrayList<>(nums);
            List<Integer> copyCurr = new ArrayList<>(curr);
            copyCurr.add((nums.get(i)));
            copyNums.remove(i);
            help(copyNums, copyCurr, ret);
        }
    }

    public static void main(String[] args) {
        Solution s = new Solution();
        int[] arr = {1,2,3};
        //s.combinationSum(arr, 7);
        //System.out.println(s.combinationSum(3));
    }
}
