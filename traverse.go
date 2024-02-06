package datastructure

// VisitFunc is the function to visit the data
// NOTE: This callback is used for readonly, changing the compare key will lead to undefined behavior.
type VisitFunc[T any] func(T)

// ConditionFunc is the function to visit the data and check condition
// NOTE: This callback is used for readonly, changing the compare key will lead to undefined behavior.
type ConditionFunc[T any] func(T) bool

// ModifyFunc is the function to modify the data
// NOTE: This callback should not change the old data directly, make a copy of old and modify the copy instead.
type ModifyFunc[T any] func(old T) (new T)
