package cli

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"

	"github.com/gordonklaus/portaudio"

	"github.com/basilysf1709/golos/internal"
	"github.com/basilysf1709/golos/processor"
)

func pidFile() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".config", "golos", "golos.pid")
}

func Run(detach bool) {
	if err := checkAlreadyRunning(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if detach {
		runDetached()
		return
	}
	runForeground()
}

func checkAlreadyRunning() error {
	data, err := os.ReadFile(pidFile())
	if err != nil {
		return nil // no pid file — not running
	}

	pid, err := strconv.Atoi(string(data))
	if err != nil {
		os.Remove(pidFile()) // corrupt pid file — clean up
		return nil
	}

	// The detached parent writes the child's PID before the child starts,
	// so the child would find its own PID and block itself.
	if pid == os.Getpid() {
		return nil
	}

	proc, err := os.FindProcess(pid)
	if err != nil {
		os.Remove(pidFile())
		return nil
	}

	// Signal 0 checks if the process exists without killing it
	if err := proc.Signal(syscall.Signal(0)); err != nil {
		os.Remove(pidFile()) // stale pid file — process is gone
		return nil
	}

	return fmt.Errorf("golos is already running (PID %d)\n  Stop it first with: golos stop", pid)
}

func DictImport(args []string) {
	if len(args) < 1 {
		fmt.Fprintln(os.Stderr, "usage: golos import <file.toml>")
		os.Exit(1)
	}

	d := internal.LoadDictionary()
	count, err := d.Import(args[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("imported %d entries\n", count)
}

func DictAdd(args []string) {
	if len(args) < 2 {
		fmt.Fprintln(os.Stderr, "usage: golos add <phrase> <replacement>")
		os.Exit(1)
	}
	phrase := args[0]
	replacement := strings.Join(args[1:], " ")

	d := internal.LoadDictionary()
	if err := d.Add(phrase, replacement); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("added: %q → %q\n", phrase, replacement)
}

func DictDelete(args []string) {
	if len(args) < 1 {
		fmt.Fprintln(os.Stderr, "usage: golos delete <phrase>")
		os.Exit(1)
	}
	phrase := strings.Join(args, " ")

	d := internal.LoadDictionary()
	if !d.Delete(phrase) {
		fmt.Fprintf(os.Stderr, "not found: %q\n", phrase)
		os.Exit(1)
	}
	fmt.Printf("deleted: %q\n", phrase)
}

func DictList() {
	d := internal.LoadDictionary()
	entries := d.List()
	if len(entries) == 0 {
		fmt.Println("dictionary is empty")
		return
	}
	for phrase, replacement := range entries {
		fmt.Printf("  %q → %q\n", phrase, replacement)
	}
}

func Stop() {
	data, err := os.ReadFile(pidFile())
	if err != nil {
		fmt.Fprintln(os.Stderr, "golos is not running")
		os.Exit(1)
	}

	pid, err := strconv.Atoi(string(data))
	if err != nil {
		fmt.Fprintln(os.Stderr, "invalid pid file")
		os.Remove(pidFile())
		os.Exit(1)
	}

	proc, err := os.FindProcess(pid)
	if err != nil {
		fmt.Fprintln(os.Stderr, "golos is not running")
		os.Remove(pidFile())
		os.Exit(1)
	}

	if err := proc.Signal(syscall.SIGTERM); err != nil {
		fmt.Fprintf(os.Stderr, "failed to stop golos: %v\n", err)
		os.Remove(pidFile())
		os.Exit(1)
	}

	os.Remove(pidFile())
	fmt.Printf("golos stopped (PID %d)\n", pid)
}

func runDetached() {
	exe, err := os.Executable()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	// Re-launch without -d, with --background flag
	args := []string{"--background"}
	for _, a := range os.Args[1:] {
		if a != "-d" && a != "--detach" {
			args = append(args, a)
		}
	}

	cmd := exec.Command(exe, args...)
	cmd.SysProcAttr = &syscall.SysProcAttr{Setsid: true}

	// Log to file
	home, _ := os.UserHomeDir()
	logDir := filepath.Join(home, ".config", "golos")
	os.MkdirAll(logDir, 0755)
	logPath := filepath.Join(logDir, "golos.log")

	logFile, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating log file: %v\n", err)
		os.Exit(1)
	}
	cmd.Stdout = logFile
	cmd.Stderr = logFile

	if err := cmd.Start(); err != nil {
		fmt.Fprintf(os.Stderr, "Error starting background process: %v\n", err)
		os.Exit(1)
	}

	// Write PID file
	os.WriteFile(pidFile(), []byte(strconv.Itoa(cmd.Process.Pid)), 0644)

	fmt.Printf("golos running in background (PID %d)\n", cmd.Process.Pid)
	fmt.Printf("  Log: %s\n", logPath)
	fmt.Println("  Stop with: golos stop")
}

func runForeground() {
	app, err := processor.Setup()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	defer portaudio.Terminate()

	// Write PID file so `golos stop` works in both modes
	os.MkdirAll(filepath.Dir(pidFile()), 0755)
	os.WriteFile(pidFile(), []byte(strconv.Itoa(os.Getpid())), 0644)
	defer os.Remove(pidFile())

	// Handle Ctrl+C
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigCh
		fmt.Println("\nShutting down...")
		os.Remove(pidFile())
		internal.StopHotkey()
	}()

	fmt.Println("golos — speech-to-text for Claude Code")
	fmt.Printf("  Output:  %s\n", app.Config.OutputMode)
	fmt.Printf("  Hotkey:  %s\n", app.Config.Hotkey)
	fmt.Printf("  Model:   Deepgram Nova-3\n")
	fmt.Println()
	fmt.Println("Ready — hold hotkey to speak")

	if err := internal.ListenHotkey(app.Hotkey, app.Proc.Start, app.Proc.Stop); err != nil {
		fmt.Fprintf(os.Stderr, "Hotkey error: %v\n", err)
		os.Exit(1)
	}
}
