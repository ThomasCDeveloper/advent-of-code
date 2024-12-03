let read_file filename =
  let ic = open_in filename in
  try
    let rec read_lines acc =
      match input_line ic with
      | line -> read_lines (line :: acc)
      | exception End_of_file -> acc
    in
    let lines = read_lines [] in
    close_in ic;
    lines
  with e ->
    close_in_noerr ic;
    raise e

let extract_lefts lines len =
  List.map (fun x -> int_of_string (String.sub x 0 len)) lines

let extract_rights lines len =
  List.map (fun x -> int_of_string (String.sub x (len + 3) len)) lines

let part1 lefts rights =
  let differences =
    List.map2 (fun left right -> abs (left - right)) lefts rights
  in
  List.fold_left ( + ) 0 differences

let count_occurences number lst =
  List.fold_left (fun acc x -> if x = number then acc + 1 else acc) 0 lst

let part2 lefts rights =
  List.fold_left
    (fun similarity element ->
      similarity + (element * count_occurences element rights))
    0 lefts

let filename = "input.txt"
let length = match filename with "input.txt" -> 5 | _ -> 1

let run () =
  let lines = read_file filename in
  let lefts = List.sort Int.compare (extract_lefts lines length) in
  let rights = List.sort Int.compare (extract_rights lines length) in
  print_string
    ("Solution to part 1: " ^ string_of_int (part1 lefts rights) ^ "\n");
  print_string
    ("Solution to part 2: " ^ string_of_int (part2 lefts rights) ^ "\n")
