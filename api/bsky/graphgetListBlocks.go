// Code generated by cmd/lexgen (see Makefile's lexgen); DO NOT EDIT.

package bsky

// schema: app.bsky.graph.getListBlocks

import (
	"context"

	"github.com/bluesky-social/indigo/xrpc"
)

// GraphGetListBlocks_Output is the output of a app.bsky.graph.getListBlocks call.
type GraphGetListBlocks_Output struct {
	Cursor *string               `json:"cursor,omitempty" cborgen:"cursor,omitempty"`
	Lists  []*GraphDefs_ListView `json:"lists" cborgen:"lists"`
}

// GraphGetListBlocks calls the XRPC method "app.bsky.graph.getListBlocks".
func GraphGetListBlocks(ctx context.Context, c *xrpc.Client, cursor string, limit int64) (*GraphGetListBlocks_Output, error) {
	var out GraphGetListBlocks_Output

	params := map[string]interface{}{
		"cursor": cursor,
		"limit":  limit,
	}
	if err := c.Do(ctx, xrpc.Query, "", "app.bsky.graph.getListBlocks", params, nil, &out); err != nil {
		return nil, err
	}

	return &out, nil
}