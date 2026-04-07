# Social Network Kata

## Source

<https://github.com/sandromancuso/social_networking_kata>

## Goal

Implement a networking application (similar to Twitter) satisfying the scenarios below using:

- Domain Driven Design and Hexagonal architecture
- Use case driven development
- In-memory database implementation

## Scenarios

### Posting: Alice can publish messages to a personal timeline

```cmd
> Alice -> I love the weather today
```

### Posting: Bob can publish messages to a personal timeline

```cmd
> Bob -> Damn! We lost!
> Bob -> Good game though.
```

### Reading: We can view Alice's timeline

```cmd
> Alice
> I love the weather today (5 minutes ago)
```

### Reading: We can view Bob's timeline

```cmd
> Bob
> Good game though. (1 minute ago)
> Damn! We lost! (2 minutes ago)
```

### Following: Charlie can subscribe to Alice's and Bob's timelines

```cmd
> Charlie -> I'm in New York today! Anyone want to have a coffee?
> Charlie follows Alice
> Charlie wall
> Charlie - I'm in New York today! Anyone want to have a coffee? (2 seconds ago)
> Alice - I love the weather today (5 minutes ago)

> Charlie follows Bob
> Charlie wall
> Charlie - I'm in New York today! Anyone want to have a coffee? (15 seconds ago)
> Bob - Good game though. (1 minute ago)
> Bob - Damn! We lost! (2 minutes ago)
> Alice - I love the weather today (5 minutes ago)
```

## Commands

| Command   | Format                               |
| --------- | ------------------------------------ |
| Posting   | `[user name] -> [message]`           |
| Reading   | `[user name]`                        |
| Following | `[user name] follows [another user]` |
| Wall      | `[user name] wall`                   |

## Implementation Notes

- Don't worry about handling exceptions or invalid commands
- Assume users will always type correct commands
- Focus on sunny day scenarios
- All in memory, no network/process communication
- Non-existing users should be created on first post
- Application should not start with pre-defined users

## Testing Strategy

| Aspect | Unit Tests        | Integration Tests | Acceptance Tests      |
| ------ | ----------------- | ----------------- | --------------------- |
| Scope  | Class/Aggregate   | Class ↔ External  | Full Application Flow |
| Data   | Mock data         | Test data         | Fake or Test Data     |
| Focus  | Isolated behavior | Integration layer | App flow              |
| Speed  | Very fast         | Medium to slow    | Very fast to fast     |

## Follow TDD Rules Strictly

1. ✅ Write production code only to pass a failing unit test
2. ✅ Write only enough of a unit test to make it fail
3. ✅ Write only enough production code to make the failing test pass

## Resources

- [TestDesiderata by Kent Beck](https://kentbeck.github.io/TestDesiderata)
- [Hexagonal Architecture](https://alistair.cockburn.us/hexagonal-architecture/)
