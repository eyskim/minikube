/*
Copyright 2020 The Kubernetes Authors All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"fmt"
	"os/exec"
	"path"

	"github.com/pkg/errors"
	"k8s.io/minikube/pkg/minikube/download"
)

func uploadTarball(tarballFilename, k8sVer string) error {
	// Upload tarball to GCS
	hostPath := path.Join("out/", tarballFilename)
	gcsDest := fmt.Sprintf("gs://%s/%s/%s", download.PreloadBucket, download.PreloadVersion, k8sVer)
	cmd := exec.Command("gsutil", "cp", hostPath, gcsDest)
	fmt.Printf("Running: %v\n", cmd.Args)
	if output, err := cmd.CombinedOutput(); err != nil {
		return errors.Wrapf(err, "uploading %s to GCS bucket: %v\n%s", hostPath, err, string(output))
	}
	return nil
}
