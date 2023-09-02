package org.example;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

public class Solution {
    public HashMap<Character, List<Character>> cMap;

    public Solution() {
        cMap = new HashMap<Character, List<Character>>();
        ArrayList<Character> l = new ArrayList<Character>();
        l.add('a');
        l.add('b');
        l.add('c');
        cMap.put('2', l);
        l = new ArrayList<>();
        l.add('d');
        l.add('e');
        l.add('f');
        cMap.put('3', l);
    }

    public List<String> letterCombinations(String digits) {
        List<String> ret = new ArrayList<String>();
        help(digits, "", ret);
        return ret;
    }

    private void help(String target, String curr, List<String> ret) {
        if (target.length() == 0 ) {
            ret.add(curr);
            System.out.println(curr);
            return;
        }
        Character c = target.charAt(0);
        if (!cMap.containsKey(c)) {
            return;
        }
        List<Character> l = cMap.get(c);
        for (Character ch: l) {
            help(target.substring(1), curr + ch, ret);
        }
    }

    public static void main(String[] args) {
        Solution s = new Solution();
        System.out.println(s.letterCombinations("23"));
    }
}
