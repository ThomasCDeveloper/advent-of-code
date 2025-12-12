open Base
open Stdio

let read_file path =
  In_channel.read_all path

let parse input =
  String.split_lines input

let part_1 data = 
  List.length data

let part_2 data = 
  1+List.length data

let solve lines =
  print_string ("Part 1: " ^ Int.to_string (part_1 lines) ^ "\n");
  print_string ("Part 1: " ^ Int.to_string (part_2 lines) ^ "\n")

let () =
  let argv = Sys.get_argv () in
  let mode = argv.(1) in
  let day  = argv.(2) in

  let file =
    Stdlib.Filename.concat
      (Stdlib.Filename.concat "solutions" day)
      (mode ^ ".txt")
  in

  let content = read_file file in
  content |> parse |> solve
