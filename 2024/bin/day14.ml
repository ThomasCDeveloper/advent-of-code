open Base
open Stdio

let filename = "inputs/day14.txt"
let max_x = if phys_equal filename "inputs/day14.tst" then 11 else 101
let max_y = if phys_equal filename "inputs/day14.tst" then 7 else 103
let med_x = (max_x - 1) / 2
let med_y = (max_y - 1) / 2

let parse_input lines =
  let out =
    List.map
      ~f:(fun x ->
        String.substr_replace_all
          (String.substr_replace_all x ~pattern:"p=" ~with_:"")
          ~pattern:" v=" ~with_:",")
      lines
  in
  List.map
    ~f:(fun x ->
      List.map
        ~f:(fun y -> Int.of_string y)
        (String.split_on_chars x ~on:[ ',' ])
      |> Array.of_list)
    out

let robot_move x y dx dy = ((x + dx + max_x) % max_x, (y + dy + max_y) % max_y)

let loop robots =
  let r =
    List.map
      ~f:(fun x ->
        let rx, ry, dx, dy = (x.(0), x.(1), x.(2), x.(3)) in
        let nrx, nry = robot_move rx ry dx dy in
        Array.of_list [ nrx; nry; dx; dy ])
      robots
  in
  r

let place_robot_in_quadrant r =
  let x, y = (r.(0), r.(1)) in
  if x = med_x || y = med_y then 0
  else if x < med_x && y < med_y then 1
  else if x < med_x && y > med_y then 3
  else if x > med_x && y < med_y then 2
  else 4

let evaluate_robots robots =
  List.fold
    ~f:(fun acc x ->
      let q1, q2, q3, q4 = acc in
      match place_robot_in_quadrant x with
      | 1 -> (q1 + 1, q2, q3, q4)
      | 2 -> (q1, q2 + 1, q3, q4)
      | 3 -> (q1, q2, q3 + 1, q4)
      | 4 -> (q1, q2, q3, q4 + 1)
      | _ -> (q1, q2, q3, q4))
    ~init:(0, 0, 0, 0) robots

let part1 input =
  let robots = ref input in
  for _ = 0 to 99 do
    robots := loop !robots
  done;
  let a, b, c, d = evaluate_robots !robots in
  printf "Part 1: %d\n" @@ (a * b * c * d)

let rec is_xy_in_robots x y robots =
  match robots with
  | [] -> false
  | (_ as rb) :: rem ->
      if rb.(0) = x && rb.(1) = y then true else is_xy_in_robots x y rem

let print_robots robots i =
  printf " === %d ===\n" @@ i;
  for y = 0 to max_y do
    for x = 0 to max_x do
      printf "%s" @@ if is_xy_in_robots x y robots then "x" else " "
    done;
    printf "\n"
  done

let evaluate_robots_2 robots =
  let robots1 = Array.of_list robots in
  let v = ref 0 in
  for i = 0 to Array.length robots1 - 1 do
    for j = 0 to Array.length robots1 - 1 do
      let r1, r2 = (robots1.(i), robots1.(j)) in
      v := !v + Int.abs (r1.(0) - r2.(0)) + Int.abs (r1.(1) - r2.(1))
    done
  done;
  !v

let part2 input =
  let robots = ref input in
  let min = ref (evaluate_robots_2 !robots) in
  let index_min = ref 0 in
  for i = 0 to 10000 do
    let eval = evaluate_robots_2 !robots in
    if eval < !min then index_min := i else ();
    (*if eval < !min then print_robots !robots i else ();*)
    if eval < !min then min := eval else ();
    robots := loop !robots
  done;
  printf "Part 2: %d\n" @@ !index_min

let run () =
  let content = In_channel.read_all filename in
  let lines = String.split_lines content in
  let input = parse_input lines in
  part1 input;
  part2 input
