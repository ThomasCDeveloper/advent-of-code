open Base
open Stdio

let read_file filename = In_channel.read_all filename
let part1 lines = 0
let part2 lines = 0
let filename = "input/day00.tst"

let () =
  let content = In_channel.read_all filename in
  let lines = String.split_lines content in
  print_string ("Solution to part 1: " ^ Int.to_string (part1 lines) ^ "\n");
  print_string ("Solution to part 2: " ^ Int.to_string (part2 lines) ^ "\n")
