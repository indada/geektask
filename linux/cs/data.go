package main
import (
	"fmt"
	"bytes"
	"math"
)
const (
	a = iota
	b = iota
	c = iota
	d = iota
)
func main()  {
	s := "hello"
	gg := s[0]
	fmt.Println(gg)
	fmt.Println(a,b,c,d)
	arr := [][]string{[]string{"1","2","3"},[]string{"4","5","6"}}
	fmt.Println(arr)
	// 声明一个数组
	var array = [10]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
	// 声明两个 slice key只能是int
	var aSlice, bSlice []string
	aSlice = array[1:4]
	bSlice = aSlice[1:2]
	fmt.Println(aSlice)
	fmt.Println(bSlice)
	array[2] = "GG" //切片有指针切片会改，数组不会 [...][1]是数组，[]空是切片
	fmt.Println("bSlice",bSlice)
	fmt.Println("aSlice",aSlice)
	ar := make([]string,10)
	ar[1] = "one"
	fmt.Println(ar)
	//map
	mm := make(map[string]int)
	mm["g"] = 5
	mm["gg"] = 55
	mm["ggg"] = 555
	fmt.Println(mm["gg"])
	delete(mm,"gg")
	i := myFunc()
	fmt.Println(i)
	type Human struct {
		name string
		age int
		weight int
	}

	type Student struct {
		Human  // 匿名字段，那么默认 Student 就包含了 Human 的所有字段
		speciality string
		name string
	}
	stu := Student{Human{"xx",18,120},"sii","99"}
	stu.Human.name = "xdx"
	fmt.Println(stu.name,stu)
	j := "abcabab"
	fmt.Println(len(j))
	nu1 := []int{1,3,5,10,12,13,18}
	nu2 := []int{2,5,8,9,15,17,20,22}
	ff := findMedianSortedArrayss(nu1,nu2)
	fmt.Println("fffffff")
	fmt.Println(ff)
	ss:= "assafdsfsdfs"
	qq := longestPalindromes(ss)
	fmt.Println("bigLong:",qq)
	fmt.Println("Z字形变换start")
	wq := convert(ss,3)
	fmt.Println("Z字形变换：",wq)
	wqq := convertys(ss,3)
	fmt.Println("Z字形变换2：",wqq)
	rr := reverses(15956156)
	fmt.Println("整数反转：",rr)
	s1 := "kkjkkkkjr"
	p1 := ".kjk*.*"
	if  isMatch(s1,p1) {
		fmt.Println("正则表达式匹配ok")
	}else {
		fmt.Println("正则表达式匹配no")
	}
}

func myFunc() (i int){
	i = 0
	defer func() {
		if x:= recover();x!=nil {
			fmt.Println(x)
			i = 1
		}
	}()

	if i == 0 {
		panic("nonono")
	}
	return

}


type ListNode struct {
	Val int
	Next *ListNode
}
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{Val: -1}
	curr, carry := dummy, 0
	for l1 != nil || l2 != nil {
		x, y := 0, 0
		if l1 != nil {
			x = l1.Val
		}
		if l2 != nil {
			y = l2.Val
		}

		total := x + y + carry
		curr.Next = &ListNode{Val: total % 10}
		// bug 修复：视频中忘了加上这一步
		curr = curr.Next
		carry = total / 10

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	if carry != 0 {
		curr.Next = &ListNode{Val: carry}
	}
	return dummy.Next
}
//aboojadpjj
func lengthOfLongestSubstring(s string) int {
	var n = len(s) //7
	if n <= 1 {
		return n
	}
	var maxLen = 1
	var left, right, window = 0, 0, make(map[byte]bool)
	for right < n {
		var rightChar = s[right]//a b a
		for window[rightChar] {
			delete(window, s[left]) // win b true left 1
			left++
		}
		if right - left + 1 > maxLen { //maxlen 2,
			maxLen = right - left + 1
		}
		window[rightChar] = true //a true,b true
		right++ //1,2
	}
	return maxLen
}
func lengthOfLongestSubstrings(s string) int {
	var n = len(s)
	if n <= 1 {
		return n
	}
	left := 0
	win := make(map[byte]bool)
	max := 1
	for r := 0; r < n; r++ {
		str := s[r]
		for win[str] {
			delete(win,s[left])
			left++
		}
		if r-left +1 >max {
			max = r-left +1
		}
		win[str] = true
	}
	return max
}

