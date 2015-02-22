# grollo
Strategy game written in Go

## Objects

### Player
* id
* created at
* updated at
* email
* username
* password

### Player Profile
* type: Human, AI
* timezone
* first name
* last name
* avatar

### Game
* id
* name
* creator
* players
* map
* rules
* cards

### Game State
* state (created, started, completed)
* active player
* turn status
* possible moves

### Action
* id
* name

### Turn
* id
* player id
* game id
* move ids

### Move
* id
* player id
* action id
* 
