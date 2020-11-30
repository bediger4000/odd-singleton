# Daily Coding Problem: Problem #726 [Medium]

This problem was asked by Microsoft.

Implement the singleton pattern with a twist.
First, instead of storing one instance, store two instances.
And in every even call of getInstance(),
return the first instance and in every odd call of getInstance(),
return the second instance.

## Analysis

* [singleton Go package](singleton)
* [Single-threaded use](singleton1.go), run like: `./singleton1 21` to see alternating "singleton".
* [Multi-threaded use](singleton2.go), run like: `./singleton2 100` to see threaded execution

Of course Microsoft asked this,
it's a strange, strange question.

First, if there are *2* instances, it's not a singleton.
The entire point of the pattern is to control access to something
that exists as a single, non-interchangeable entity.

Second, there are [well-known problems](https://stackoverflow.com/questions/1392315/problems-with-singleton-pattern)
with a lot of implementations of the Singleton Pattern.
The biggest problem in the programming domain is that
it's easy to introduce data race conditions in a multi-threaded program.

If the job candidate gets past the cognitive dissonance of writing
a multi-instance *singleton* class,
there's plenty of pitfalls to trigger.
I'm not claiming my code doesn't trigger any of those problems
except that it probably doesn't hit the data race.
You can compile like this to test:

```sh
$ go build -race singleton2.go
$ ./singleton2 100
```

The candidate should talk about the data race problem,
and mention the other problems, which include hard-to-test systems.
Since a singleton pattern does not make sense in a single-threaded-program
(any use of a resource is a singleton use),
talking about the data race and how to avoid it seem paramount.

It's difficult for Go programmers to talk about how one might make a coherent
package out of this singleton.
The problem statement specifies no behavior
of a singleton object, making it hard to write a Go interface.
Further, the problem statement specifies a `getSingleton()` function
that doesn't allow a programmer to pass an argument specifying
which implementation of an interface to return.

The interviewer shouldn't expect much from a candidate,
as this problem is almost incoherent.
