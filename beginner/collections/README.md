# Collections

## Notes on Collections in Go

Collections in Go refer to data structures that hold multiple values. The primary built-in collections are **arrays**, **slices**, and **maps**. Understanding these is essential for efficient data handling in Go programs.

### Arrays
- Fixed-size, contiguous blocks of memory holding elements of the same type.
- Declared with syntax: `var arr [size]Type` or `arr := [size]Type{values}`.
- Example: `var nums [5]int = [5]int{1, 2, 3, 4, 5}`.
- Length is part of the type; cannot resize dynamically.

### Slices
- Dynamic arrays built on top of arrays; more flexible.
- Declared with syntax: `var slice []Type` or `slice := []Type{values}`.
- Can be created from arrays: `slice := arr[1:4]`.
- Key operations: `append()`, `len()`, `cap()`.
- Example: `slice := []int{1, 2, 3}; slice = append(slice, 4)`.

### Maps
- Key-value pairs; unordered collections.
- Declared with syntax: `var m map[KeyType]ValueType` or `m := make(map[KeyType]ValueType)`.
- Operations: `m[key] = value`, `value := m[key]`, `delete(m, key)`.
- Check existence: `value, exists := m[key]`.
- Example: `m := map[string]int{"a": 1, "b": 2}`.

Collections are fundamental for storing and manipulating data efficiently. Use slices for most dynamic needs, arrays for fixed-size data, and maps for associative data.

## Practice Questions

1. **Array Basics**: Declare an array of 3 strings and initialize it with "Go", "is", "fun". Print the array and its length.

2. **Slice Operations**: Create a slice of integers with values 1, 2, 3, 4, 5. Append 6 and 7 to it, then print the slice, its length, and capacity.

3. **Map Usage**: Create a map with string keys and int values representing ages: "Alice": 25, "Bob": 30. Add "Charlie": 35, then delete "Bob" and print the map.

4. **Slice from Array**: Given `arr := [6]int{10, 20, 30, 40, 50, 60}`, create a slice from index 1 to 4 (inclusive start, exclusive end) and print it.

5. **Map Existence Check**: Using the map from question 3, check if "Alice" exists and print her age if she does; otherwise, print "Not found".
