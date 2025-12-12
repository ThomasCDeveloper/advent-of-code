open Base
open Stdio

let filename = "inputs/day09.txt"

let constitute_string seq =
  let rec c_s current remaining is_zero i =
    match remaining with
    | [] -> printf "%s\n" current
    | (_ as num) :: rem ->
        let dsk =
          Int.to_string num ^ "x:"
          ^ (if is_zero then "." else Int.to_string num)
          ^ " "
        in
        c_s (current ^ dsk) rem (not is_zero) (if is_zero then i + 1 else i)
  in
  c_s "" seq false 0

let part1 input =
  constitute_string input;
  printf "Part 1: %d\n" @@ List.length input

let part2 input = printf "Part 2: %d\n" @@ List.length input

let run () =
  let input =
    In_channel.read_all filename
    |> String.to_list
    |> List.map ~f:(fun x -> Int.of_string (String.of_char x))
  in
  part1 input;
  part2 []
