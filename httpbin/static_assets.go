package httpbin

import (
	"embed"
)

//go:embed static/*
var staticAssets embed.FS

// staticAsset loads an embedded static asset by name.
func staticAsset(name string) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// mustStaticAsset loads an embedded static asset by name, panicking on error.
func mustStaticAsset(name string) []byte { _ = "STUB: not implemented"; return nil }

func mustRenderTemplate(name string, data any) []byte { _ = "STUB: not implemented"; return nil }
