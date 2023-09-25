package org.example;

import javax.transaction.TransactionRequiredException;
import java.util.*;

class ListNode {
     int val;
    ListNode next;
    ListNode() {}
    ListNode(int val) { this.val = val; }
    ListNode(int val, ListNode next) { this.val = val; this.next = next; }
 }
 class TreeNode {
    int val;
     TreeNode left;
     TreeNode right;
    TreeNode() {}
    TreeNode(int val) { this.val = val; }
    TreeNode(int val, TreeNode left, TreeNode right) {
        this.val = val;
        this.left = left;
         this.right = right;
     }
}

public class Solution {
    public int search(int[] nums, int target) {
        int beg = 0;
        int end = nums.length-1;

        while (beg <= end) {
            int mid = (end-beg)/2 + beg;
            if (nums[mid] == target) {
                return mid;
            } else {
                if (nums[mid] > nums[beg]) {
                    // 前半截正常，旋转出现在后半截
                    if (nums[beg]<=target && mid>0 && nums[mid-1] >= target) {
                        end = mid-1;
                    } else {
                        beg = mid+1;
                    }
                } else {
                    // 后半截正常，旋转出现在前半截
                    if (nums[end]>=target && mid+1 < nums.length && nums[mid+1] <= target) {
                        beg = mid + 1;
                    } else {
                        end = mid-1;
                    }
                }
            }
        }
        return -1;
    }

    public static void main(String[] args) {
        Solution s = new Solution();
        int[] arr = {3,2,1,0,4};
        //s.canJump(arr);
        //System.out.println(s.combinationSum(3));
    }
}
