package mail

import (
	"testing"

	"github.com/alim7007/go_bank_k8s/util"
	"github.com/stretchr/testify/require"
)

func TestSendEmailWithGmail(t *testing.T) {
	//github workflow will not test it, because of flag -short in "go test -v -cover -short ./..."
	if testing.Short() {
		t.Skip()
	}

	config, err := util.LoadConfig("..")
	require.NoError(t, err)

	sender := NewGmailSender(config.EmailSenderName, config.EmailSenderAddress, config.EmailSenderPassword)

	subject := "A test email"
	content := `
	<h1>Hello world</h1>
	<p>This is a test message from <a href="http://techschool.guru">Olim Bank</a></p>
	`
	to := []string{"toni7007olim@gmail.com"}
	attachFiles := []string{"./text.txt"}

	err = sender.SendEmail(subject, content, to, nil, nil, attachFiles)
	require.NoError(t, err)
}
