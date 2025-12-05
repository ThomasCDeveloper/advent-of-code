# Chameau

Props to AlbertCalmus

## Getting started

### Installation

Install Opam using [Homebrew](https://brew.sh/) and initialize it.

```bash
brew install opam
opam init -y
```

Run the following command (even better to add it in your `.zshrc` or equivalent)

```bash
eval $(opam env)
```

Install required libraries with `opam install`

```bash
opam install dune ocamlformat ocaml-lsp-server odoc core base merlin
```

### VSCode setup

Install the [OCaml Platform](https://marketplace.visualstudio.com/items?itemName=ocamllabs.ocaml-platform) extension.

Open your VSCode settings using `Cmd + shift + P > Open user Settings` and add the following:

```json
"[ocaml]": {
  "editor.tabSize": 2,
  "editor.rulers": [ 100 ],
  "editor.defaultFormatter": "ocamllabs.ocaml-platform",
  "editor.formatOnSave": true,
},
```

### Run the project

Open a terminal and run the following command to continuously build your OCaml files:

```bash
make watch
```

To run a specific day of the challenge, run the following command:

```bash
make run # prompts the user to select a day
make run DAY=1 # to skip the prompt and run day 1
```
