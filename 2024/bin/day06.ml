open Base
open Stdio

let filename = "inputs/day06.txt"

let parse_input lines =
  let walls = ref [] in
  let start = ref (0, 0) in
  let height = Array.length lines - 1 in
  let width = String.length lines.(0) - 1 in
  for y = 0 to height do
    for x = 0 to width do
      match String.sub lines.(y) ~pos:x ~len:1 with
      | "#" -> walls := (x, y) :: !walls
      | "^" -> start := (x, y)
      | _ -> walls := !walls
    done
  done;
  (!walls, !start, width, height)

let rec is_wall pos walls =
  match walls with
  | [] -> false
  | hd :: rest ->
      let hx, hy = hd in
      let px, py = pos in
      if px = hx && py = hy then true else is_wall pos rest

let turn_right dir =
  let dx, dy = dir in
  (-dy, dx)

let part1 walls start width height =
  let current_dir = (0, -1) in
  let pos = start in
  let rec walk pos current_dir walls seen =
    let x, y = pos in
    if x < 0 || y < 0 || x >= width || y >= height then List.length seen + 1
    else
      let dx, dy = current_dir in
      let new_pos = (x + dx, y + dy) in
      if is_wall new_pos walls then walk pos (turn_right current_dir) walls seen
      else
        walk new_pos current_dir walls
          (if is_wall pos seen then seen else pos :: seen)
  in
  print_string (Int.to_string (walk pos current_dir walls []) ^ "\n")

let rec been_there status visited =
  match visited with
  | [] -> false
  | hd :: rest ->
      let hx, hy, hdx, hdy = hd in
      let px, py, pdx, pdy = status in
      if px = hx && py = hy && pdx = hdx && pdy = hdy then true
      else been_there status rest

let rec check_is_loop pos dir walls visited width height =
  let x, y = pos in
  let dx, dy = dir in
  if been_there (x, y, dx, dy) visited then 1
  else if x < 0 || y < 0 || x >= width || y >= height then 0
  else
    let new_pos = (x + dx, y + dy) in
    if is_wall new_pos walls then
      check_is_loop pos (turn_right dir) walls visited width height
    else
      check_is_loop new_pos dir walls ((x, y, dx, dy) :: visited) width height

let part2 walls start width height =
  let current_dir = (0, -1) in
  let pos = start in
  let px, py = pos in
  let loops = ref 0 in
  for zy = 0 to height do
    for zx = 0 to width do
      if (not (phys_equal zx px)) || not (phys_equal zy py) then
        let new_walls = (zx, zy) :: walls in
        loops :=
          !loops + check_is_loop pos current_dir new_walls [] width height
    done
  done;
  print_string (Int.to_string !loops ^ "\n")

let run () =
  let content = In_channel.read_all filename in
  let lines = Array.of_list (String.split_lines content) in
  let walls, start, width, height = parse_input lines in
  part1 walls start width height;
  part2 walls start width height
