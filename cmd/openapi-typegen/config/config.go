package config

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/signal"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

func NewDefaultConfig() Config {

	ctx, cf := signal.NotifyContext(context.Background(), os.Interrupt)

	return Config{
		OutFilePath: "stdout",
		outFile:     os.Stdout,

		ctx: ctx,
		cf:  cf,
	}
}

type Config struct {
	OpenAPIFilePath string `koanf:"file" short:"f" description:"only compare directories"`
	OutFilePath     string `koanf:"out" short:"o" description:"out file path or 'stdout'"`

	openApiSpec *openapi3.T    `koanf:"-"`
	outFile     io.WriteCloser `koanf:"-"`
	ctx         context.Context
	cf          context.CancelFunc
}

func (c *Config) Close() error {
	c.cf()
	return c.outFile.Close()
}

func (c *Config) Context() context.Context {
	return c.ctx
}

func (c *Config) Out() io.Writer {
	return c.outFile
}

func (c *Config) Document() *openapi3.T {
	return c.openApiSpec
}

func (c *Config) Validate() error {
	if !strings.EqualFold(c.OutFilePath, "stdout") {
		f, err := os.Create(c.OutFilePath)
		if err != nil {
			return fmt.Errorf("failed to open out file %q: %w", c.OutFilePath, err)
		}
		c.outFile = f
	} else {
		c.outFile = os.Stdout
	}

	loader := openapi3.NewLoader()
	loader.Context = c.ctx

	doc, err := loader.LoadFromFile(c.OpenAPIFilePath)
	if err != nil {
		return fmt.Errorf("failed to load openapi specification: %w", err)
	}

	err = doc.Validate(loader.Context)
	if err != nil {
		return fmt.Errorf("failed to validate openapi specification: %w", err)
	}

	c.openApiSpec = doc

	return nil
}
