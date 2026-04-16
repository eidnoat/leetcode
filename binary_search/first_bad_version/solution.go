/** 
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad 
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
    l, r, ret := 1, n, 0
	for l <= r {
		m := (l+r)/2
		if isBadVersion(m) {
			if m-1 >= l && isBadVersion(m-1) {
				r = m-1
			} else {
				ret = m
				break
			}
		} else {
			l = m+1
		}
	}

	return ret
}

/*
	执行用时分布：0ms，击败 100.00%
	消耗内存分布：3.67MB，击败 35.35%
*/
