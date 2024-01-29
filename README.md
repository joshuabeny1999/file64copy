# file64copy
[![Actions Status](https://github.com/joshuabeny1999/file64copy/workflows/Test/badge.svg)](https://github.com/joshuabeny1999/file64copy/actions)

This is a command-line program written in Go that encodes a file to Base64 and generates a command to decode it. It also provides an option to copy the output to the clipboard.

```bash
Usage: file64copy [OPTIONS] FILE
Encode a file to Base64 and generate a command to decode it.

Options:
  -c, --clipboard
        Copy output to clipboard
  -h, --help
        Display help information
  -v, --version
        Display version information
```

To paste the file back, you can use the command generated by `file64copy`. This command will run on any linux and macos system without the need to install any additional software.

This is an example of the command generated by `file64copy`:
```bash
echo "SGVsbG8gV29ybGQK" | openssl base64 -d -A -out Hello.txt
```

You can share the command with anyone and they can use it to decode the file. 

Helpful to share files with people over messaging apps and email preserving the content of the file and simple copy paste. You could also use it yourself to copy files locally on an ubuntu server per ssh via your clipboard. Just use `file64copy -c` and then paste the command into your ssh session and run it. So you have the file on your server. 
## Installation

You can download the latest version of `file64copy` from the [releases page](https://github.com/joshuabeny1999/file64copy/releases).

### Debian/Ubuntu/Linux Mint

Download the `.deb` package and install it using `dpkg`:

```bash
wget https://github.com/joshuabeny1999/file64copy/releases/latest/download/file64copy.deb
sudo dpkg -i file64copy.deb
```

### Fedora
Download the .rpm package and install it using dnf:

```bash
wget https://github.com/joshuabeny1999/file64copy/releases/latest/download/file64copy.rpm
sudo dnf install file64copy.rpm
```

### Arch Linux
Download the .pkg.tar.xz package and install it using pacman:
```bash
wget https://github.com/joshuabeny1999/file64copy/releases/latest/download/file64copy.pkg.tar.zst
sudo pacman -U file64copy.pkg.tar.zst
```
