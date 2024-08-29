
<h1 align="center"> 
  <strong>Snake</strong>
</h1>

<p align="center">
    <img src="./.docs/game_logo.jpeg" width="100" height="100">
</p>

<div align="center">

  <a href="go version">![go_version](https://img.shields.io/badge/go-1.22.5-blue)</a>
  <a href="code coverage">![coverage](https://img.shields.io/badge/coverage-95.77%25-green)</a>
  <a href="tests">![tests](https://img.shields.io/badge/tests-26%20passed%2C%200%20failed-brightgreen)</a>

</div>

#### **Description**

This is a simple snake game made in Golang using the ebitengine library.

#### **How to play**

- Use the arrow keys to move the snake.
- Eat the food to grow.
- Don't hit the walls or yourself.

#### **Play**

To play the game, run the following command:

```
go run cmp/snake/main.go
```

<p align="center">
    <img src="./.docs/gameplay.png" width="400" height="400">
</p>

#### **Acknowledgments**

There's a few things that I want to highlight:
- Sometimes the piece allocated to the snake's body is not really accurate considering the movement of the snake.
That probably has to do with the fact that the game clock and goroutine that handles the snake's movement are not really sync.
