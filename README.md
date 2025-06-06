# taskctl

```
████████╗ █████╗ ███████╗██╗  ██╗ ██████╗████████╗██╗     
╚══██╔══╝██╔══██╗██╔════╝██║ ██╔╝██╔════╝╚══██╔══╝██║     
   ██║   ███████║███████╗█████╔╝ ██║        ██║   ██║     
   ██║   ██╔══██║╚════██║██╔═██╗ ██║        ██║   ██║     
   ██║   ██║  ██║███████║██║  ██╗╚██████╗   ██║   ███████╗
   ╚═╝   ╚═╝  ╚═╝╚══════╝╚═╝  ╚═╝ ╚═════╝   ╚═╝   ╚══════╝
       A simple CLI Task Manager (taskctl)
```

---

## 📌 What is taskctl?

taskctl is a lightweight, cross-platform **command-line task manager** written in Go. It allows you to quickly create, view, update, delete, and organize your to-do list with ease — all in the terminal.

---

## ✨ Features

* ✅ Add tasks with a **priority** and optional **due date**
* ✅ Mark tasks as **done / not done**
* ✅ Sort and view tasks:

  * by **priority**
  * by **completion status**
  * by **due date**
* ✅ Color-coded output for task status
* ✅ Persist tasks in `tasks.json`
* ✅ Clear all tasks
* ✅ Polished interactive UI via `promptui`
* ✅ Easy keyboard navigation

---

## Requirements

* [Go 1.20+](https://golang.org/dl/)
* A terminal (Windows Terminal, macOS Terminal, Linux shell)

---

## Getting Started

Windows Beeping Warning

If you're using PowerShell or Command Prompt and hear annoying beeps when navigating menus, this is due to limited keyboard support in legacy Windows terminals.

It is recommended to run taskctl via the VS Code integrated terminal

### 1. Clone the repo

```bash
git clone https://github.com/yourusername/taskctl.git
cd taskctl
```

### 2. Build the app

```bash
go build -o taskctl
```

### 3. Run it

```bash
./taskctl
```

> On Windows, run `taskctl.exe` from a terminal (best with VS Code terminal)

---

## File Structure

```
taskctl/
├── main.go         # Main source code
├── go.mod          # Module and dependency definitions
├── go.sum          # Dependency checksums
├── tasks.json      # Auto-generated task data store
└── README.md       # You're reading it!
```

---

## Sample Usage

```text
████████╗ █████╗ ███████╗██╗  ██╗ ██████╗████████╗██╗     
╚══██╔══╝██╔══██╗██╔════╝██║ ██╔╝██╔════╝╚══██╔══╝██║     
   ██║   ███████║███████╗█████╔╝ ██║        ██║   ██║     
   ██║   ██╔══██║╚════██║██╔═██╗ ██║        ██║   ██║     
   ██║   ██║  ██║███████║██║  ██╗╚██████╗   ██║   ███████╗
   ╚═╝   ╚═╝  ╚═╝╚══════╝╚═╝  ╚═╝ ╚═════╝   ╚═╝   ╚══════╝
           A simple CLI Task Manager (taskctl)

Use the arrow keys to navigate: ↓ ↑ → ←
? Welcome to the Task Manager:
  > Create | Read | Update | Delete Tasks
    View Tasks
    Clear Tasks
    Exit
```

You can:

* Create a new task: `Buy milk`, priority `high`, due `2025-06-10`
* View all tasks by priority or due date
* Mark it complete and come back later

---

## Data Persistence

All your tasks are stored in a local file:

```bash
tasks.json
```

This file is created automatically when you first use the app. No database required!

---

## Future Ideas

- [ ] Introduce key-based option choices
- [ ] Cron job integration for scheduled task reminders
- [ ] Sync tasks in the cloud using storage providers
- [ ] Export/import to and from `.csv`, `.json`, or `.md`
- [ ] Command-line flags for task creation
- [ ] Migrate to a more Windows-friendly TUI like `bubbletea`


## Challenges & Known Issues

Console Beeping on Windows

* Some terminals like PowerShell or cmd.exe trigger beeping noises when using arrow keys.

This is caused by how raw input is handled in these shells. It is recommended to use taskctl in a VS Code terminal.

---

## License

MIT License — use it, modify it, share it.

---

> *taskctl is your personal productivity assistant, right in the terminal.*
