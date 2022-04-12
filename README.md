# golang-di

A simple project demonstrating the idiom "accept interfaces return structs".

In `main.go` we create a mock (in-memory) and a real (persistence-aware) repository instances. 
Each of these is independently injected into a service as a dependency. Being able to do this 
allows for looser coupling and easier testing.

