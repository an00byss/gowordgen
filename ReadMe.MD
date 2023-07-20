# Gowordgen

 A quick and (super) dirty wordlist generator / mangler that takes a list of keywords and appends them with years, songs, seasons, nfl, nba, and nhl teams.


## Features

- Import a list of words and appends sports teams, songs and years to the wordlist.
- Leet mode to change up special characters
- It's quick

### Usage

```
# Normal
go run main.go -i .\keywords.txt

# Leet Mode
go run main.go -i .\keywords.txt -l yes
```