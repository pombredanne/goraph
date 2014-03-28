package dfs

import (
	"log"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"

	"github.com/gyuho/goraph/graph/gsd"
)

// DeleteNonAlnum removes all alphanumeric characters.
func DeleteNonAlnum(str string) string {
	// alphanumeric (== [0-9A-Za-z])
	// \s is a white space character
	validID := regexp.MustCompile(`[^[:alnum:]\s]`)
	return validID.ReplaceAllString(str, "")
}

// ShowDFS shows the traversed DFS.
func ShowDFS(g *gsd.Graph, filename string) string {
	result := "graph " + DeleteNonAlnum(filename) + " {" + "\n"
	for _, edge := range *g.GetEdges() {
		wt := strconv.FormatFloat(edge.(*gsd.Edge).Weight, 'f', -1, 64)
		result += "\t" + edge.(*gsd.Edge).Src.ID + " -- " + edge.(*gsd.Edge).Dst.ID + " [label=" + wt + "]" + "\n"
	}
	tb := DFS(g)
	tb = strings.Replace(tb, "→", "--", -1)
	result += "\t" + tb + " [label=DFS, color=blue]" + "\n"
	result += "}"

	file, err := os.Create(filename)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	file.WriteString(result)
	cmd := exec.Command("open", filename)
	er := cmd.Run()
	if er != nil {
		log.Fatal(er)
	}

	return result
}
