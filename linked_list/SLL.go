package linked_list

type Node struct {
	Name string
	Next *Node
}

type LinkedList struct {
	Head *Node
	End *Node
	Length int
}

func NewLinkedList() *LinkedList{
	return &LinkedList{}
}

func (ll *LinkedList) Enqueue(value string) int {
	
	var node Node
	node.Name = value
	node.Next = nil
	
	if ll.Length == 0 {
		
		ll.Head = &node
		ll.End = &node
		ll.Length++
		return ll.Length
	} 

	ll.End.Next = &node
	ll.End = &node
	ll.Length++
	return ll.Length 
}

func (ll *LinkedList) Dequeue() *Node {
	
	temp := ll.Head
	if ll.Length == 0 {
		return nil
	}
	if ll.Length == 1 {
		
		ll.Head = nil
		ll.End = nil
		ll.Length = 0
		return temp
	} 
	
	ll.Head = ll.Head.Next
	
	ll.Length--
	return temp
}

func (ll *LinkedList) Delete(value string) bool {
	
	if ll.Length == 0 {
		return true
	}

	node := ll.Head

	for (node.Next != nil){
		
		if node.Name == value {

			if node == ll.Head {
				ll.Head = ll.Head.Next
			} 

			if node == ll.End {
				ll.Dequeue()
			}

			ll.Length--
			return true
		}
		node.Next = node.Next.Next
	}
	return false
}