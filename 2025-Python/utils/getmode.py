def get_mode(args: list[str]) -> str:
    file_path = "input.txt"
    if len(args) > 1:
        if args[1] == "test":
            file_path = "test.txt"
    return file_path
