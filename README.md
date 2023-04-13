# Odyssey
Repo for solving Fly.io's Gossip Glomers distributed systems challenges.

## Purpose
I wanted to go about bettering skills in distributed systems and a set of good challenges arose from [fly.io](https://fly.io/dist-sys/).
I also wanted to better my Rust skills so my intention is to do these challenges in two parts - first with go and then 
with rust.

## Learnings

## Scratch Notes
- I believe that this is an iterative process and a system passing challenge 3 will still pass challenge 1. Therefore, I will
just create new branches and merge them into main as we go. However, if this does not hold, I will create and hold onto branches.

## How to Run
For the Go side, cd into Godyssey. Then you can build the go binary and then
execute as need be with the maelstrom terminal
- First lesson: `/maelstrom test -w echo --bin ./Godyssey --node-count 1 --time-limit 10`
- Second lesson: `/maelstrom test -w unique-ids --bin ./Godyssey --time-limit 30 --rate 1000 --node-count 3 --availability total --nemesis partition`