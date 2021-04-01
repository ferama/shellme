package shwscontroller

import (
	"log"
	"net/http"
	"os/exec"
	"shellme/utils"
	"strconv"
	"strings"

	"github.com/creack/pty"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

const (
	logContext = "[SHWSCONTROLLER]"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// WsHandler ...
func WsHandler(c *gin.Context) {
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("%v", err)
		return
	}

	flags := utils.GetFlags()

	cmd := exec.Command("sh", "-c", *flags.PtyCmd)

	// Start the command with a pty.
	ptmx, err := pty.Start(cmd)
	if err != nil {
		return
	}

	defer func() {
		_ = ptmx.Close()
		cmd.Process.Kill()
		cmd.Process.Wait()
		ws.Close()
	}()

	go func() {
		for {
			buf := make([]byte, 2048)
			l, _ := ptmx.Read(buf)

			// nothing to read anymore
			if l == 0 {
				return
			}
			if err = ws.WriteMessage(websocket.BinaryMessage, buf[0:l]); err != nil {
				// fmt.Println(err)
				return
			}
		}
	}()

	for {
		_, msg, err := ws.ReadMessage()
		if err != nil {
			return
		}

		// ignore Ctrl+D - EOF. Prevents shell from exit
		// To actually kill the shell, refresh the browser window
		// TODO: the following solution prevent to exit byobu sessions too!
		// if msg[0] == 4 {
		// 	continue
		// }

		// Handles terminal resize requests
		// They are special messages in the form
		// 		|RESIZE|85:24
		// If the message is detected, Setsize is called
		// and the message is not written to the shell input
		msgString := string(msg)
		if strings.HasPrefix(msgString, "|RESIZE|") {
			// logger.Debug(logContext, "resize request")
			values := strings.Replace(msgString, "|RESIZE|", "", 1)
			valuesArray := strings.Split(values, ":")
			cols, _ := strconv.Atoi(valuesArray[0])
			rows, _ := strconv.Atoi(valuesArray[1])
			size := pty.Winsize{
				Rows: uint16(rows),
				Cols: uint16(cols),
			}
			pty.Setsize(ptmx, &size)
		} else {
			ptmx.Write(msg)
		}
	}
}
