// Connections support one concurrent reader and one concurrent writer.
// https://github.com/gorilla/websocket/tree/master/examples
package wsg

import "github.com/gorilla/websocket"

// WriteClose 写关闭消息
func WriteClose(conn *websocket.Conn) error {
	data := websocket.FormatCloseMessage(websocket.CloseNormalClosure, "")
	err := conn.WriteMessage(websocket.CloseMessage, data)
	return err
}
