Requirements:

Build a simple web server that tracks how many games players have won.

- `GET /players/{name}` should return a number indicating the total number of wins
- `POST /players{name}` should record a win for that name, incrementing every subsequent `POST` call.
