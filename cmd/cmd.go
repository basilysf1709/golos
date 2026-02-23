package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"strconv"
	"syscall"

	"github.com/gordonklaus/portaudio"

	"github.com/iqbalyusuf/golos/internal"
	"github.com/iqbalyusuf/golos/processor"
)

func pidFile() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".config", "golos", "golos.pid")
}

func Run(detach bool) {
	if detach {
		runDetached()
		return
	}
	runForeground()
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
