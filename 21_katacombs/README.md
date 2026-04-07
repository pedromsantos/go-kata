# Katacombs

## Source

<https://github.com/conso/katacombs>

## Project Requirements

- Implement using Domain Driven Design and Hexagonal architecture
- Focus on use cases (transport layer implementation not required)
- Use in-memory database implementation

## Game Overview

A text-based adventure game where players:

- Explore interconnected locations
- Move using cardinal directions or special connections
- Collect treasures
- Find the katacomb exit to win
- Score based on collected treasure value

## Gameplay Guide

### Game Start

```text
LOST IN SHOREDITCH.
YOU ARE STANDING AT THE END OF BRICK LANE BEFORE A SMALL BRICK BUILDING CALLED THE OLD TRUMAN BREWERY.
AROUND YOU IS A FOREST OF RESTAURANTS.
A SMALL STREAM OF CRAFTED BEER FLOWS OUT OF THE BUILDING AND DOWN A GULLY.
>
```

### Available Commands

#### Movement

- `GO N/E/S/W` - Move in cardinal directions
- `GO UP/DOWN` - Use stairs
- `LOOK [direction/item]` - Examine surroundings or items

#### Interaction

- `OPEN [item]` - Open doors, gates, etc.
- `CLOSE [item]` - Close doors, gates, etc.
- `TAKE [item]` - Collect items (max 10 in bag)
- `DROP [item]` - Leave items
- `BAG` - View inventory
- `USE [item]` - Interact with items in specific locations

#### System Commands

- `?` - Show help menu
- `QUIT` - End game

### Game Mechanics

#### Items & Inventory

- Bag limit: 10 items
- Items can be taken and dropped anywhere
- Location descriptions list available items

#### Gold Collection

- Automatic collection when first visiting a location with gold
- Automatic collection when first opening an item containing gold
- Total gold viewable in bag

#### World Rules

1. Each location must have a unique title
2. Locations must have matching reverse connections
   - If going South from A leads to B, going North from B must lead to A

## Features (Gherkin Scenarios)

### Look and Move

```gherkin
Scenario: New Katacombs Player
  Given I'm a new player of Katacombs
  When I register to play the game
  Then I'm listed as a player in the players list
  And I'm shown the description of the entry of Katacombs

Scenario: Look in a direction with a path
  Given I'm playing Katacombs
  When I look in a direction
  And there is a path to that direction
  Then I'm shown a description of the path

Scenario: Move in a direction with a path
  Given I'm playing Katacombs
  When I move in a direction with a path
  Then I'm shown a description of the new location
```

### Pick and Drop Items

```gherkin
Scenario: Take an item from a location
  Given I'm playing Katacombs
  When I take an item that exists in the location
  Then I'm shown a message "Took %s"
  And the item is in my bag

Scenario: Drop an item from the bag
  Given I'm playing Katacombs
  When I drop an item that exists in my bag
  Then I'm shown a message "Dropped %s"
  And the item is no longer in my bag
```

### Using Items

```gherkin
Scenario: Use a key to open a locked door
  Given I'm in a location with a locked door
  And I have the key in my bag
  When I use the key
  Then the door opens
```

## Follow TDD Rules Strictly

1. ✅ Write production code only to pass a failing unit test
2. ✅ Write only enough of a unit test to make it fail
3. ✅ Write only enough production code to make the failing test pass

## Resources

- [MUD](https://en.wikipedia.org/wiki/Multi-user_dungeon)
- [How to program a text adventure in C](https://helderman.github.io/htpataic/htpataic01.html)
- [TestDesiderata by Kent Beck](https://kentbeck.github.io/TestDesiderata)
