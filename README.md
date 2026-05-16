# 🍷 Winebox

Winebox is a lightweight CLI tool to manage Wine prefixes and applications in a clean, structured way.

It provides a simple registry system per prefix and a unified way to run, organize, and inspect Windows applications under Wine.

---

## ✨ Features

- Create and manage Wine prefixes
- Register applications per prefix
- Define a main application per prefix
- Run applications through a unified interface
- Tree-style listing (lsblk-inspired)
- Optional file size display per prefix
- Optional executable path display
- Safe name sanitization system

---

## 📦 Installation

### Build from source

```bash
git clone https://github.com/aguadecoco1301/winebox
cd winebox
go build -o winebox
```

Move binary to your PATH:

```bash
sudo mv winebox /usr/local/bin/
```

---

## 🚀 Usage

### Create a prefix

```bash
winebox create <path-name>
```

### Open a built-in shell to install app if needed

```bash
winebox shell <path-name>
```

### Add a sub-application

```bash
winebox app add <prefix-name> <app-name> /path/to/app.exe
```

### Set main application

```bash
winebox app main <prefix-name> /path/to/app.exe
```

### Run application

```bash
winebox run <prefix-name>
winebox run <prefix-name> <app-name>
```

### List prefixes and apps

```bash
winebox list
winebox list --size --paths
winebox list --prefixes
```

---

## 📁 Data structure

Winebox stores data per prefix:

```
~/.winebox/
  prefixes/
    gta-sa/
      apps.json
```

---

## 🛣️  Roadmap

- v1.1
  - Sorting (name / size)
  - UX improvements in listing
  - Create .desktop files
- v1.2
  - Caching of prefix sizes
  - Performance improvements
  - Colors
  - Shell completion
- v1.3
  - Fuzzy search for run
  - Improved visual UI mode

---

## 👨‍💻 Author

Adriel Ulloa (AguaDeCoco1301)
