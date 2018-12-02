package splaytree

import (
	"strings"
	"testing"
)

type ComparableString string

func (sv ComparableString) CompareTo(i interface{}) int {
	if ts, ok := i.(ComparableString); ok {
		return strings.Compare(string(sv), string(ts))
	}

	return -1
}

type orderStep struct {
	s string
	l int
}

func assertExpectedOrder(actual, expected []orderStep, t *testing.T) {
	for i, v := range actual {
		if v.s != expected[i].s || v.l != expected[i].l {
			t.Fatalf("Incorrect order")
		}
	}
}

func executeTreeTraverse(tree *SplayTree) []orderStep {
	result := []orderStep{}
	tree.Traverse(func(c Comparable, level int) bool {
		result = append(result, orderStep{
			s: string(c.(ComparableString)),
			l: level,
		})

		return true
	})

	return result
}

func TestSplayTree(t *testing.T) {
	tree := NewSplayTree()
	tree.Insert(ComparableString("A"))
	tree.Insert(ComparableString("B"))
	tree.Insert(ComparableString("C"))
	tree.Insert(ComparableString("D"))
	tree.Insert(ComparableString("E"))

	t1 := executeTreeTraverse(tree)
	assertExpectedOrder(t1, []orderStep{
		{"A", 4}, {"B", 3}, {"C", 2}, {"D", 1}, {"E", 0},
	}, t)

	tree.Has(ComparableString("A"))
	t2 := executeTreeTraverse(tree)
	assertExpectedOrder(t2, []orderStep{
		{"A", 0}, {"B", 2}, {"C", 3}, {"D", 1}, {"E", 2},
	}, t)

	tree.Has(ComparableString("C"))
	t3 := executeTreeTraverse(tree)
	assertExpectedOrder(t3, []orderStep{
		{"A", 1}, {"B", 2}, {"C", 0}, {"D", 1}, {"E", 2},
	}, t)

	tree.Insert(ComparableString("F"))
	t4 := executeTreeTraverse(tree)
	assertExpectedOrder(t4, []orderStep{
		{"A", 2}, {"B", 3}, {"C", 1}, {"D", 3}, {"E", 2}, {"F", 0},
	}, t)

	tree.Insert(ComparableString("G"))
	t5 := executeTreeTraverse(tree)
	assertExpectedOrder(t5, []orderStep{
		{"A", 3}, {"B", 4}, {"C", 2}, {"D", 4}, {"E", 3}, {"F", 1}, {"G", 0},
	}, t)

	tree.Has(ComparableString("D"))
	t6 := executeTreeTraverse(tree)
	assertExpectedOrder(t6, []orderStep{
		{"A", 2}, {"B", 3}, {"C", 1}, {"D", 0}, {"E", 2}, {"F", 1}, {"G", 2},
	}, t)

	tree.Has(ComparableString("B"))
	t7 := executeTreeTraverse(tree)
	assertExpectedOrder(t7, []orderStep{
		{"A", 1}, {"B", 0}, {"C", 2}, {"D", 1}, {"E", 3}, {"F", 2}, {"G", 3},
	}, t)

	tree.Remove(ComparableString("D"))
	t8 := executeTreeTraverse(tree)
	assertExpectedOrder(t8, []orderStep{
		{"A", 2}, {"B", 1}, {"C", 0}, {"E", 2}, {"F", 1}, {"G", 2},
	}, t)

	tree.Remove(ComparableString("B"))
	t9 := executeTreeTraverse(tree)
	assertExpectedOrder(t9, []orderStep{
		{"A", 1}, {"C", 0}, {"E", 2}, {"F", 1}, {"G", 2},
	}, t)

	tree.Remove(ComparableString("E"))
	t10 := executeTreeTraverse(tree)
	assertExpectedOrder(t10, []orderStep{
		{"A", 1}, {"C", 0}, {"F", 1}, {"G", 2},
	}, t)

	tree.Remove(ComparableString("F"))
	t11 := executeTreeTraverse(tree)
	assertExpectedOrder(t11, []orderStep{
		{"A", 1}, {"C", 0}, {"G", 1},
	}, t)

	tree.Insert(ComparableString("B"))
	t12 := executeTreeTraverse(tree)
	assertExpectedOrder(t12, []orderStep{
		{"A", 1}, {"B", 0}, {"C", 1}, {"G", 2},
	}, t)

	tree.Remove(ComparableString("A"))
	t13 := executeTreeTraverse(tree)
	assertExpectedOrder(t13, []orderStep{
		{"B", 0}, {"C", 1}, {"G", 2},
	}, t)

	tree.Insert(ComparableString("A"))
	t14 := executeTreeTraverse(tree)
	assertExpectedOrder(t14, []orderStep{
		{"A", 0}, {"B", 1}, {"C", 2}, {"G", 3},
	}, t)

	tree.Has(ComparableString("G"))
	t15 := executeTreeTraverse(tree)
	assertExpectedOrder(t15, []orderStep{
		{"A", 1}, {"B", 3}, {"C", 2}, {"G", 0},
	}, t)

	tree.Remove(ComparableString("G"))
	t16 := executeTreeTraverse(tree)
	assertExpectedOrder(t16, []orderStep{
		{"A", 0}, {"B", 2}, {"C", 1},
	}, t)

	mustBeFalse := tree.Has(ComparableString("G"))
	if mustBeFalse {
		t.Fatalf("Tree contains value it should not")
	}
	t17 := executeTreeTraverse(tree)
	assertExpectedOrder(t17, []orderStep{
		{"A", 0}, {"B", 2}, {"C", 1},
	}, t)

	mustBeFalse = tree.Remove(ComparableString("G"))
	if mustBeFalse {
		t.Fatalf("Removed values that should not be int the tree")
	}
	t18 := executeTreeTraverse(tree)
	assertExpectedOrder(t18, []orderStep{
		{"A", 0}, {"B", 2}, {"C", 1},
	}, t)

	tree.Remove(ComparableString("C"))
	t19 := executeTreeTraverse(tree)
	assertExpectedOrder(t19, []orderStep{
		{"A", 0}, {"B", 1},
	}, t)
}
