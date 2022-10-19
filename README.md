cross-platform gui dialog program. base on <https://github.com/sqweek/dialog>

- show message
- select file
- select directory

# Installation

requirements:

- OSX: uses Cocoa's NSAlert/NSSavePanel/NSOpenPanel classes
- Win32: uses MessageBox/GetOpenFileName/GetSaveFileName (via package github.com/TheTitanrain/w32)
- Linux: uses Gtk's MessageDialog/FileChooserDialog (via cgo; requires gtk3 development packages)

  - Ubuntu

    ```bash
    sudo apt install libgtk-3-dev
    ```

install gui-dialog:

```
go install github.com/jiazhoulvke/gui-dialog@latest
```

# Usage

```
Usage of gui-dialog:
      --file_dialog_type string      load,save (default "load")
      --filter strings               file filter. example: jpg,png,gif
      --filter_desc string           file filter description. example: pictures
      --message_dialog_type string   info,error,yes_or_no (default "info")
  -m, --msg string                   message info
  -o, --output_type string           json,text (default "text")
  -d, --start_dir string             start directory
      --start_file string            start file
      --title string                 title
  -t, --type string                  file,dir,msg (default "file")
```

## Message

### Info Message

```
gui-dialog -t msg --title=greeting --msg="你好，世界！"
```

![msg info](https://raw.githubusercontent.com/jiazhoulvke/gui-dialog/master/assets/example_msg_info.png?sanitize=true&raw=true)

### Error Message

```
gui-dialog -t msg --message_dialog_type=error --title="Alert!" --msg="dangerous"
```

![msg error](https://raw.githubusercontent.com/jiazhoulvke/gui-dialog/master/assets/example_msg_error.png?sanitize=true&raw=true)

### Yes Or No

```
gui-dialog -t msg --message_dialog_type=yes_or_no --msg="Are you ok?" -o json
```

![msg yes or no](https://raw.githubusercontent.com/jiazhoulvke/gui-dialog/master/assets/example_msg_yes_or_no.png?sanitize=true&raw=true)

console output:

```json
{ "value": true, "error": "" }
```

```
gui-dialog -t msg --message_dialog_type=yes_or_no --msg="Are you ok?" -o text
```

console output:

> true

## File

```
gui-dialog --start_file="C:\Windows\System32\help.exe" --filter_desc="programs" --filter="exe" -o json
```

![file](https://raw.githubusercontent.com/jiazhoulvke/gui-dialog/master/assets/example_file.png?sanitize=true&raw=true)

console output:

```json
{ "value": "C:\\Windows\\System32\\help.exe", "error": "" }
```

if canceled, console output:

```json
{ "value": "", "error": "Cancelled" }
```

## Directory

```
gui-dialog -t dir --start_dir="C:\Windows\System32\"
```

![dir](https://raw.githubusercontent.com/jiazhoulvke/gui-dialog/master/assets/example_directory.png?sanitize=true&raw=true)
