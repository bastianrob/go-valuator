# go-valuator

With this package, you can evaluate simple expression

## Supported Operator

* =
* !=
* in
* !in
* intersect
* \>
* \>=
* <
* <=

## Dependencies

```bash
"github.com/bastianrob/arrayutil"
```

## Example

### Equality evaluation

```go
eval, err := valuator.NewValuator("name", "=", "John Doe", "Name must be equal to John Doe")
result := eval.Evaluate(map[string]interface{}{
    "name": "John Doe"
})
fmt.Println("Err:",err, "Res:", result)
//Err: <nil> Res: true
```

### In evaluation

```go
eval, err := valuator.NewValuator("status", "in", `["Open", "Close"]`, "Status must be one of")
result1 := eval.Evaluate(map[string]interface{}{
    "status": "Void"
})
result2 := eval.Evaluate(map[string]interface{}{
    "status": "Open"
})
fmt.Println("Err:",err, "Res1:", result1, "Res2:", result2)
//Err: <nil> Res1: false Res2: true
```

### Intersect evaluation

```go
eval, err := valuator.NewValuator("number", "intersect", `[1, 2, 3, 4, 5]`, "number intersects")
result1 := eval.Evaluate(map[string]interface{}{
    "number": []int{6, 7},
})
result2 := eval.Evaluate(map[string]interface{}{
    "number": []int{1, 3, 7},
})
fmt.Println("Err:",err, "Res1:", result1, "Res2:", result2)
//Err: <nil> Res1: false Res2: true
```

### > evaluation

```go
eval, err := valuator.NewValuator("number", ">", "100", "Number must be more than 100")
result1 := eval.Evaluate(map[string]interface{}{
    "number": 100
})
result2 := eval.Evaluate(map[string]interface{}{
    "number": "100.1" //accept either string, int, or float
})
fmt.Println("Err:",err, "Res1:", result1, "Res2:", result2)
//Err: <nil> Res1: false Res2: true
```

### >= evaluation

```go
eval, err := valuator.NewValuator("number", ">=", "100", "Number must be more than 100")
result1 := eval.Evaluate(map[string]interface{}{
    "number": 99.9
})
result2 := eval.Evaluate(map[string]interface{}{
    "number": "100" //accept either string, int, or float
})
fmt.Println("Err:",err, "Res1:", result1, "Res2:", result2)
//Err: <nil> Res1: false Res2: true
```

### < evaluation

```go
eval, err := valuator.NewValuator("number", "<", "100", "Number must be less than 100")
result1 := eval.Evaluate(map[string]interface{}{
    "number": 100
})
result2 := eval.Evaluate(map[string]interface{}{
    "number": "100.1" //accept either string, int, or float
})
fmt.Println("Err:",err, "Res1:", result1, "Res2:", result2)
//Err: <nil> Res1: false Res2: true
```

### <= evaluation

```go
eval, err := valuator.NewValuator("number", "<=", "100", "Number must be less than 100")
result1 := eval.Evaluate(map[string]interface{}{
    "number": 100.1
})
result2 := eval.Evaluate(map[string]interface{}{
    "number": "100" //accept either string, int, or float
})
fmt.Println("Err:",err, "Res1:", result1, "Res2:", result2)
//Err: <nil> Res1: false Res2: true
```
