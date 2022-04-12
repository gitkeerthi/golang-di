# golang-di

A simple project demonstrating the idiom "accept interfaces return structs" to allow
seamless injection of dependencies for looser coupling and easier testing.

In `main.go` we create a mock (in-memory) and a real (persistence-aware) repository. 
Each of these is independently injected into a service as a dependency. 

