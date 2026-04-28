// Copyright (c) The OpenTofu Authors
// SPDX-License-Identifier: MPL-2.0


package handlers

import (
	"fmt"
	"context"
	lsp "github.com/opentofu/tofu-ls/internal/protocol"


)

func (s *service) Rename(ctx context.Context, params *lsp.RenameParams) (*lsp.WorkspaceEdit, error) {
    return nil, fmt.Errorf("rename not implemented")
}