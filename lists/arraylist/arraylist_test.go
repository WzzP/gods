package arraylist

import "testing"

//arraylist 测试

func TestList_New(t *testing.T) {
	list1 := New()

	if actualValue := list1.Empty(); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}

	list2 := New(1, "b")

	if actualValue := list2.Size(); actualValue != 2 {
		t.Errorf("Got %v expected %v", actualValue, 2)
	}

	if actualValue, ok := list2.Get(0); actualValue != 1 || !ok {
		t.Errorf("Got %v expected %v ", actualValue, 1)
	}

	if actualValue, ok := list2.Get(1); actualValue != "b" || !ok {
		t.Errorf("Got %v expected %v", actualValue, "b")
	}

	if actualValue, ok := list2.Get(2); actualValue != nil || ok {
		t.Errorf("Got %v expected %v", actualValue, nil)
	}

}

func TestList_Add(t *testing.T) {

}

func TestList_Clear(t *testing.T) {

}

func TestList_Contains(t *testing.T) {

}

func TestList_Empty(t *testing.T) {

}

func TestList_Get(t *testing.T) {

}

func TestList_Insert(t *testing.T) {

}

func TestList_Remove(t *testing.T) {

}

func TestList_Set(t *testing.T) {

}

func TestList_Size(t *testing.T) {

}
func TestList_String(t *testing.T) {

}
func TestList_Swap(t *testing.T) {

}

func TestList_Values(t *testing.T) {

}
