# Pointers

## Notes on Pointers in Go

Pointers in Go are variables that store the memory address of another variable. They allow you to indirectly access and modify the value of a variable through its address.

### Key Concepts

1. **Declaration**: Use `*T` for pointer type, where `T` is the type being pointed to.
   ```go
   var p *int  // p is a pointer to an int
   ```

2. **Initialization**: Use `&` to get the address of a variable.
   ```go
   x := 42
   p := &x  // p now points to x
   ```

3. **Dereferencing**: Use `*` to access the value at the pointer.
   ```go
   *p = 50  // changes x to 50
   fmt.Println(*p)  // prints 50
   ```

4. **Zero Value**: Uninitialized pointers are `nil`.
   ```go
   var p *int  // p is nil
   ```

5. **Passing by Reference**: Functions can modify the original value by accepting pointers.
   ```go
   func modify(x *int) {
       *x = 100
   }
   ```

### Common Use Cases

- Modifying values in functions
- Working with large structs (avoid copying)
- Implementing data structures like linked lists
- Interfacing with C code

### Important Notes

- Go doesn't have pointer arithmetic like C
- Pointers can be compared with `==` and `!=`
- Be careful with nil pointer dereferences (runtime panic)

## Practice Questions

1. **Basic Declaration**: Declare a pointer to an `int` and initialize it to point to a variable `x` with value 10.

2. **Dereferencing**: Given `p := &x` where `x = 5`, write code to change `x` to 15 using the pointer.

3. **Nil Pointers**: What happens if you try to dereference a nil pointer? How can you check for nil?

4. **Function Modification**: Write a function that takes a pointer to an `int` and doubles its value.

5. **Struct Pointers**: Create a struct `Person` with fields `Name` and `Age`. Write a method that modifies the `Age` using a pointer receiver.

6. **Pointer to Pointer**: Declare a pointer to a pointer to an `int`. Initialize it properly and access the final value.

7. **Slice Pointers**: Explain why you don't usually need pointers to slices in Go.

8. **Memory Address**: Write code to print the memory address of a variable.

### Answers

1. ```go
   x := 10
   p := &x
   ```

2. ```go
   *p = 15
   ```

3. It causes a runtime panic. Check with `if p != nil { *p = ... }`

4. ```go
   func double(n *int) {
       *n *= 2
   }
   ```

5. ```go
   type Person struct {
       Name string
       Age  int
   }
   
   func (p *Person) Birthday() {
       p.Age++
   }
   ```

6. ```go
   x := 42
   p := &x
   pp := &p
   fmt.Println(**pp)  // prints 42
   ```

7. Slices are already reference types; they contain a pointer to the underlying array.

8. ```go
   x := 10
   fmt.Printf("%p\n", &x)
   ```