func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	
    var ans *ListNode = l1
	ans = l1
	var r int
	r = 0

	for l1 != nil && l2 != nil {
		l1.Val = (l1.Val + l2.Val + r)
		r = l1.Val / 10
        l1.Val %= 10
        if (l1.Next == nil && l2.Next == nil && r > 0){
            temp := ListNode{r, nil}
            l1.Next = &temp
            r = 0
            break
        } else if (l1.Next == nil){
            l1.Next = l2.Next
            l1 = l1.Next
            break
        }
        l1 = l1.Next
		l2 = l2.Next
	}

    for l1 != nil {
		l1.Val += r
        r = l1.Val / 10
        l1.Val %= 10
        if (l1.Next == nil && r > 0){
            temp := ListNode{r, nil}
            l1.Next = &temp
            break
        }
		l1 = l1.Next
	}

	return ans
}
