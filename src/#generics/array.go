/*
 * Array utilties.
 *
 * RESERVED GENERIC NAMES:
 *   - TYPE: type of the array to fill.
 */

func FillArray(arr []TYPE, val TYPE) {
	for i := range arr {
		arr[i] = val
	}
}
