package mpwx

import (
	"fmt"
	"log"
)

// non-blocking is appropriate for most case
func (c *MpwxClient) PushToAdminf(format string, a ...any) {
	text := fmt.Sprintf(format, a...)
	c.PushToAdmin(text)
}
func (c *MpwxClient) PushToAdmin(text string) {
	go func() {
		if err := c.pushToAdmin(text); err != nil {
			log.Println("failed to pushToAdmin:", err)
		}
	}()
}

// blocking
func (c *MpwxClient) pushToAdmin(text string) error {
	return c.SendTemplateMessage(c.adminUser, c.plainTemplate,
		// 坑！text不加引号就变成了动态取其value作为键值
		// map[string]string{text: text})
		map[string]string{"text": text})
}
