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
    public List<List<Integer>> combinationSum(int[] candidates, int target) {
        // 排序
        Arrays.sort(candidates);

        List<List<Integer>> ret = new ArrayList<>();
        List<Integer> curr = new ArrayList<>();
        help(candidates, 0, target, curr, ret);
        return ret;
    }

    private void help(int[] candidates,int begIdx, int target, List<Integer> curr, List<List<Integer>> ret) {
        for (int i = 0; i< candidates.length; i++) {
            if (i < begIdx) {
                continue;
            }
            List<Integer> newCopy = new ArrayList<>(curr);
            if (target == candidates[i]) {
                // 找到一个符合预期的结果
                newCopy.add(candidates[i]);
                ret.add(newCopy);
            } else if (target > candidates[i]) {
                newCopy.add(candidates[i]);
                help(candidates, i, target - candidates[i], newCopy, ret);
            }
        }
    }

    public static void main(String[] args) {
        Solution s = new Solution();
        int[] arr = {2,7,6,3};
        s.combinationSum(arr, 7);
        //System.out.println(s.combinationSum(3));
    }
}
