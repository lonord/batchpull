package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"strings"
	"syscall"

	"github.com/gookit/color"
	"github.com/subchen/go-cli/v3"
)

func runApp(c *cli.Context) {
	checkGit()
	dir := parseDir(c)
	walkGitDir(dir, depFlagVal(c), handleRepo)
}

func handleRepo(p string) {
	color.Yellow.Print("[PULL] ")
	fmt.Print(p, "...")
	if isGitRepoDirty(p) {
		gitStash(p)
		defer gitStashPop(p)
	}
	var errBuf bytes.Buffer
	ret := gitPull(p, &errBuf)
	if ret == 0 {
		color.LightGreen.Println(" OK")
	} else {
		color.LightRed.Printf(" ERROR(%d)\n", ret)
		errStr := errBuf.String()
		if !strings.HasSuffix(errStr, "\n") {
			errStr += "\n"
		}
		color.Red.Print(errStr)
	}
}

func walkGitDir(pathAll string, depLeft int, fn func(p string)) {
	info, err := os.Stat(pathAll)
	if err != nil {
		errorExit(err.Error())
	}
	if !info.IsDir() {
		return
	}
	if isGitDir(pathAll) {
		fn(pathAll)
	}
	if depLeft > 0 {
		files, err := ioutil.ReadDir(pathAll)
		if err != nil {
			error(fmt.Sprintln("could not read files in directory", pathAll))
			return
		}
		for _, f := range files {
			walkGitDir(path.Join(pathAll, f.Name()), depLeft-1, fn)
		}
	}
}

func isGitDir(p string) bool {
	info, err := os.Stat(path.Join(p, ".git"))
	return err == nil && info.IsDir()
}

func isGitRepoDirty(p string) bool {
	cmd := exec.Command("git", "-C", p, "status", "--short", "--porcelain")
	b, err := cmd.CombinedOutput()
	if err != nil {
		errorExit(err.Error())
	}
	lines := strings.Split(string(b), "\n")
	for _, l := range lines {
		if len(strings.TrimSpace(l)) > 0 && !strings.HasPrefix(l, "??") {
			return true
		}
	}
	return false
}

func gitStash(p string) {
	cmd := exec.Command("git", "-C", p, "stash")
	cmd.Run()
}

func gitStashPop(p string) {
	cmd := exec.Command("git", "-C", p, "stash", "pop")
	cmd.Run()
}

func gitPull(p string, wErr io.Writer) int {
	cmd := exec.Command("git", "-C", p, "pull")
	if wErr != nil {
		cmd.Stderr = wErr
	}
	if err := cmd.Run(); err != nil {
		if exiterr, ok := err.(*exec.ExitError); ok {
			if status, ok := exiterr.Sys().(syscall.WaitStatus); ok {
				return status.ExitStatus()
			}
		}
		return 1
	}
	return 0
}

func checkGit() {
	_, err := exec.LookPath("git")
	if err != nil {
		if err == exec.ErrNotFound {
			errorExit("could not find git in your PATH")
		}
		errorExit(err.Error())
	}
}

func parseDir(c *cli.Context) string {
	if c.NArg() == 1 {
		return c.Arg(0)
	}
	d, err := os.Getwd()
	if err != nil {
		errorExit(err.Error())
	}
	return d
}
