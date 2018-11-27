<h1 align="center">Mistakes I Made</h1>

## Never write comments

Don't forget to delete and replace all the stub comments with real comments. You might be surprised at how important accurate, helpful comments are to Go developers. This is a great guide to writing good comments: [Effective Go](https://golang.org/doc/effective_go.html#commentary): Commentary. Remember, the godoc tool will produce documentation from your code comments automatically. See [godoc.org](https://godoc.org/) for examples of what this looks like.

## Repeating most of the output string

Currently you are repeating most of the output string. This is OK, but in general, Go programmers try to avoid duplicating code where they can avoid it. Can you work out how to restructure your code so that you don't repeat yourself?

## Concurrency is not Parallelism

Concurrency is often mentioned as one of the things Go does really well. But what actually is it, how does it relate to parallelism, and how do we use Go to solve problems in a concurrent way? Rob Pike's talk entitled [Concurrency is not Parallelism explains](https://www.youtube.com/watch?v=cN_DpYBzKso&feature=youtu.be).

## Code Reviews

Having other people review your code can be incredibly helpful, and of course one of the great things about Exercism is that you get code reviews from experienced Go programmers. [Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments) is a great collection of common mistakes and Go style errors which should be picked up in code reviews. It can help you avoid these problems in your own code.

## Declare a variable too earlier

You are declaring distance before you check if the length is equal, so the declaration may end up being unnecessary. It's better to declare a variable right before you use it.

## Pointer to an interface

A pointer to a struct and a pointer to an interface are not the same. An interface can store either a struct directly or a pointer to a struct. In the latter case, you still just use the interface directly, not a pointer to the interface.
