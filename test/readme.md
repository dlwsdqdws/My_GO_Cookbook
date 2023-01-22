- [Test](#test)
  - [Unit Test](#unit-test)
  - [Mock](#mock)
  - [Benchmark](#benchmark)


# Test

## Unit Test

- All test files end with `_test.go`
- Test function `func TestXxxx(t *testing.T)`
- Initialization should be included in `TestMain` function.

```go
func TestMain(m *testing.M) {
   // do init
   code := m.Run() // run test
   // do close and release
   os.Exit(code)
}
```

- `assert` package

```go
import "github.com/stretchr/testify/assert"

func TestFunc(t *testing.T) {
    output := Func()
    expectOutput := ""
    assert.Equal(t, expectOutput, output)
}
```

- **Code coverage** is the standard for judging whether the tests are good or bad. Using `--cover` can get the code coverage. It is good to make it more than 50%, better for more than 80%.

For more examples, please refer to https://juejin.cn/post/6908938380114034701.

## Mock

- Use the reflection mechanism to replace some functions such as R/W to achieve idempotent and make sure that the test is not affected by the environment or other users.
- [monkey package](https://github.com/bouk/monkey) : can mock methods or instances.

```go
// replace a funtion with another
func Patch(target, replacement interface{}) *PatchGuard {
	t := reflect.ValueOf(target)
	r := reflect.ValueOf(replacement)
	patchValue(t, r)
	return &PatchGuard{t, r}
}

// remove all monkey patches on target
// return if the target was patched in the first place
func Unpatch(target interface{}) bool {
	return unpatchValue(reflect.ValueOf(target))
}

func TestProcessFirstLineWithMock(t *testing.T) {
	monkey.Patch(ReadFirstLine, func() string {
		return "line110"
	})
	defer monkey.Unpatch(ReadFirstLine)
	line := ProcessFirstLine()
	assert.Equal(t, "line000", line)
}
```

## Benchmark

- Very similar with unit test. It is to test the performance of the code and the consumption of CPU.
- Start with `Benchmark` and using `gobench xxx_test.go` to run.

```go
// Serial
func BenchmarkRandom(b *testing.B) {
   InitServerIndex()
   b.ResetTimer()
   for i := 0; i < b.N; i++ {
      Random(10)
   }
}

// Parallel
func BenchmarkRandomParallel(b *testing.B) {
   InitServerIndex()
   b.ResetTimer()
   b.RunParallel(func(pb *testing.PB) {
      for pb.Next() {
         Random(10)
      }
   })
}
```