package stepper_test

import (
	"context"
	"github.com/ghodss/yaml"
	"github.com/google/go-cmp/cmp"
	"github.com/pkg/errors"
	"github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1"
	"github.com/tektoncd/pipeline/pkg/remote"
	"github.com/tektoncd/pipeline/pkg/remote/fake"
	"github.com/tektoncd/pipeline/pkg/stepper"
	"github.com/tektoncd/pipeline/test/diff"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

var (
	// generateTestOutput enable to regenerate the expected output
	generateTestOutput = os.Getenv("REGENERATE_TEST_OUTPUT") == "true"
)

func TestStepper(t *testing.T) {
	sourceDir := filepath.Join("test_data", "tests")
	fs, err := ioutil.ReadDir(sourceDir)
	if err != nil {
		t.Errorf(errors.Wrapf(err, "failed to read source dir %s", sourceDir).Error())
	}

	fakeResolver := fake.NewFileResolver(filepath.Join("test_data", "git"))

	remoteResolver := func(ctx context.Context, uses *v1beta1.Uses) (remote.Resolver, error) {
		return fakeResolver, nil
	}

	// make it easy to run a specific test only
	runTestName := os.Getenv("TEST_NAME")
	for _, f := range fs {
		if !f.IsDir() {
			continue
		}
		name := f.Name()
		if strings.HasPrefix(name, ".") {
			continue
		}
		if runTestName != "" && runTestName != name {
			t.Logf("ignoring test %s\n", name)
			continue
		}

		dir := filepath.Join(sourceDir, name)
		path := filepath.Join(dir, "input.yaml")
		expectedPath := filepath.Join(dir, "expected.yaml")
		data, err := ioutil.ReadFile(path)
		if err != nil {
			t.Errorf(errors.Wrapf(err, "failed to read file %s", path).Error())
		}

		prs := &v1beta1.PipelineRun{}
		err = yaml.Unmarshal(data, prs)
		if err != nil {
			t.Errorf(errors.Wrapf(err, "failed to unmarshal file %s", path).Error())
		}

		ctx := context.TODO()
		s := &stepper.Resolver{ResolveRemote: remoteResolver}
		err = s.Resolve(ctx, prs)
		if err != nil {
			t.Errorf(errors.Wrapf(err, "failed to invoke stepper on file %s", path).Error())
		}

		data, err = yaml.Marshal(prs)
		if err != nil {
			t.Errorf(errors.Wrapf(err, "failed to marshal output of stepper on file %s", path).Error())
		}

		if generateTestOutput {
			err = ioutil.WriteFile(expectedPath, data, 0666)
			if err != nil {
				t.Errorf(errors.Wrapf(err, "failed to save file %s", expectedPath).Error())
			}
			continue
		}
		expectedData, err := ioutil.ReadFile(expectedPath)
		if err != nil {
			t.Errorf(errors.Wrapf(err, "failed to load file %s", expectedPath).Error())
		}

		got := strings.TrimSpace(string(data))
		want := strings.TrimSpace(string(expectedData))

		if d := cmp.Diff(want, got); d != "" {
			t.Errorf("path %s diff %s", path, diff.PrintWantGot(d))
			t.Errorf("actual content for %s was: %s", path, got)
		}
	}
}
