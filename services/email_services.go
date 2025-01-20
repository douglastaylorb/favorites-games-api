package services

import (
	"fmt"
	"net/smtp"
	"os"
)

func SendPasswordResetEmail(to, resetLink string) error {
	// Configurações do Mailtrap
	smtpHost := os.Getenv("MAILTRAP_HOST")
	smtpPort := os.Getenv("MAILTRAP_PORT")
	smtpUsername := os.Getenv("MAILTRAP_USERNAME")
	smtpPassword := os.Getenv("MAILTRAP_PASSWORD")

	from := "noreply@gamelist.com" // E-mail remetente

	// Autenticação
	auth := smtp.PlainAuth("", smtpUsername, smtpPassword, smtpHost)

	// Mensagem
	subject := "Redefinição de Senha - GameList"
	body := fmt.Sprintf(`
    <html>
        <body>
            <h2>Redefinição de Senha - GameList</h2>
            <p>Você solicitou a redefinição de sua senha. Clique no link abaixo para criar uma nova senha:</p>
            <p><a href="%s">Redefinir Minha Senha</a></p>
            <p>Se você não solicitou esta redefinição, por favor ignore este e-mail.</p>
            <p>Atenciosamente,<br>Equipe GameList</p>
        </body>
    </html>
    `, resetLink)

	msg := []byte(fmt.Sprintf("To: %s\r\n"+
		"From: %s\r\n"+
		"Subject: %s\r\n"+
		"MIME-Version: 1.0\r\n"+
		"Content-Type: text/html; charset=UTF-8\r\n"+
		"\r\n"+
		"%s\r\n", to, from, subject, body))

	// Enviando e-mail
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, msg)
	if err != nil {
		return fmt.Errorf("erro ao enviar e-mail: %v", err)
	}

	return nil
}
