open Base
open Stdio

let filename = "inputs/day07.txt"

let parse_input filename =
  let content = In_channel.read_all filename in
  let lines = String.split_lines content in
  List.map
    ~f:(fun x ->
      let linestring = String.substr_replace_all x ~pattern:":" ~with_:"" in
      List.map ~f:(fun x -> Int.of_string x) (String.split linestring ~on:' '))
    lines

let check_line_1 line =
  let rec rec_check_1 target remaining current =
    match remaining with
    | [] -> if current = target then 1 else 0
    | (_ as number) :: rest ->
        rec_check_1 target rest (current * number)
        + rec_check_1 target rest (current + number)
  in
  match line with
  | [] -> 0
  | [ _ ] -> 0
  | (_ as tar) :: seq ->
      if not (phys_equal (rec_check_1 tar seq 0) 0) then tar else 0

let check_line_2 line =
  let rec rec_check_2 target remaining current =
    match remaining with
    | [] -> if current = target then 1 else 0
    | (_ as number) :: rest ->
        rec_check_2 target rest (current * number)
        + rec_check_2 target rest (current + number)
        + rec_check_2 target rest
            (Int.of_string (Int.to_string current ^ Int.to_string number))
  in
  match line with
  | [] -> 0
  | [ _ ] -> 0
  | (_ as tar) :: seq ->
      if not (phys_equal (rec_check_2 tar seq 0) 0) then tar else 0

let part1 input =
  printf "Part 1: %d\n"
  @@ List.fold_left ~f:(fun acc x -> acc + check_line_1 x) ~init:0 input

let part2 input =
  printf "Part 1: %d\n"
  @@ List.fold_left ~f:(fun acc x -> acc + check_line_2 x) ~init:0 input

let run () =
  let input = parse_input filename in
  part1 input;
  part2 input
