/*
Given a 32-bit signed integer, reverse digits of an integer.

Example 1:

Input: 123
Output:  321
Example 2:

Input: -123
Output: -321
Example 3:

Input: 120
Output: 21
Note:
Assume we are dealing with an environment which could only hold integers within the 32-bit signed integer range. For the purpose of this problem, assume that your function returns 0 when the reversed integer overflows.
*/

func reverse(x int) int {
    ret := 0
    for ; x!=0; {
        ret = ret*10 + x%10
        x = x/10
    }

    if ret > 0x7fffffff || ret < -2147483648{
        return 0
    }
    
    return ret
}
