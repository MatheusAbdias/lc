package hashset

type MyHashSet struct {
	Keys []int
}

func Constructor() MyHashSet {
	return MyHashSet{}
}

func (this *MyHashSet) Add(key int) {
	if key >= len(this.Keys) {
		for index := len(this.Keys); index <= key; index++ {
			this.Keys = append(this.Keys, -1)
		}
		this.Keys[key] = key
	} else if this.Keys[key] != key {
		this.Keys[key] = key
	}

}

func (this *MyHashSet) Remove(key int) {
	if key < len(this.Keys) {
		this.Keys[key] = -1
	}
}

func (this *MyHashSet) Contains(key int) bool {
	if key < len(this.Keys) {
		return this.Keys[key] == key
	}
	return false
}
