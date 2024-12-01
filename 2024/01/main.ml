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
  with
  | e ->
      close_in_noerr ic;
      raise e

let rec print_list list =
  List.iter (fun x -> print_string ((string_of_int x) ^ "\n")) list;
  print_newline ()

let extract_lefts lines len =
  let rec extract_lefts_acc acc = function
    | [] -> acc
    | x :: xs ->
        let left = int_of_string (String.sub x 0 len) in
        extract_lefts_acc (left :: acc) xs
  in
  extract_lefts_acc [] lines

let extract_rights lines len =
  let rec extract_rights_acc acc = function
    | [] -> acc
    | x :: xs ->
        let right = int_of_string (String.sub x (len+3) len) in
        extract_rights_acc (right :: acc) xs
  in
  extract_rights_acc [] lines

let part1 lefts rights =
  let differences = List.map2 (fun left right -> abs (left - right)) lefts rights in
  List.fold_left ( + ) 0 differences

let count_occurences number list =
  List.fold_left (fun acc x -> if x = number then acc + 1 else acc) 0 list

let part2 lefts rights =
  List.fold_left(fun similarity element -> similarity + element * count_occurences element rights) 0 lefts

let filename = "input.txt"
let length =
  match filename with
  | "input.txt" -> 5
  | _ -> 1


let () =
  let lines = read_file filename in
  let lefts = List.sort Int.compare (extract_lefts lines length) in
  let rights = List.sort Int.compare (extract_rights lines length) in
  print_string ("Solution to part 1: " ^ string_of_int (part1 lefts rights) ^ "\n");
  print_string ("Solution to part 2: " ^ string_of_int (part2 lefts rights) ^ "\n")