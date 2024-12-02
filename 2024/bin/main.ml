let () =
  let day = Sys.argv.(1) in
  match day with
  | "0" -> Day00.run ()
  | "1" -> Day01.run ()
  | "2" -> Day02.run ()
  | _ -> Printf.printf "Invalid day: %s\n" day
