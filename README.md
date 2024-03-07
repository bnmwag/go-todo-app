# GO todo app

This is a simple todo app written in Go. It uses a simple .json file to store the todos.
Wanted to learn other languages and Go was the first one I wanted to learn thanks to it's simplicity and efficiency.
I also wanted a new way to keep track of my todos and thought it would be a good idea to combine the two "problems" into on solution.

## Getting started

1. Clone the repository
2. Run the following in the root of the repository

```bash
`go build && mv go-todo-app todo && sudo mv todo /usr/local/bin`
```

3. Create a file called `.todos.json` in your `/usr/local/bin` directory
4. Get started by running `todo add "Your first task"` in your terminal

```bash
todo add "Your first task"
```

## Usage

Commands:

| name    | arg 1           | arg 2     | description                                                                             |
| ------- | --------------- | --------- | --------------------------------------------------------------------------------------- |
| List    | `list`, `l`     |           | Lists all the tasks in the todo list <br /> `todo list`                                 |
| Add     | `add`, `a`      | `<task>`  | Adds a task with the name specified in the `<task>` value <br /> `todo add "Some task"` |
| Delete  | `delete`, `d`   | `<index>` | Removes a task with the index specified in the `<index>` value <br /> `todo delete 1`   |
| Check   | `check`, `c`    | `<index>` | Checks a task with the index specified in the `<index>` value <br /> `todo check 1`     |
| Uncheck | `uncheck`, `uc` | `<index>` | Unchecks a task with the index specified in the `<index>` value <br /> `todo uncheck 1` |
| Help    | `help`, `h`     |           | Shows the help menu                                                                     |
| Version | `version`, `v`  |           | Shows the version of the app                                                            |