//[1,3,5,10,12,13,18]  [2,5,8,9,15,17,20,22]
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	totalLength := len(nums1) + len(nums2)
	if totalLength%2 == 1 { //只有一个中位数
		midIndex := totalLength/2 //取整 7
		return float64(getKthElement(nums1, nums2, midIndex + 1))
	} else {
		midIndex1, midIndex2 := totalLength/2, totalLength/2 +1
		return float64(getKthElement(nums1, nums2, midIndex1) + getKthElement(nums1, nums2, midIndex2)) / 2.0
	}
	return 0
}
//[1,3,5,10,12,13,18]  [2,5,8,9,15,17,20,22]
func findMedianSortedArrayss(nums1 []int, nums2 []int) float64 {
	totalLength := len(nums1) + len(nums2)
	if totalLength%2 == 1 { //只有一个中位数
		midIndex := totalLength/2 //取整 7
		return float64(getE(nums1, nums2, midIndex +1)) //8
	} else {
		midIndex1, midIndex2 := totalLength/2 + 1, totalLength/2
		n1 := getE(nums1, nums2, midIndex1)
		n2 := getE(nums1, nums2, midIndex2)
		fmt.Println("midIndex1:",midIndex1,"midIndex2:",midIndex2,"n1:",n1,"n2:",n2)
		return float64( n1 + n2) / 2.0
	}
	return 0
}
// 中位数二分查找 k=8
func getE(nums1, nums2 []int, k int) int {
	//k是取剩余几位数
	index1,index2 := 0,0 //两个数组当前的位置
	for {
		if index1 == len(nums1) {
			return nums2[index2+k-1]
		}
		if index2 == len(nums2) {
			return nums1[index1+k-1]
		}
		if k == 1 {
			return min(nums1[index1],nums2[index2])
		}
		pp := k/2
		p1 := min(index1+pp,len(nums1))-1
		p2 := min(index2+pp,len(nums2))-1
		n1,n2 := nums1[p1],nums2[p2]
		if n1<n2 {
			k-= (p1+1-index1)
			index1 = p1+1
		}else {
			k-= (p2+1-index2)
			index2= p2+1
		}
	}
	return 0
}
func minList(x int,y int) int {
	if x > y {
		return y
	}
	return x
}

//k = 5 中位数 二分查找法
func getKthElement(nums1, nums2 []int, k int) int {
	index1, index2 := 0, 0 //1 nums1
	for {
		if index1 == len(nums1) {
			return nums2[index2 + k - 1]
		}
		if index2 == len(nums2) {
			return nums1[index1 + k - 1]
		}
		if k == 1 {
			return min(nums1[index1], nums2[index2])
		}
		half := k/2 //4
		newIndex1 := min(index1 + half, len(nums1)) - 1 // 3 前进k/2格 前进4格，index要-1
		newIndex2 := min(index2 + half, len(nums2)) - 1 // 3 前进k/2格
		pivot1, pivot2 := nums1[newIndex1], nums2[newIndex2]
		if pivot1 <= pivot2 {
			k -= (newIndex1 - index1 + 1)
			index1 = newIndex1 + 1
		} else {
			k -= (newIndex2 - index2 + 1)
			index2 = newIndex2 + 1
		}
	}
	return 0
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

//最长回文子串 中心扩展算法
func longestPalindrome(s string) string {
	if s == "" {
		return ""
	}
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		left1, right1 := expandAroundCenter(s, i, i)
		left2, right2 := expandAroundCenter(s, i, i + 1)
		if right1 - left1 > end - start {
			start, end = left1, right1
		}
		if right2 - left2 > end - start {
			start, end = left2, right2
		}
	}
	return s[start:end+1]
}

