open Base
open Stdio

let filename = "inputs/day05.tst"

let is_a_before_b_in_c a b c =
  if c.(0) = a then if c.(1) = b then true else false else false

let get_rules_and_instruction lines =
  let rec get_rules_and_instruction' lines rules instructions =
    match lines with
    | [] -> (rules, instructions)
    | "" :: rest -> get_rules_and_instruction' rest rules instructions
    | (_ as line) :: rest ->
        if String.length line < 6 then
          get_rules_and_instruction' rest
            (String.split_on_chars ~on:[ '|' ] line :: rules)
            instructions
        else
          get_rules_and_instruction' rest rules
            (String.split_on_chars ~on:[ ',' ] line :: instructions)
  in
  get_rules_and_instruction' lines [] []

let check_instruction _instruction = 0

let part1 lines =
  let _rules, _instruction = get_rules_and_instruction lines in
  List.iter
    ~f:(fun x -> List.iter ~f:(fun y -> print_string (y ^ " ")) x)
    _rules;
  0

let run () =
  let content = In_channel.read_all filename in
  let lines = String.split_lines content in
  print_string ("Solution to part 1: " ^ Int.to_string (part1 lines) ^ "\n");
  print_string ("Solution to part 2: " ^ Int.to_string 0 ^ "\n")
