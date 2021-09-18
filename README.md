# Performance test => Go vs Javascript

This repository contains code to compare Javascript and Go code in reguards to performance.

For the test, I implemented a collatz conjecture algorithm in both languages.

- OS: Windows 10 (desktop);
- Terminal: git (in VS Code);

# Benchmarks

Here are some results on my machine (in ms) per branch:

## less_prints:

Execution time (smaller = better)

Iteration | Go | JavaScript | Winner | Difference
--- | --- | --- | --- | --- |
0 | 1 | 6 | Go | 5
1 | 2 | 6 | Go | 4
2 | 1 | 7 | Go | 6
3 | 1 | 7 | Go | 6
4 | 1 | 7 | Go | 6
5 | 2 | 7 | Go | 5

Winner: Go

### Difference average (bigger = better) 
Go | JavaScript
--- | --- |
5.33 | NaN

Winner:  Go

## more_prints:

Execution time (smaller = better)

Iteration | Go | JavaScript | Winner | Difference
--- | --- | --- | --- | --- |
0 | 8275 | 8296  | Go | 21
1 | 8241 | 7728 | JavaScript | 513
2 | 8390 | 8423 | Go | 33
3 | 8579 | 8064 | JavaScript | 515
4 | 8452 | 8384 | JavaScript | 68
5 | 8086 | 8636 | Go | 550

Winner: Draw

### Difference average (bigger = better) 
Go | JavaScript
--- | --- |
201.33 | 365.33

Winner: JavaScript

# Conclusion

Go seems to be way more performant when there are few to none logs, which is the really for most cases. However, it lost to JavaScript by a small margin when there were more logs.

The second there should be taken into account, but only secondarily. Its only tell us that go interface with the terminal is not as good as JavaScript's. And that may not even be true for other terminals or operating systems.

Taken that into consideration, Go won.