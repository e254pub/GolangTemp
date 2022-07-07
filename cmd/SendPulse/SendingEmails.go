package SendPulse

import (
	"os"
	"strconv"
)

// SendingReg рассылка сообщений на регистрацию/**
func SendingReg(receiverEmail string, token string) error {
	initialize(os.Getenv("SEND_PULSE_ID"), os.Getenv("SEND_PULSE_SECRET"), os.Getenv("SEND_PULSE_NAME"), os.Getenv("SEND_PULSE_EMAIL"))

	templateId, _ := strconv.Atoi(os.Getenv("SEND_PULSE_TEMPLATE_ID_REG"))
	subject := os.Getenv("SEND_PULSE_SUBJECT_REG")
	recipientsData := []Recipient{
		{receiverEmail},
	}

	template, subject, recipients := composeMessage(templateId, recipientsData, subject, token)

	err := sendEmail(template, subject, recipients)

	if err != nil {
		return err
	}

	return nil
}

// composeMessage Формирует сообщение и список получателей /**
func composeMessage(templateId int, recipientsData []Recipient, subjectName string, token string) (Template, string, []Recipient) {
	var recipients []Recipient

	for i := 0; i < len(recipientsData); i++ {
		recipientsSlice := Recipient{Email: recipientsData[i].Email}
		recipients = append(recipients, recipientsSlice)
	}
	template := Template{Id: templateId, Variables: TemplateVariable{Token: token}}
	subject := subjectName

	return template, subject, recipients
}
