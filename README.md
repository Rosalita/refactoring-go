# refactoring-go
Martin Fowlers refactoring book adapted to Go

"Any fool can write code that a computer can understand. Good programmers write code that humans can understand."

# Advice from Martin.
- "Before you start refactoring, make sure you have a solid suite of tests. These tests must be self-checking."
- "Run tests immediately after every change. Commit after each successful refactoring." 
- "Many programmers are uncomfortable with splitting a single loop that does two different things into two 
   separate loops, as this forces two loops to execute. My reminder as usual is to separate refactoring from 
   optimisation. Once code is clear, then optimise it. If loop traversal is a bottle neck, its easy to slam two
   loops back together. But the actual iteration through even a large list is rarely a bottle neck and splitting 
   loops enables other, more powerful, optimisations.

# Learning
- It might not be a good idea to pass data down to a function, if that function can be empowered to read the data for itself.
  This is replacing 'temp with query' e.g. replacing a temporary variable with a function that knows how to find the value.
- If you see a function with the wrong name, change it as soon as you understand what a better name could be. Then you won't
  have to figure out what it does again.
- If a value returned from a function is always transformed in the same way, that transformation should probably move into the function.
- If a loop does two different things at once, when you need to modify the loop you must understand both things. Splitting a loop that
  does two things into two loops ensures you only need to understand the behaviour you need to modify.
- Declare variables just before using them, move related code together

# Observations about Go
- It's not possible to inline a Go function which returns a value and an error as errors must be handled.