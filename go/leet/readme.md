### Testing

run `go test` from this directory, and check out the `*_test.go` files for test input params

run `go test -v -cover -benchmem` for verbosity, coverage analysis & memory allocation stats

```bash
go test $PWD/leet -v -cover -benchmem
# get a pretty UI detailing test coverage
go test $PWD/leet -v -cover -benchmem -coverprofile=coverage.out
go tool cover -html=coverage.out -o coverage.html
```

### List

- sums: https://leetcode.com/problems/two-sum/description/ (1)
- palindrome: https://leetcode.com/problems/palindrome-number/description/ (9)
- LongestPrefix: https://leetcode.com/problems/longest-common-prefix/description/ (14)
- Valid Parentheses: https://leetcode.com/problems/valid-parentheses/description/ (20)
- mistmatch: https://leetcode.com/problems/set-mismatch/description/ (645)
- findcenter: https://leetcode.com/problems/find-center-of-star-graph/description/ (1791)
- happiness: https://leetcode.com/problems/maximize-happiness-of-selected-children/description/ (3075)

go test -coverprofile=coverage.out
go tool cover -html=coverage.out -o coverage.html

outputs

// boolean
%t

// string
%s

// int
%d
