package util

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"archive/zip"
)

// ExtractTemplateArchive unpacks an uploaded email-template archive into dest.
func ExtractTemplateArchive(archive string, dest string) error {
	zr, err := zip.OpenReader(archive)
	if err != nil {
		return err
	}
	defer zr.Close()

	root, err := filepath.Abs(dest)
	if err != nil {
		return err
	}
	prefix := root + string(os.PathSeparator)

	for _, f := range zr.File {
		target := filepath.Join(root, f.Name)
		// An entry name may contain ".." or an absolute path; refuse anything
		// that would land outside dest.
		if target != root && !strings.HasPrefix(target, prefix) {
			return fmt.Errorf("archive entry %q escapes the destination directory", f.Name)
		}
		if f.FileInfo().IsDir() {
			if err := os.MkdirAll(target, 0755); err != nil {
				return err
			}
			continue
		}
		if err := os.MkdirAll(filepath.Dir(target), 0755); err != nil {
			return err
		}
		rc, err := f.Open()
		if err != nil {
			return err
		}
		out, err := os.Create(target)
		if err != nil {
			rc.Close()
			return err
		}
		_, err = io.Copy(out, rc)
		out.Close()
		rc.Close()
		if err != nil {
			return err
		}
	}
	return nil
}
