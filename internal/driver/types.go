// -*- Mode: Go; indent-tabs-mode: t -*-
//
// Copyright (C) 2023 Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

package driver

type RTSPAuthRequest struct {
	Ip       string `json:"ip"`
	User     string `json:"user"`
	Password string `json:"password"`
	Path     string `json:"path"`
	Protocol string `json:"protocol"`
	Id       string `json:"id"`
	Action   string `json:"action"`
	Query    string `json:"query"`
}
