# Parser Luna

This is my solution to the take home exercise for Luna.

## Instructions

(This instructions asume go is installed)

1. On one terminal run:

```bash
    go run cmd/main.go
```

2. On another one run the provided binary with:

```bash
    ./bin/test-program/darwin_arm64/rockets launch "http://localhost:8088/messages" --message-delay=500ms --concurrency-level=1
```

3. Then you can use (or import it into postman):

    a. `curl --location 'localhost:8088/rockets'`

    b. `curl --location 'localhost:8088/rockets/467cbde6-01a1-5883-b6df-9bb52c2a468d'`

## Design choices & comments

- The state is computed when reading. Not when ingesting messages. Main reason being the messages are out of order.
- Storage layer using in-memory implementation. For the sake of speed. It mimics the semantics typically used when working with real storaging solutions (ordering, eventual errors, no direct internal access, etc.)
- Full standard library. The test was simple enough to not use anything fancy.
- Not fully unit tested because of time and effort. Most of the tests would be very dummy and simple in this case. Added unit test using table tests for the app component since it's the one with the "business logic".
- Project structure is really anemic (most of the files are called "messages.go"). Folder structure is the one defining proper responsabilities.
- In `infrastructure/http/message.go` file the function `ToDomainMessage()` is probably the ugliest part of the codebase. At the same time, it's self-contained in the http package where it belongs, so didn't iterate further. I accept most probably there is a better approach to it.
- I did not see any Failure event. I think the implementation was correct, but worth mentioning the potential issue here.
- Next steps:
    - Add more testing coverage.
    - Building the binaries for the initially provided platforms would be my second choice.
    - In order to automate the previous, some makefile to automate the building and running the two services would be the next.
    - For a bigger feature, I would look into better (persistent) storaging solutions.
    -  There was a bonut point for sorting. I didn't do it (basically because of time).
