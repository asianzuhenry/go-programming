# Generics
Generics allow you to write flexible, reusable functions and types that can work with any type, while still maintaining type safety. In Swift, you can use generics to create functions, types, and protocols that can operate on any type.

## Generic Functions
You can define a generic function by using angle brackets (`<T>`) to specify a placeholder type. For example:
```swift
func swapTwoValues<T>(_ a: inout T, _ b: inout T) {
    let temporaryA = a
    a = b
    b = temporaryA
}
```
In this example, `T` is a placeholder type that can be replaced with any type when the function is called. The `swapTwoValues` function can swap values of any type, as long as both parameters are of the same type.

## Generic Types
You can also define generic types, such as classes, structures, and enumerations. For example:
```swift
struct Stack<Element> {
    var items = [Element]()
    
    mutating func push(_ item: Element) {
        items.append(item)
    }
    
    mutating func pop() -> Element {
        return items.removeLast()
    }
}
```
In this example, `Element` is a placeholder type that can be replaced with any type when an instance of `Stack` is created. The `Stack` structure can be used to create stacks of any type, such as `Stack<Int>` or `Stack<String>`. 

## Generic Protocols
You can also define generic protocols. For example:
```swift
protocol Container {
    associatedtype Item
    mutating func append(_ item: Item)
    var count: Int { get }
    subscript(i: Int) -> Item { get }
}
```
In this example, `Item` is an associated type that can be replaced with any type when a type conforms to the `Container` protocol. Types that conform to this protocol must specify what type they will use for `Item` and implement the required methods and properties.

Generics are a powerful feature in Swift that allow you to write flexible and reusable code while maintaining type safety. By using generics, you can create functions, types, and protocols that can work with any type, making your code more versatile and easier to maintain.
