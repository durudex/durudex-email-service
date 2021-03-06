/*
 * Copyright © 2022 Durudex
 *
 * This file is part of Durudex: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as
 * published by the Free Software Foundation, either version 3 of the
 * License, or (at your option) any later version.
 *
 * Durudex is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with Durudex. If not, see <https://www.gnu.org/licenses/>.
 */

package v1

import (
	"context"

	"github.com/durudex/durudex-email-service/internal/service"
	v1 "github.com/durudex/durudex-email-service/pkg/pb/durudex/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Email user gRPC server handler.
type UserHandler struct {
	service service.User
	v1.UnimplementedEmailUserServiceServer
}

// Creating a new email gRPC handler.
func NewEmailHandler(service service.User) *UserHandler {
	return &UserHandler{service: service}
}

// Send to user email verification code.
func (h *UserHandler) SendEmailUserCode(ctx context.Context, input *v1.SendEmailUserCodeRequest) (*v1.SendEmailUserCodeResponse, error) {
	err := h.service.Code(input.Email, input.Username, input.Code)
	if err != nil {
		return &v1.SendEmailUserCodeResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &v1.SendEmailUserCodeResponse{}, nil
}

// Send to user email logged in information.
func (h *UserHandler) SendEmailUserLoggedIn(ctx context.Context, input *v1.SendEmailUserLoggedInRequest) (*v1.SendEmailUserLoggedInResponse, error) {
	err := h.service.LoggedIn(input.Email, input.Ip)
	if err != nil {
		return &v1.SendEmailUserLoggedInResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &v1.SendEmailUserLoggedInResponse{}, nil
}

// Send to user email register information.
func (h *UserHandler) SendEmailUserRegister(ctx context.Context, input *v1.SendEmailUserRegisterRequest) (*v1.SendEmailUserRegisterResponse, error) {
	err := h.service.Register(input.Email, input.Username)
	if err != nil {
		return &v1.SendEmailUserRegisterResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &v1.SendEmailUserRegisterResponse{}, nil
}
