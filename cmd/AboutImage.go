package cmd

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strconv"

	"github.com/cheggaaa/pb/v3"
)

func GenDomDisk(image string, Num int) (output string) {
	output = Datadir + "/volumes/virt-go-" + image + "-" + strconv.Itoa(Num) + "root"

	in, err := os.Open(Datadir + "/images/" + image)
	if err != nil {
		fmt.Println("Open base image", err)
		os.Exit(70)
	}
	defer in.Close()

	out, err := os.Create(output)
	if err != nil {
		fmt.Println("Open output file", err)
		os.Exit(71)
	}
	defer out.Close()

	if _, err = io.Copy(out, in); err != nil {
		fmt.Println("Copy err", err)
		os.Exit(71)
	}

	err = out.Sync()
	if err != nil {
		fmt.Println("Syn err", err)
	}

	cmd := exec.Command("qemu-img", "resize", output, "20G")

	if err = cmd.Run(); err != nil {
		fmt.Println(err)
		os.Exit(90)
	}

	return output
}

func GenImage(base string, image string) {
	in, err := os.Open(base)
	if err != nil {
		fmt.Println("Open base image", err)
		os.Exit(70)
	}
	defer in.Close()

	out, err := os.Create(Datadir + "/images/" + image)
	if err != nil {
		fmt.Println("Open output file", err)
		os.Exit(71)
	}
	defer out.Close()

	info, _ := in.Stat()
	len := info.Size()
	bar := pb.Full.Start64(len)

	barReader := bar.NewProxyReader(in)

	if _, err = io.Copy(out, barReader); err != nil {
		fmt.Println("Copy err", err)
		os.Exit(71)
	}

	bar.Finish()

	err = out.Sync()
	if err != nil {
		fmt.Println("Syn err", err)
	}
}
