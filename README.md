# refactoring-go
Martin Fowlers refactoring book adapted to Go

"Any fool can write code that a computer can understand. Good programmers write code that humans can understand."

# Advice from Martin.
- "Before you start refactoring, make sure you have a solid suite of tests. These tests must be self-checking."
- "Run tests immediately after every change. Commit after each successful refactoring." 

# Learning
- It might not be a good idea to pass data down to a function, if that function can be empowered to read the data for itself.
  This is replacing 'temp with query' e.g. replacing a temporary variable with a function that knows how to find the value.
