package cmd

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/atotto/clipboard"
	ci "github.com/skanehira/clipboard-image/v2"
	"github.com/spf13/cobra"
)

var pasteImg = &cobra.Command{
	Use: "pasteimg",
	Run: func(cmd *cobra.Command, args []string) {
		r, err := ci.Read()
		if err != nil {
			log.Fatal(err)
		}

		output := fmt.Sprintf("assets/image/%d.png", time.Now().Unix())
		f, err := os.OpenFile(output, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		if _, err := io.Copy(f, r); err != nil {
			log.Fatal(err)
		}
		// write ref link
		str := fmt.Sprintf("<img width=400 src='/%s'/>\n", output)
		str += fmt.Sprintf("![img](/%s)", output)
		clipboard.WriteAll(str)
	},
}

func init() {
	rootCmd.AddCommand(pasteImg)
}
