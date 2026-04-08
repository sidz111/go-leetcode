/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeElements(head *ListNode, val int) *ListNode {
    
    dummy := &ListNode{Next: head}
    current := dummy

    for current.Next != nil {
        if current.Next.Val == val {
            // this is for skip node
            current.Next = current.Next.Next
        } else {
            current = current.Next
        }
    }

    return dummy.Next
}