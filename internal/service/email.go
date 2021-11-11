/*
	Copyright © 2021 Durudex

	This file is part of Durudex: you can redistribute it and/or modify
	it under the terms of the GNU Affero General Public License as
	published by the Free Software Foundation, either version 3 of the
	License, or (at your option) any later version.

	Durudex is distributed in the hope that it will be useful,
	but WITHOUT ANY WARRANTY; without even the implied warranty of
	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
	GNU Affero General Public License for more details.

	You should have received a copy of the GNU Affero General Public License
	along with Durudex. If not, see <https://www.gnu.org/licenses/>.
*/

package service

import (
	"github.com/Durudex/durudex-notif-service/internal/config"
	"github.com/Durudex/durudex-notif-service/pkg/email"
)

type EmailService struct {
	email email.Email
	cfg   config.EmailConfig
}

// Creating a new email service.
func NewEmailService(email email.Email, emailConfig config.EmailConfig) *EmailService {
	return &EmailService{
		email: email,
		cfg:   emailConfig,
	}
}

type verificationEmailInput struct {
	Name string
	Code int32
}

// Send to user email code.
func (s *EmailService) UserCode(to, name string, code int32) (bool, error) {
	msg := email.SendEmailInput{
		To:      to,
		Subject: "Verification Code",
	}

	// Generate email html template.
	templateInput := verificationEmailInput{Name: name, Code: code}
	err := msg.GenerateBodyFromHTML(s.cfg.Template.Verification, templateInput)
	if err != nil {
		return false, err
	}

	// Send email message.
	status, err := s.email.Send(msg)
	if err != nil {
		return status, err
	}

	return status, nil
}
