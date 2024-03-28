package main

import (
	"bytes"
	"os/exec"
	"time"

	"github.com/gofiber/fiber/v2"
)

// Request structure to receive code execution requests
type ExecuteRequest struct {
	Code string `json:"code"`
}

// Response structure to send back execution results
type ExecuteResponse struct {
	Status string `json:"status"`
	Result string `json:"result"`
}

func executeCode(c *fiber.Ctx) error {
	req := new(ExecuteRequest)

	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse request"})
	}

	cmd := exec.Command("python", "-c", req.Code)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out

	// Set a 5-second timeout for code execution
	err := runCmdWithTimeout(cmd, 5*time.Second)

	resp := new(ExecuteResponse)
	if err != nil {
		resp.Status = "error"
		resp.Result = err.Error() + ": " + out.String()
	} else {
		resp.Status = "success"
		resp.Result = out.String()
	}

	return c.JSON(resp)
}

// runCmdWithTimeout runs the given command with a specified timeout duration.
func runCmdWithTimeout(cmd *exec.Cmd, timeout time.Duration) error {
	if err := cmd.Start(); err != nil {
		return err
	}

	done := make(chan error)
	go func() { done <- cmd.Wait() }()

	select {
	case <-time.After(timeout):
		cmd.Process.Kill()
		return exec.ErrNotFound // You can define a custom error for timeout
	case err := <-done:
		return err
	}
}

func main() {
	app := fiber.New()

	app.Post("/execute", executeCode)

	app.Listen(":5000")
}
