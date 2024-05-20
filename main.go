package gosubst

import (
	"bufio"
	"context"
	"fmt"
	"os"
)

type CLI struct {
	input   string
	result  string
	output  *bufio.Writer
	context context.Context
}

func (c *CLI) Run(ctx context.Context) error {
	if err := c.prepare(); err != nil {
		return fmt.Errorf("prepare failed: %w", err)
	}

	if err := c.process(); err != nil {
		return fmt.Errorf("process failed: %w", err)
	}

	if err := c.flush(); err != nil {
		return fmt.Errorf("flush failed: %w", err)
	}

	return nil
}

func (c *CLI) prepare() error {
	if err := c.optParse(); err != nil {
		return fmt.Errorf("optParse: %w", err)
	}

	if err := c.scanStdin(); err != nil {
		return fmt.Errorf("scanStdin: %w", err)
	}
	return nil
}

func (c *CLI) process() error {
	// TODO: impl template process
	result := c.input
	c.result = result
	return nil
}

func (c *CLI) flush() error {
	if _, err := c.output.WriteString(c.result); err != nil {
		return fmt.Errorf("WriteString failed: %w", err)
	}
	if err := c.output.Flush(); err != nil {
		return fmt.Errorf("Flush failed: %w", err)
	}

	return nil
}

func (c *CLI) optParse() error {
	// TODO: impl --output
	c.output = bufio.NewWriter(os.Stdout)
	return nil
}

func (c *CLI) scanStdin() error {
	scanner := bufio.NewScanner(os.Stdin)
	if ok := scanner.Scan(); !ok {
		return fmt.Errorf("scanner.Scan failed: %w", scanner.Err())
	}
	c.input = scanner.Text()
	return nil
}
