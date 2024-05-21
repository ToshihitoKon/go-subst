package gosubst

import (
	"bufio"
	"fmt"
	"os"

	goconfig "github.com/kayac/go-config"
)

type CLI struct {
	input  []byte
	result []byte
	output *bufio.Writer
}

func (c *CLI) Run() error {
	if err := c.prepare(); err != nil {
		return err
	}

	if err := c.process(); err != nil {
		return err
	}

	if err := c.flush(); err != nil {
		return err
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
	result, err := goconfig.ReadWithEnvBytes(c.input)
	if err != nil {
		return fmt.Errorf("goconfig.ReadWithEnvBytes failed: %w", err)
	}

	c.result = result
	return nil
}

func (c *CLI) flush() error {
	if _, err := c.output.Write(c.result); err != nil {
		return fmt.Errorf("Write failed: %w", err)
	}
	if err := c.output.Flush(); err != nil {
		return fmt.Errorf("Flush failed: %w", err)
	}
	return nil
}

func (c *CLI) optParse() error {
	// TODO: impl options input,output
	c.output = bufio.NewWriter(os.Stdout)
	return nil
}

func (c *CLI) scanStdin() error {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(func(data []byte, atEOF bool) (int, []byte, error) {
		if atEOF && len(data) == 0 {
			return 0, nil, nil
		}
		return len(data), data, nil
	})

	for scanner.Scan() {
		c.input = append(c.input, scanner.Bytes()...)
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("scanner.Scan failed: %w", err)
	}

	return nil
}
