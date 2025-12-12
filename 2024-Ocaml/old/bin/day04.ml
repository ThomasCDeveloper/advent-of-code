open Base
open Stdio

let filename = "inputs/day04.txt"
let directions = [ (1, 0); (0, 1); (1, 1); (1, -1) ]

let get_char_at lines x y =
  if y < 0 || y >= Array.length lines || x < 0 || x >= String.length lines.(0)
  then ""
  else String.sub lines.(y) ~pos:x ~len:1

let get_word_at_and_dir lines x y direction len =
  let rec get_w lines x y direction remaining acc =
    match remaining with
    | 0 -> acc
    | _ ->
        let dx, dy = direction in
        get_w lines (x + dx) (y + dy) direction (remaining - 1)
          (acc ^ get_char_at lines x y)
  in
  get_w lines x y direction len ""

let check_word_1 word = match word with "XMAS" | "SAMX" -> 1 | _ -> 0
let check_word_2 word = match word with "MAS" | "SAM" -> 1 | _ -> 0

let part1 lines =
  let result = ref 0 in
  for x = 0 to String.length lines.(0) - 1 do
    for y = 0 to Array.length lines - 1 do
      result :=
        !result
        + List.fold_left
            ~f:(fun acc d ->
              acc + check_word_1 (get_word_at_and_dir lines x y d 4))
            ~init:0 directions
    done
  done;
  !result

let part2 lines =
  let result = ref 0 in
  for x = 1 to String.length lines.(0) - 2 do
    for y = 1 to Array.length lines - 2 do
      result :=
        !result
        +
        if
          check_word_2 (get_word_at_and_dir lines (x - 1) (y - 1) (1, 1) 3) = 1
          && check_word_2 (get_word_at_and_dir lines (x - 1) (y + 1) (1, -1) 3)
             = 1
        then 1
        else 0
    done
  done;
  !result

let run () =
  let content = In_channel.read_all filename in
  let lines = Array.of_list (String.split_lines content) in
  print_string ("Solution to part 1: " ^ Int.to_string (part1 lines) ^ "\n");
  print_string ("Solution to part 2: " ^ Int.to_string (part2 lines) ^ "\n")
