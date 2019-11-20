package bugs

import (
	"fmt"
	"os"
)

//var dops = Directory(os.PathSeparator)
//var sops = string(os.PathSeparator)

func (b *Bug) Read(p []byte) (int, error) {
	if b.DescriptionFileName == "" {
		return 0, ErrNoDescription
	}
	if b.descFile == nil {
		dir := b.Direr()
		fp, err := os.OpenFile(string(dir)+sops+b.DescriptionFileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
		b.descFile = fp
		if err != nil {
			fmt.Fprintf(os.Stderr, "err: %s", err.Error())
			return 0, ErrNoDescription
		}
	}

	return b.descFile.Read(p)
}

func (b *Bug) Write(data []byte) (n int, err error) {
	if b.DescriptionFileName == "" {
		return 0, ErrNoDescription
	}
	if b.descFile == nil {
		dir := b.Direr()
		os.MkdirAll(string(dir), 0755)
		fp, err := os.OpenFile(string(dir)+sops+b.DescriptionFileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error writing to issue: %s", err.Error())
			return 0, err
		}
		b.descFile = fp
	}
	return b.descFile.Write(data)
}

// WriteAt makes a directory, writes a byte string to the Description using an offset.
// It returns the number of bytes written and an error.
func (b *Bug) WriteAt(data []byte, off int64) (n int, err error) {
	if b.DescriptionFileName == "" {
		return 0, ErrNoDescription
	}
	if b.descFile == nil {
		dir := b.Direr()
		os.MkdirAll(string(dir), 0755)
		fp, err := os.OpenFile(b.DescriptionFileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error writing to issue: %s", err.Error())
			return 0, err
		}
		b.descFile = fp
	}
	return b.descFile.WriteAt(data, off)
}

// Close returns an error if there is an error closing the descFile of an issue.
func (b Bug) Close() error {
	if b.descFile != nil {
		err := b.descFile.Close()
		b.descFile = nil
		return err
	}
	return nil
}

// Remove deletes the directory and files of an issue.
func (b *Bug) Remove() error {
	dir := b.Direr()
	if dir != "" {
		return os.RemoveAll(string(dir))
	}
	return ErrNotFound
}
