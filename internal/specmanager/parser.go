package specmanager

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
)

func ParseSpecDir(root string) ([]Spec, error) {
	specsDir := filepath.Join(root, ".kiro", "specs")

	entries, err := os.ReadDir(specsDir)
	if err != nil {
		if os.IsNotExist(err) {
			return []Spec{}, nil
		}
		return nil, err
	}

	var specs []Spec
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		name := entry.Name()
		if strings.HasPrefix(name, "_") {
			continue
		}

		folderPath := filepath.Join(specsDir, name)
		info, err := os.Stat(folderPath)
		if err != nil {
			continue
		}

		spec := Spec{
			Name:    name,
			Path:    filepath.Join(".kiro", "specs", name),
			ModTime: info.ModTime(),
		}

		files, err := os.ReadDir(folderPath)
		if err != nil {
			continue
		}

		hasRequirements := false
		hasTasks := false
		for _, f := range files {
			switch f.Name() {
			case "requirements.md":
				hasRequirements = true
			case "tasks.md":
				hasTasks = true
			}
		}

		if hasTasks {
			done, total := parseCheckboxes(filepath.Join(folderPath, "tasks.md"))
			spec.TasksDone = done
			spec.TasksTotal = total
		}

		spec.Status = DeriveStatus(spec, len(files) == 0, hasRequirements, hasTasks)
		specs = append(specs, spec)
	}

	return specs, nil
}

func parseCheckboxes(path string) (done, total int) {
	f, err := os.Open(path)
	if err != nil {
		return 0, 0
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "- [x]") || strings.Contains(line, "- [X]") {
			done++
		} else if strings.Contains(line, "- [ ]") {
			total++
		}
	}
	total += done
	return done, total
}
