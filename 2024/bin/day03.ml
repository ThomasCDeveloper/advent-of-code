open Base
open Stdio

let filename = "inputs/day03.txt"

let mult rule =
  let numbers =
    String.split ~on:','
      (String.substr_replace_all
         (String.substr_replace_all rule ~pattern:")" ~with_:"")
         ~pattern:"mul(" ~with_:"")
  in
  List.fold_left ~f:(fun acc x -> acc * Int.of_string x) ~init:1 numbers

let filter1 = Re.Pcre.re {|mul\(\d+,\d+\)|} |> Re.compile
let extract_mults line filter = Re.matches filter line

let part1 line =
  let mults = extract_mults line filter1 in
  List.fold_left ~f:(fun acc x -> acc + mult x) ~init:0 mults

let filter2 = Re.Pcre.re {|do\(\)|don't\(\)|mul\(\d+,\d+\)|} |> Re.compile

let part2 line =
  let mults_and_dos = extract_mults line filter2 in
  let rec do_or_dont elements do_ acc =
    match elements with
    | [] -> acc
    | "do()" :: rest -> do_or_dont rest 1 acc
    | "don't()" :: rest -> do_or_dont rest 0 acc
    | (_ as mul) :: rest -> do_or_dont rest do_ (acc + (do_ * mult mul))
  in
  do_or_dont mults_and_dos 1 0

let run () =
  let line = In_channel.read_all filename in
  print_string ("Solution to part 1: " ^ Int.to_string (part1 line) ^ "\n");
  print_string ("Solution to part 2: " ^ Int.to_string (part2 line) ^ "\n")
