
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

#### **Things missing**

- Replace pixels by images on the snake and the food:
    - Snake head
    - Snake tail
    - Snake body:
        - when down to right: `curve_1.png`
        - when down to left: `curve_2.png`
        - when up to right: `curve_4.png`
        - when up to left: `curve_3.png`
        - when right to up: `curve_2.png`
        - when right to down: `curve_3.png`
        - when left to up: `curve_1.png`
        - when left to down: `curve_4.png`
        - else continue straight: `straight.png`
    - Food
    _Note_: To make this happen we might need to had a new object to the Snake struct `Edges` which will be a Map that to a coordinate will have a string that will be the image path. In case the coordinate is no longer part of the snake body then we remove it from the map.