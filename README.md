# Go Kata

A collection of coding katas in Go for practicing Test-Driven Development (TDD), refactoring, and software design.

## Prerequisites

- Go 1.21 or later
- A code editor (VS Code recommended with Go extension)

## Installing Go

### macOS

**Option 1: Official Installer (Recommended)**

1. Download from [go.dev/dl](https://go.dev/dl/)
2. Run the `.pkg` installer
3. Restart your terminal

**Option 2: Homebrew**

```bash
brew install go
```

### Windows

**Option 1: Official Installer (Recommended)**

1. Download the `.msi` from [go.dev/dl](https://go.dev/dl/)
2. Run the installer
3. Restart your terminal/PowerShell

**Option 2: Chocolatey**

```bash
choco install golang
```

**Option 3: Scoop**

```bash
scoop install go
```

### Linux

**Option 1: Official Tarball (Recommended)**

```bash
# Download latest (replace version as needed)
wget https://go.dev/dl/go1.22.0.linux-amd64.tar.gz

# Remove old installation and extract
sudo rm -rf /usr/local/go
sudo tar -C /usr/local -xzf go1.22.0.linux-amd64.tar.gz

# Add to PATH (add to ~/.bashrc or ~/.zshrc)
export PATH=$PATH:/usr/local/go/bin
```

**Option 2: Package Manager**

```bash
# Ubuntu/Debian (may not be latest version)
sudo apt install golang-go

# Fedora
sudo dnf install golang
```

### Verify Installation

```bash
go version
# Should output: go version go1.2x.x ...
```

## Installing Development Tools

### golangci-lint (Linter)

> ⚠️ **Note**: The golangci-lint team recommends using binary installation or package managers over `go install` for consistent behavior.

**macOS**

```bash
brew install golangci-lint
```

**Windows**

```bash
# Chocolatey
choco install golangci-lint

# Scoop
scoop install golangci-lint
```

**Linux / macOS / Windows (Binary Install - Recommended for CI)**

```bash
# Install latest version to $(go env GOPATH)/bin
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

# Or install specific version
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.55.2
```

**Docker (Alternative)**

```bash
docker run --rm -v $(pwd):/app -w /app golangci/golangci-lint golangci-lint run
```

### Other Go Tools

```bash
# Install mutation testing tool
go install github.com/zimmski/go-mutesting/cmd/go-mutesting@latest

# Install watch mode tool (macOS/Linux)
brew install entr  # or: apt install entr
```

## Quick Start (First Time Setup)

```bash
# 1. Verify Go is installed
go version

# 3. Clone and enter the project
cd go-kata

# 4. Download dependencies
go mod download

# 5. Run all tests to verify setup
make test
# Or without make:
go test ./...

# 6. You're ready! Start with FizzBuzz
make fizz
```

## Available Commands

Run `make help` to see all available commands:

```bash
make help         # Show all commands
make test         # Run all tests
make watch        # Watch mode (requires: brew install entr)
make coverage     # Run tests with coverage
make coverage-html # Generate HTML coverage report
make lint         # Run linter (requires: brew install golangci-lint)
make fmt          # Format code
make ci           # Run all CI checks
make install-tools # Install dev tools
```

### Individual Kata Tests

```bash
make fizz       # FizzBuzz
make leap       # Leap Year
make stats      # Stats Calculator
make anagrams   # Anagrams
make fib        # Fibonacci
make stack      # Stack
make roman      # Roman Numerals
make prime      # Prime Factors
make tic        # Tic Tac Toe
make yahtzee    # Yahtzee
make mars       # Mars Rover
make tennis     # Tennis Refactoring
make rose       # Gilded Rose
make golf       # All Refactoring Golf
make smelly     # Smelly Tic Tac Toe
make copier     # Character Copier
make tac        # Tic Tac Toe (output)
make esa        # ESA Mars Rover
make cart       # Shopping Cart
make social     # Social Network
make katacombs  # Katacombs
make smellymars # Smelly Mars Rover
make smellycart # Smelly Shopping Cart
make smellyyahtzee # Smelly Yahtzee
```

### Refactoring Golf Holes

```bash
make golf1     # Hole 1
make golf2     # Hole 2
# ... through golf13
```

### Watch Mode for Individual Katas

```bash
make watch-fizz  # Watch FizzBuzz
make watch-rose  # Watch Gilded Rose
```

## Without Make

If you prefer not to use `make`, here are the raw Go commands:

```bash
# Run all tests
go test ./...

# Run specific kata
go test -v ./01_fizzbuzz/...

# Run with coverage
go test -cover ./...

# Format code
go fmt ./...

# Check for issues
go vet ./...
```

## Katas

| #   | Kata                                          | Description                               | Difficulty   |
| --- | ---------------------------------------------- | ------------------------------------------ | ------------ |
| 01  | [FizzBuzz](01_fizzbuzz/)                      | Classic TDD starter                       | Beginner     |
| 02  | [LeapYear](02_leapyear/)                      | Date validation logic                     | Beginner     |
| 03  | [StatsCalculator](03_statscalculator/)        | Statistical calculations on a sequence    | Beginner     |
| 04  | [Anagrams](04_anagrams/)                      | Anagram generation                        | Beginner     |
| 05  | [Fibonacci](05_fibonacci/)                    | Sequence generation                       | Beginner     |
| 06  | [Stack](06_stack/)                            | Data structure implementation             | Beginner     |
| 07  | [RomanNumerals](07_romannumerals/)            | Number conversion                         | Intermediate |
| 08  | [PrimeFactors](08_primefactors/)              | Mathematical decomposition                | Intermediate |
| 09  | [TicTacToe](09_tictactoe/)                    | Game logic with OOP                       | Intermediate |
| 10  | [Yahtzee](10_yahtzee/)                        | Scoring system                            | Intermediate |
| 11  | [MarsRover](11_marsrover/)                    | Command pattern & state                   | Advanced     |
| 12  | [Tennis](12_tennis/)                          | Refactoring kata                          | Advanced     |
| 13  | [GildedRose](13_gildedrose/)                  | Legacy code & testing                     | Advanced     |
| 14  | [RefactoringGolf](14_refactoringgolf/)        | Progressive refactoring holes             | Advanced     |
| 15  | [SmellyTicTacToe](15_smellytictactoe/)        | Code smells identification                | Advanced     |
| 16  | [CharacterCopier](16_charactercopier/)        | Mocking & test doubles                    | Intermediate |
| 17  | [TicTacToe (Output)](17_tictactoe/)           | Game with output interface                | Intermediate |
| 18  | [ESA MarsRover](18_esamarsrover/)             | Multilingual rover with radio             | Advanced     |
| 19  | [ShoppingCart](19_shoppingcart/)              | DDD & Hexagonal architecture              | Advanced     |
| 20  | [SocialNetwork](20_socialnetwork/)            | Use case driven development               | Advanced     |
| 21  | [Katacombs](21_katacombs/)                    | Text adventure game                       | Advanced     |
| 22  | [SmellyMarsRover](22_smellymarsrover/)        | Legacy code & test smells fixture         | Advanced     |
| 23  | [SmellyShoppingCart](23_smellyshoppingcart/)  | Hexagonal ports & adapters smells fixture | Advanced     |
| 25  | [SmellyYahtzee](25_smellyyahtzee/)            | Legacy code & test smells fixture (dice)  | Advanced     |

## TDD Rules

1. ✅ Write production code only to pass a failing unit test
2. ✅ Write only enough of a unit test to make it fail
3. ✅ Write only enough production code to make the failing test pass

## Running Mutation Tests

```bash
# Run mutation tests (requires go-mutesting, see Installing Development Tools)
go-mutesting ./...
```

## Project Structure

```
go-kata/
├── go.mod
├── go.sum
├── README.md
├── 01_fizzbuzz/
│   ├── fizzbuzz.go
│   ├── fizzbuzz_test.go
│   └── README.md
├── 02_leapyear/
│   └── ...
└── ...
```

## Resources

- [TestDesiderata by Kent Beck](https://kentbeck.github.io/TestDesiderata)
- [Object Calisthenics](https://williamdurand.fr/2013/06/03/object-calisthenics/)
- [Transformation Priority Premise](https://blog.cleancoder.com/uncle-bob/2013/05/27/TheTransformationPriorityPremise.html)
