let () =
  let day = Sys.argv.(1) in
  match day with
  | "0" -> Day00.run ()
  | "1" -> Day01.run ()
  | "2" -> Day02.run ()
  | "3" -> Day03.run ()
  | "4" -> Day04.run ()
  | "5" -> Day05.run ()
  | "6" -> Day06.run ()
  | "7" -> Day07.run ()
  | "8" -> Day08.run ()
  | "9" -> Day09.run ()
  | "10" -> ()
  | "14" -> Day14.run ()
  | _ -> Printf.printf "Invalid day: %s\n" day
