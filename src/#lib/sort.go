/**
 * To use custom sorts, follow these steps:
 *   1.  Replace T with the type of the array.
 *   2.  Replace SortName with a descriptive name of the sorting class.
 *   3.  Update the Less functon to perform the custom comparison.
 *   4.  Call `sort.Sort(SortName(arrayName))` to sort the array called arrayName.
 */

type SORTNAME []TYPE

func (s SORTNAME) Len() int {
    return len(s)
}

func (s SORTNAME) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}

func (s SORTNAME) Less(i, j int) bool {
    return s[i] < s[j]
}
