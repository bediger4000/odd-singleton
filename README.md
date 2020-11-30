# Daily Coding Problem: Problem #726 [Medium]

This problem was asked by Microsoft.

Implement the singleton pattern with a twist.
First, instead of storing one instance, store two instances.
And in every even call of getInstance(),
return the first instance and in every odd call of getInstance(),
return the second instance.

## Analysis

* [Single-threaded use](singleton1.go)
* [Multi-threaded use](singleton2.go)

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

The candidate should talk about the data race problem,
and mention the other problems, which include hard-to-test systems.
Since a singleton pattern does not make sense in a single-threaded-program
(any use of a resource is a singleton use),
talking about the data race and how to avoid it seem paramount.

The interviewer shouldn't expect much from a candidate,
as this problem is almost incoherent.
