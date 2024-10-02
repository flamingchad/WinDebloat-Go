# WinDebloat-Go

## 🛠️ A Simple Tool to Debloat Windows 11

WinDebloat-Go is a lightweight CLI tool built with Go to help users remove unwanted telemetry, tracking, and unnecessary features from Windows 11. This tool helps you disable Windows services, modify registry entries, and block telemetry by editing the hosts file.

---

## 🚀 Features

- **Disable DiagTrack and other telemetry services**: Stops and disables unnecessary Windows telemetry services.
- **Block Microsoft telemetry via the hosts file**: Appends blocklists to prevent data from being sent to Microsoft.
- **Disable Cortana, Bing integration**: Removes Bing search suggestions from the Start Menu.
- **Disable Weather, News, and Interests**: Remove distracting news and weather popups from the taskbar.
- **Disable Windows Copilot**: Remove the new Windows Copilot feature from Windows 11.
- **Apply all settings at once**: Use a single command to apply all debloating features in one go.

---

## 📦 Installation

1. **Install Go**: First, make sure you have Go installed. You can download it from [here](https://golang.org/dl/).
2. **Clone the repository**:

    ```bash
    git clone https://github.com/flamingchad/WinDebloat-Go.git
    cd WinDebloat-Go
    ```

3. **Install Cobra CLI**:

    ```bash
    go get -u github.com/spf13/cobra/cobra
    ```

4. **Build the project**:

    ```bash
    go build -o windebloat
    ```

5. **Run the tool**:

    ```bash
    ./windebloat
    ```

---

## 📖 Usage

WinDebloat-Go is a command-line tool. After building the project, you can use the following commands to debloat your Windows installation:

```bash
windebloat <command>
Available Commands
disable-diagtrack	Disable the DiagTrack telemetry service
disable-telemetry	Disable Windows telemetry in the registry
disable-weather-news    Disable weather, news, and interests on the taskbar
remove-bing	        Disable Bing search suggestions in the Start Menu
disable-copilot	        Disable the Windows Copilot feature
disable-microsoft-telemetry	Block Microsoft telemetry servers by adding entries to the hosts file
all	                Apply all settings at once for a complete debloating experience

```
Example:

```bash
windebloat all
```

🛑 Important Notes

    Admin privileges: Most of the operations require elevated privileges. 
    Make sure to run the executable with administrator rights.
    Backup: Before making any changes to your system, it is always recommended to create a system backup to avoid any unwanted issues.

⚙️ How It Works

    Services: The tool uses PowerShell commands to stop and disable services like DiagTrack.
    Registry: Edits Windows Registry values to disable telemetry-related features.
    Hosts File: Appends a blocklist to the hosts file to block telemetry and tracking servers.
    Cobra CLI: Uses Cobra CLI for creating easy-to-use subcommands for each debloating feature.

💻 Technologies Used

    Golang: The core language for building the application.
    Cobra: A CLI library to manage commands and flags.
    PowerShell: Executes system commands to stop services and make system-level changes.
    Windows Registry: For modifying system configurations.

🛡️ Disclaimer

WinDebloat-Go is intended for users who understand the risks of modifying their system. The authors are not responsible for any unintended behavior that may result from using this tool. Always use caution and create backups.