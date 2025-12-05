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
    List.rev lines
  with e ->
    close_in_noerr ic;
    raise e

let get_numbers_from_line line =
  List.map int_of_string (String.split_on_char ' ' line)

let is_acceptable_value value =
  match value with 1 | 2 | 3 -> true | _ -> false

let evaluate_numbers numbers =
  let is_ascending =
    List.fold_left2
      (fun acc normal sorted -> acc && normal = sorted)
      true numbers
      (List.sort compare numbers)
  in
  let is_descending =
    List.fold_left2
      (fun acc normal sorted -> acc && normal = sorted)
      true numbers
      (List.rev (List.sort compare numbers))
  in
  let is_gradual =
    let differences lst =
      let rec compute_differences acc lst =
        match lst with
        | a :: (b :: _ as rest) -> compute_differences ((b - a) :: acc) rest
        | _ -> acc
      in
      compute_differences [] lst
    in
    List.fold_left
      (fun acc diff -> acc && is_acceptable_value diff)
      true
      (differences (List.sort compare numbers))
  in
  (is_ascending || is_descending) && is_gradual

let all_lists_without_one_element lst =
  let rec aux i acc =
    match lst with
    | [] -> []
    | _ ->
        if i < List.length lst then
          let new_list = List.filteri (fun j _ -> j <> i) lst in
          aux (i + 1) (new_list :: acc)
        else List.rev acc
  in
  aux 0 []

let evaluate_numbers2 numbers =
  List.fold_left
    (fun acc sub_numbers -> acc || evaluate_numbers sub_numbers)
    false
    (all_lists_without_one_element numbers)

let part1 lines =
  let all_numbers = List.map get_numbers_from_line lines in
  List.fold_left
    (fun acc numbers -> if evaluate_numbers numbers then acc + 1 else acc)
    0 all_numbers

let part2 lines =
  let all_numbers = List.map get_numbers_from_line lines in
  List.fold_left
    (fun acc numbers -> if evaluate_numbers2 numbers then acc + 1 else acc)
    0 all_numbers

let filename = "input.txt"

let run () =
  let lines = read_file filename in
  print_string ("Solution to part 1: " ^ string_of_int (part1 lines) ^ "\n");
  print_string ("Solution to part 2: " ^ string_of_int (part2 lines) ^ "\n")
