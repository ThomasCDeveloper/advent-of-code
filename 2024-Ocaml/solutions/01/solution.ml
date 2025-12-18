open Base
open Stdio

let read_file path =
  In_channel.read_all path

let parse input =
  String.split_lines input

let part_1 lines =
  let left, right =
    List.fold lines ~init:([], []) ~f:(fun (l, r) line ->
      match String.split ~on:' ' line |> List.filter ~f:(Fn.non String.is_empty) with
      | [a; b] -> (Int.of_string a :: l, Int.of_string b :: r)
      | _ -> (l, r))
  in
  let left  = List.sort left  ~compare:Int.compare in
  let right = List.sort right ~compare:Int.compare in
  List.map2_exn left right ~f:(fun a b -> Int.abs (a - b))
  |> List.fold ~init:0 ~f:( + )

let part_2 lines =
  let left, right =
    List.fold lines ~init:([], []) ~f:(fun (l, r) line ->
      match String.split ~on:' ' line |> List.filter ~f:(Fn.non String.is_empty) with
      | [a; b] -> (Int.of_string a :: l, Int.of_string b :: r)
      | _ -> (l, r))
  in
  let freq =
    List.fold right ~init:(Map.empty (module Int)) ~f:(fun acc x ->
      Map.update acc x ~f:(function
        | None -> 1
        | Some n -> n + 1))
  in
  List.fold left ~init:0 ~f:(fun acc x ->
    let count = Map.find freq x |> Option.value ~default:0 in
    acc + (x * count))

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
