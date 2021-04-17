package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/docker"
	"github.com/stretchr/testify/assert"
)

// Standard Go test, with the "Test" prefix and accepting the *testing.T struct.
func TestDockerImage(t *testing.T) {
	// Define the docker tag
	tag := "terratest-examples:docker"

	// The build options we would pass to `docker build`
	buildOptions := &docker.BuildOptions{
		Tags: []string{tag},
		OtherOptions: []string{
			"--pull",
			"--no-cache",
			"-f",
			"../docker/Dockerfile",
		},
	}

	// The wrapped docker build command, with the `../docker` folder as the
	// build context
	docker.Build(t, "../docker", buildOptions)

	// A testing table to test different aspects of the image.
	tt := []struct {
		name       string
		entrypoint string
		command    string
		expected   string
	}{
		{name: "test that node is installed", entrypoint: "node", command: "--version", expected: "14"},
		{name: "test that the testing.txt is present", entrypoint: "ls", command: "/tmp/testing.txt", expected: "testing.txt"},
	}

	// Iterate over the testing table to create test cases
	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			// Allow the tests to run in parallel
			t.Parallel()

			// The docker run options
			opts := &docker.RunOptions{
				// Remove the container once finished
				Remove: true,
				// Entrypoint is variable from the test table
				Entrypoint: tc.entrypoint,
				// The command we will run for the test
				Command: []string{tc.command},
			}

			// Run the container, and get the output
			output := docker.Run(t, tag, opts)

			// The test check to assert we get what we expected.
			assert.Contains(t, output, tc.expected)
		})
	}
}
