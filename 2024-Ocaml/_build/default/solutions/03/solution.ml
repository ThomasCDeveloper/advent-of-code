open Base
open Stdio

let read_file filename =
  In_channel.read_lines filename

let () =
  let current_dir = Filename.dirname Sys.argv.(0) |> Stdlib.Filename.dirname in
  let filepath name = Stdlib.Filename.concat current_dir name in

  let file =
    if Array.length Sys.argv > 1 && String.(Sys.argv.(1) = "test") then
      filepath "test.txt"
    else
      filepath "input.txt"
  in

  let lines = read_file file in

  printf "Loaded %d lines from %s\n"
    (List.length lines) file;

  (* TODO: implement part 1 and part 2 *)