func expandAroundCenter(s string, left, right int) (int, int) {
	for ; left >= 0 && right < len(s) && s[left] == s[right]; left, right = left-1 , right+1 { }
	return left + 1, right - 1
}
//最长回文子串 动态规划
func longestPalindromes(s string) string {
	arr := make([][]int,len(s))
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		arr[i] = make([]int,len(s))
		arr[i][i] = 1
	}
	for L := 1; L < len(s); L++ {
		for i := 0; i < len(s); i++ {
			j := L+i
			if j>=len(s) {
				break
			}
			if s[i] == s[j] {
				if j-i<=2 {
					arr[i][j] = 1
				}else {
					arr[i][j] = arr[i+1][j-1]
				}
			}else {
				arr[i][j] = 2
			}
			if arr[i][j] == 1 &&  j-i>end-start{
				start = i
				end = j
			}
		}
	}
	return s[start:end+1]
}
//Z字形变换 压缩矩阵空间
func convert(s string, numRows int) string {
	r := numRows
	if r == 1 || r >= len(s) {
		return s
	}
	mat := make([][]byte, r)
	t, x := r*2-2, 0 //t4
	for i, ch := range s {
		mat[x] = append(mat[x], byte(ch))
		fmt.Println(x)
		if i%t < r-1 {
			x++
		} else {
			x--
		}
	}
	fmt.Println(mat)
	return string(bytes.Join(mat, nil)) //数组转byte，byte字符串，
}
func convertys(s string,numRows int) string {
	if numRows == 1 || numRows >= len(s) {
		return s
	}
	zq := 2*numRows -2 //一个周期 6 numRows 4
	down := 0 //向下的key
	arr := make([][]byte,numRows)
	for i, v := range s {
		arr[down] = append(arr[down],byte(v))
		if i%zq < numRows-1 {
			down++
		}else {
			down--
		}
	}
	return string(bytes.Join(arr,nil))
}
//Z字形变换 直接构造
func converts(s string, numRows int) string {
	n, r := len(s), numRows
	if r == 1 || r >= n {
		return s
	}
	t := r*2 - 2
	ans := make([]byte, 0, n)
	for i := 0; i < r; i++ { // 枚举矩阵的行
		for j := 0; j+i < n; j += t { // 枚举每个周期的起始下标
			ans = append(ans, s[j+i]) // 当前周期的第一个字符
			if 0 < i && i < r-1 && j+t-i < n {
				ans = append(ans, s[j+t-i]) // 当前周期的第二个字符
			}
		}
	}
	return string(ans)
}
func reverse(x int) (rev int) {
	for x != 0 {
		if rev < math.MinInt32/10 || rev > math.MaxInt32/10 {
			return 0
		}
		digit := x % 10
		x /= 10
		rev = rev*10 + digit
	}
	return
}

func reverses(x int) (rev int) {
	for x!=0 {
		if rev>math.MaxInt32/10 || rev<math.MinInt32/10 {
			return 0
		}
		yu := x%10
		rev = rev*10+yu
		x/=10

	}
	return rev
}

// -ddf81555
func myAtoi(s string) int {
	abs, sign, i, n := 0, 1, 0, len(s)
	//丢弃无用的前导空格
	for i < n && s[i] == ' ' {
		i++
	}
	//标记正负号
	if i < n {
		if s[i] == '-' {
			sign = -1
			i++
		} else if s[i] == '+' {
			sign = 1
			i++
		}
	}
	for i < n && s[i] >= '0' && s[i] <= '9' {
		abs = 10*abs + int(s[i]-'0')  //字节 byte '0' == 48
		if sign*abs < math.MinInt32 { //整数超过 32 位有符号整数范围
			return math.MinInt32
		} else if sign*abs > math.MaxInt32 {
			return math.MaxInt32
		}
		i++
	}
	return sign * abs
}

//回文数
func isPalindrome(x int) bool {
	if x<0 || (x%10 == 0 && x!= 0){
		return false
	}
	me := 0
	for x>me {
		me = me*10 + x%10
		x = x/10
	}
	return me == x || x == me/10
}

//10 正则表达式匹配s kkjkkkkjr ， p .kjk*.*
func isMatch(s string, p string) bool {
	m, n := len(s), len(p) //m9 , n7
	matches := func(i, j int) bool {
		if i == 0 {
			return false
		}
		if p[j-1] == '.' {
			return true
		}
		return s[i-1] == p[j-1]
	}

	f := make([][]bool, m + 1)
	for i := 0; i < len(f); i++ {
		f[i] = make([]bool, n + 1)
	}
	f[0][0] = true

	for i := 0; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == '*' {
				f[i][j] = f[i][j] || f[i][j-2]
				if matches(i, j - 1) {
					f[i][j] = f[i][j] || f[i-1][j]
				}
			} else if matches(i, j) {
				f[i][j] = f[i][j] || f[i-1][j-1]
			}
		}
	}
	fmt.Println(f[1])
	return f[m][n]
}