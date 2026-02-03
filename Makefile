# Go Kata Makefile
# Run 'make help' to see available commands

.PHONY: help test watch coverage lint fmt vet build clean install-tools mutants
.PHONY: fizz leap fib stack roman prime tic yahtzee mars tennis rose golf smelly
.PHONY: copier tac esa cart social katacombs
.PHONY: golf1 golf2 golf3 golf4 golf5 golf6 golf7 golf8 golf9 golf10 golf11 golf12 golf13

# Default target
.DEFAULT_GOAL := help

# Colors for help
CYAN := \033[36m
RESET := \033[0m

help: ## Show this help message
	@echo "Go Kata - Available Commands"
	@echo ""
	@echo "Usage: make <command>"
	@echo ""
	@awk 'BEGIN {FS = ":.*##"} /^[a-zA-Z0-9_-]+:.*##/ {printf "  $(CYAN)%-12s$(RESET) %s\n", $$1, $$2}' $(MAKEFILE_LIST)
	@echo ""
	@echo "Individual Kata Tests:"
	@echo "  $(CYAN)fizz$(RESET)        FizzBuzz"
	@echo "  $(CYAN)leap$(RESET)        Leap Year"
	@echo "  $(CYAN)fib$(RESET)         Fibonacci"
	@echo "  $(CYAN)stack$(RESET)       Stack"
	@echo "  $(CYAN)roman$(RESET)       Roman Numerals"
	@echo "  $(CYAN)prime$(RESET)       Prime Factors"
	@echo "  $(CYAN)tic$(RESET)         Tic Tac Toe (basic)"
	@echo "  $(CYAN)yahtzee$(RESET)     Yahtzee"
	@echo "  $(CYAN)mars$(RESET)        Mars Rover"
	@echo "  $(CYAN)tennis$(RESET)      Tennis Refactoring"
	@echo "  $(CYAN)rose$(RESET)        Gilded Rose"
	@echo "  $(CYAN)golf$(RESET)        All Refactoring Golf"
	@echo "  $(CYAN)smelly$(RESET)      Smelly Tic Tac Toe"
	@echo "  $(CYAN)copier$(RESET)      Character Copier"
	@echo "  $(CYAN)tac$(RESET)         Tic Tac Toe (with output)"
	@echo "  $(CYAN)esa$(RESET)         ESA Mars Rover"
	@echo "  $(CYAN)cart$(RESET)        Shopping Cart"
	@echo "  $(CYAN)social$(RESET)      Social Network"
	@echo "  $(CYAN)katacombs$(RESET)   Katacombs"
	@echo ""
	@echo "Refactoring Golf Holes:"
	@echo "  $(CYAN)golf1-13$(RESET)    Individual holes (e.g., make golf1)"

# ============================================================================
# Main Commands
# ============================================================================

test: ## Run all tests
	go test ./...

watch: ## Run tests in watch mode (requires entr)
	@command -v entr >/dev/null 2>&1 || { echo "Install entr: brew install entr"; exit 1; }
	@echo "Watching for changes... (Ctrl+C to stop)"
	@find . -name "*.go" | entr -c go test ./...

coverage: ## Run tests with coverage report
	go test -cover ./...

coverage-html: ## Generate HTML coverage report
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report: coverage.html"
	@open coverage.html 2>/dev/null || xdg-open coverage.html 2>/dev/null || echo "Open coverage.html in your browser"

lint: ## Run linter (requires golangci-lint)
	@command -v golangci-lint >/dev/null 2>&1 || { echo "Install: brew install golangci-lint"; exit 1; }
	golangci-lint run

fmt: ## Format all Go files
	go fmt ./...

vet: ## Run go vet
	go vet ./...

build: ## Build all packages
	go build ./...

clean: ## Clean build artifacts
	go clean ./...
	rm -f coverage.out coverage.html

install-tools: ## Install development tools
	@echo "Installing development tools..."
	go install github.com/zimmski/go-mutesting/cmd/go-mutesting@latest
	@echo ""
	@echo "Optional tools (install manually):"
	@echo "  brew install golangci-lint  # Linter"
	@echo "  brew install entr           # Watch mode"

mutants: ## Run mutation tests (requires go-mutesting)
	@command -v go-mutesting >/dev/null 2>&1 || { echo "Run 'make install-tools' first"; exit 1; }
	go-mutesting ./...

ci: ## Run all CI checks (fmt, vet, lint, test)
	@echo "Running CI checks..."
	go fmt ./...
	go vet ./...
	go build ./...
	@command -v golangci-lint >/dev/null 2>&1 && golangci-lint run || echo "Skipping lint (golangci-lint not installed)"
	go test ./...
	@echo "✅ All checks passed!"

# ============================================================================
# Individual Kata Tests
# ============================================================================

fizz: ## Test FizzBuzz
	go test -v ./01_fizzbuzz/...

leap: ## Test Leap Year
	go test -v ./02_leapyear/...

fib: ## Test Fibonacci
	go test -v ./03_fibonacci/...

stack: ## Test Stack
	go test -v ./04_stack/...

roman: ## Test Roman Numerals
	go test -v ./05_romannumerals/...

prime: ## Test Prime Factors
	go test -v ./06_primefactors/...

tic: ## Test Tic Tac Toe (basic)
	go test -v ./07_tictactoe/...

yahtzee: ## Test Yahtzee
	go test -v ./08_yahtzee/...

mars: ## Test Mars Rover
	go test -v ./09_marsrover/...

tennis: ## Test Tennis
	go test -v ./10_tennis/...

rose: ## Test Gilded Rose
	go test -v ./11_gildedrose/...

golf: ## Test all Refactoring Golf holes
	go test -v ./12_refactoringgolf/...

smelly: ## Test Smelly Tic Tac Toe
	go test -v ./13_smellytictactoe/...

copier: ## Test Character Copier
	go test -v ./14_charactercopier/...

tac: ## Test Tic Tac Toe (with output)
	go test -v ./15_tictactoe/...

esa: ## Test ESA Mars Rover
	go test -v ./16_esamarsrover/...

cart: ## Test Shopping Cart
	go test -v ./17_shoppingcart/...

social: ## Test Social Network
	go test -v ./18_socialnetwork/...

katacombs: ## Test Katacombs
	go test -v ./19_katacombs/...

# ============================================================================
# Refactoring Golf Individual Holes
# ============================================================================

golf1:
	go test -v ./12_refactoringgolf/hole1/...

golf2:
	go test -v ./12_refactoringgolf/hole2/...

golf3:
	go test -v ./12_refactoringgolf/hole3/...

golf4:
	go test -v ./12_refactoringgolf/hole4/...

golf5:
	go test -v ./12_refactoringgolf/hole5/...

golf6:
	go test -v ./12_refactoringgolf/hole6/...

golf7:
	go test -v ./12_refactoringgolf/hole7/...

golf8:
	go test -v ./12_refactoringgolf/hole8/...

golf9:
	go test -v ./12_refactoringgolf/hole9/...

golf10:
	go test -v ./12_refactoringgolf/hole10/...

golf11:
	go test -v ./12_refactoringgolf/hole11/...

golf12:
	go test -v ./12_refactoringgolf/hole12/...

golf13:
	go test -v ./12_refactoringgolf/hole13/...

# ============================================================================
# Watch Individual Katas (requires entr)
# ============================================================================

watch-fizz:
	@find ./01_fizzbuzz -name "*.go" | entr -c go test -v ./01_fizzbuzz/...

watch-leap:
	@find ./02_leapyear -name "*.go" | entr -c go test -v ./02_leapyear/...

watch-fib:
	@find ./03_fibonacci -name "*.go" | entr -c go test -v ./03_fibonacci/...

watch-rose:
	@find ./11_gildedrose -name "*.go" | entr -c go test -v ./11_gildedrose/...
