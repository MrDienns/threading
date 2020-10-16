# Threading
This homework assignment demonstrates several concepts of multi threading; it demonstrates how things can go wrong (race-condition) and it demonstrates how things can work correctly (password-cracker). Below are the assignment questions with their respective answers. Below that is a quick introduction into the two projects in this repository.

### What is multi threading?
Multi threading allows a CPU to execute instructions concurrently and independently within the same process and the process resources. Even if a computer has multiple cores & threads, if your application doesn't utilise these threads then this single application will not have any direct benefits of the additional cores & threads. To put it simple; multi threading allows you to do work concurrently (in parallel), which can speed up an application that's performing heavy tasks (see password-cracker project).

### When to use multiple threads?
Multi threading should be applied when we have the option of segregating our work across multiple parallel tasks, in order to speed up the larger task. An extremely common use-case is any IO related work. IO, such as reading files, contacting a REST API, or executing database queries will take extremely long compared to traditional operations within an application. As a result, these operations can take milliseconds, or even seconds depending on the IO action that is being done. If the application isn't multi threaded, then all work has to wait for this IO operation to complete. The application would freeze as a result. For a desktop application, this can be annoying. For an important backend service, this can be catastrophic. It is thus common practice to perform IO actions in a separate thread, to prevent blocking the main (or any other important thread).

### What are three common issues surrounding multi threading?
1) Increased complexity.

Correct application of multi threading is not easy. You can't simply create a chuckload of threads and tell them to get cracking. When applying multi threading, you have to carefully align the individual tasks to make sure they do the individual work correctly.

2) Unpredictable results.

If multi threading is not done correctly, then unexpected results may pop up. You can run into race conditions which can give your application a completely different outcome, without actually throwing an error clearly stating that this happened.

3) Difficult to identify errors

When errors, race conditions or otherwise bugs created by multi threading show up, it is often problematic to debug these. The application essentially runs in parallel. Unless this is all handled safely, something can go wrong. Specific issues, such as race conditions, don't throw any errors. As a result, figuring out where race conditions occur can be time and hair consuming.

## Race condition
In the race condition project, we see a very simple counter. This counter is a shared object, shared by two threads. Two different threads will increment and decrement the counter at the same time, unsafely. One counts up a million times, the other counts down a million times. If the threading was done safely, the outcome of the number should be zero. However, running the application we see a completely random number being outputted each time.

## Password cracker
In this project, I built a small Golang based BCrypt password cracker. BCrypt is quite a heavy hashing algorithm, which means cracking passwords one by one in sequence will take a long time. In fact, testing on my desktop shows that comparing roughly 100 passwords takes roughly 19 seconds without any Go routines. After applying Go routines (100 by default), the processing time is dramatically dropped to just over one second while still comparing as many passwords as it did previously.