This repository contains [splay tree](https://en.wikipedia.org/wiki/Splay_tree) implementation.

This implementation does not allow duplicate elements in the tree, so it may be used as _set_ data type.

Usage.
```
    type ComparableString string

    func (sv ComparableString) CompareTo(i interface{}) int {
        if ts, ok := i.(ComparableString); ok {
            return strings.Compare(string(sv), string(ts))
        }

        return -1
    }

    ...

    tree := splaytree.NewSplayTree()
    tree.Insert(ComparableString("A"))
    tree.Insert(ComparableString("B"))
    tree.Insert(ComparableString("C"))
    ...

    if tree.Has(ComparableString("C")) {
        ...
    }
    ...

    if tree.Remove(ComparableString("B")) {
        // B was removed
        ...
    }

    tree.Traverse(func(c splaytree.Comparable, level int) bool {
        // level shows how deep in the tree the value is, root has value 0

        return canContinueTraverse // false to stop traverse
    })
```
