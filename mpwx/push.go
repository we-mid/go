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
		if err := c.PushToAdminSync(text); err != nil {
			log.Println("failed to PushToAdminSync:", err)
		}
	}()
}

// blocking
func (c *MpwxClient) PushToAdminSync(text string) error {
	return c.SendTemplateMessage(c.adminUser, c.plainTemplate, map[string]string{text: text})
}
