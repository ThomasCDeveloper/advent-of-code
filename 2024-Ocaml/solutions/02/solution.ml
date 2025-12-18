open Base
open Stdio

let read_file path =
  In_channel.read_all path

let parse input =
  String.split_lines input

let sign x =
  if x > 0 then 1 else -1

let is_safe levels =
  match levels with
  | [] | [_] -> true
  | a :: b :: rest ->
      let diff = b - a in
      if diff = 0 || Int.abs diff > 3 then
        false
      else
        let direction = sign diff in
        let rec loop prev = function
          | [] -> true
          | x :: xs ->
              let d = x - prev in
              if d = 0
                 || Int.abs d > 3
                 || sign d <> direction
              then false
              else loop x xs
        in
        loop b rest

let part_1 lines =
  lines
  |> List.map ~f:(fun line ->
       String.split line ~on:' '
       |> List.filter ~f:(Fn.non String.is_empty)
       |> List.map ~f:Int.of_string)
  |> List.count ~f:is_safe

let remove_at lst idx =
  List.filteri lst ~f:(fun i _ -> i <> idx)

let is_safe_with_dampener levels =
  is_safe levels
  || List.existsi levels ~f:(fun i _ ->
       remove_at levels i |> is_safe)

let part_2 lines =
  lines
  |> List.map ~f:(fun line ->
       String.split line ~on:' '
       |> List.filter ~f:(Fn.non String.is_empty)
       |> List.map ~f:Int.of_string)
  |> List.count ~f:is_safe_with_dampener

let solve lines =
  print_string ("Part 1: " ^ Int.to_string (part_1 lines) ^ "\n");
  print_string ("Part 2: " ^ Int.to_string (part_2 lines) ^ "\n")

let () =
  let argv = Sys.get_argv () in
  let mode = argv.(1) in
  let day  = argv.(2) in

  let file =
    Stdlib.Filename.concat
      (Stdlib.Filename.concat "solutions" day)
      (mode ^ ".txt")
  in
  file |> read_file |> parse |> solve
