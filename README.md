<h1 align="center">TicTacGo</h1>

<br>

<p align="center">
  <img src="./gifs/basic.gif" />
</p>

<br>

## Install

```shell
go get github.com/josephthomashines/TicTacGo

cd ~/go/src/github.com/josephthomashines/TicTacGo

go build

# Run in interactive mode
./TicTacGo

# Run n simulations
./TicTacGo n
```

## Results

On a 12" Macbook (2016)

```
Model Name: MacBook
Model Identifier: MacBook9,1
Processor Name: Intel Core m3
Processor Speed: 1.1 GHz
Number of Processors: 1
Total Number of Cores: 2
L2 Cache (per Core): 256 KB
L3 Cache: 4 MB
Memory: 8 GB
```

Playing 1,000,000 games

```
TicTacGo: Simulated 1000000 Games (played randomly, X goes first)

 Took: 43.570393334s

Totals:
  X wins: 584971
  O wins: 288443
  Draws: 126586
  Plays: 7626199
Rates:
  X winrate: 58.50%
  O winrate: 28.84%
  Draw rate: 12.66%
  Plays per game: 7.63
```

## Checklist

| Status |                Feature                |
| :----: | :-----------------------------------: |
|   ✅   |   Allow for CLI input to play game    |
|   ✅   |     Error checking for user input     |
|   ✅   |          Add simulation mode          |
|   ✅   | Use goroutines to speed up simulation |
|   ❌   |            Add unit tests             |
|   ❌   |         Basic TravisCI setup          |
