open Base
open Stdio

let filename = "inputs/day08.txt"

let parse_input filename =
  let dict = ref (Hashtbl.create (module String)) in
  let content = In_channel.read_all filename in
  let lines = Array.of_list (String.split_lines content) in
  let height = Array.length lines - 1 in
  let width = String.length lines.(0) - 1 in
  for y = 0 to height do
    for x = 0 to width do
      match String.sub lines.(y) ~pos:x ~len:1 with
      | "." -> ()
      | _ as freq -> (
          let current_list = Hashtbl.find !dict freq in
          match current_list with
          | None -> Hashtbl.set !dict ~key:freq ~data:[ (x, y) ]
          | Some lst -> Hashtbl.set !dict ~key:freq ~data:((x, y) :: lst))
    done
  done;
  (!dict, width, height)

let is_pos_in_list list pos =
  let rec ipil' list pos =
    match list with
    | [] -> false
    | (_ as p) :: rest ->
        let px, py = p in
        let x, y = pos in
        if px = x && py = y then true else ipil' rest pos
  in
  ipil' list pos

let add_pos_to_list list poss width height =
  List.fold_left ~init:list
    ~f:(fun acc pos ->
      let x, y = pos in
      if x < 0 || y < 0 || x > width || y > height then acc
      else if not (is_pos_in_list acc (x, y)) then (x, y) :: acc
      else acc)
    poss

let analyse_list_antennas_1 lst other w h =
  List.fold_left ~init:other
    ~f:(fun acc antenna1 ->
      let x1, y1 = antenna1 in
      List.fold_left ~init:acc
        ~f:(fun acc2 antenna2 ->
          let x2, y2 = antenna2 in
          if x1 = x2 && y1 = y2 then acc2
          else
            add_pos_to_list acc2
              [ ((2 * x1) - x2, (2 * y1) - y2); ((2 * x2) - x1, (2 * y2) - y1) ]
              w h)
        lst)
    lst

let part1 input width height =
  let nodes =
    List.fold_left ~init:[]
      ~f:(fun acc x ->
        let check = Hashtbl.find input x in
        let list_antennas = match check with None -> [] | Some lst -> lst in
        analyse_list_antennas_1 list_antennas acc width height)
      (Hashtbl.keys input)
  in
  printf "Part 1: %d\n" @@ List.length nodes

let get_list_of_pos x1 y1 x2 y2 =
  let lst = ref [] in
  let dx, dy = (x1 - x2, y1 - y2) in
  for i = 0 to 50 do
    lst := (x1 + (i * dx), y1 + (i * dy)) :: !lst
  done;
  for i = 0 to 50 do
    lst := (x2 - (i * dx), y2 - (i * dy)) :: !lst
  done;
  !lst

let analyse_list_antennas_2 lst other w h =
  List.fold_left ~init:other
    ~f:(fun acc antenna1 ->
      let x1, y1 = antenna1 in
      List.fold_left ~init:acc
        ~f:(fun acc2 antenna2 ->
          let x2, y2 = antenna2 in
          if x1 = x2 && y1 = y2 then acc2
          else add_pos_to_list acc2 (get_list_of_pos x1 y1 x2 y2) w h)
        lst)
    lst

let part2 input width height =
  let nodes =
    List.fold_left ~init:[]
      ~f:(fun acc x ->
        let check = Hashtbl.find input x in
        let list_antennas = match check with None -> [] | Some lst -> lst in
        analyse_list_antennas_2 list_antennas acc width height)
      (Hashtbl.keys input)
  in
  printf "Part 2: %d\n" @@ List.length nodes

let run () =
  let antennas, width, height = parse_input filename in
  part1 antennas width height;
  part2 antennas width height
