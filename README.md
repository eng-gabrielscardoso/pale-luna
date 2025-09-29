# Pale Luna

A recreation of the legendary "Pale Luna" creepypasta game - a mysterious text-based horror experience that supposedly appeared on an old computer with no known creator.

## The Legend

According to the creepypasta, "Pale Luna" was discovered on an abandoned computer. The game consisted of simple text commands and responses, but players reported strange occurrences when playing at specific times, particularly around 3:00 AM.

## Features

- **Time-sensitive gameplay**: The game behaves differently depending on when you play
- **Interactive text commands**: Type various commands to explore the game world
- **Atmospheric horror**: Experience the eerie ambiance of the original story
- **Session tracking**: The game remembers how many times you've played
- **Hidden commands**: Discover secret interactions through experimentation

## How to Play

1. **Build and run the game**:

```bash
go run main.go
```

2. **Basic commands**:

- `help` - Show available commands
- `time` - Check the current time
- `status` - View your current game status
- `pale luna` - The most important command (try it at different times)
- `quit` - Exit the game

3. **Tips**:
- Try playing at different hours of the day
- The game is most "active" during certain time periods
- Experiment with different word combinations
- Pay attention to the responses - they might change based on timing

## The 3 AM Rule

According to the legend, something special happens when you invoke "Pale Luna" during the 3:00 AM hour. The original players reported that the game would respond differently during this "witching hour"...

## Warning

This is a recreation for entertainment purposes. The original "Pale Luna" may or may not have existed, but the stories surrounding it have become part of internet folklore. Play at your own discretion, especially during late night hours.

## Technical Details

- Written in Go
- Terminal/console-based interface
- Cross-platform compatible (Windows, macOS, Linux)
- No external dependencies required

## Building

```bash
# Clone or download the repository
cd pale-luna

# Run directly
go run main.go

# Or build an executable
go build -o pale-luna main.go
./pale-luna
```

---

_"In the pale light of 3 AM, when the veil is thinnest, she waits for you to call her name..."_
