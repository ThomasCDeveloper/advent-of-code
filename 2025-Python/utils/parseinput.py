def parse_input_lines(file_path: str) -> list[str]:
    with open(file_path) as f:
        return [line.strip() for line in f.readlines()]


def parse_input(file_path: str) -> str:
    with open(file_path) as f:
        return f.read()


def parse_raw_lines(file_path: str) -> list[str]:
    with open(file_path) as f:
        return [line for line in f.readlines()]
