// Test package
package artifact_registry_test

import (
	pb "artifact-registry/protos"
	registry "artifact-registry/registry"
	"fmt"
)

// Usage guide to get artifacts by IDs
func ExampleMLArtifactStore_GetArtifactsByID() {
	artifactStore := registry.ArtifactStore("run-uuid")

	artifact := &pb.MLArtifact{
		// ArtifactType: pb.MLArtifact_MODEL,
		Ids: []int64{5312},
	}
	response, _ := artifactStore.GetArtifactsByID(artifact)

	for _, artifactData := range response.Artifacts {
		fmt.Println(artifactData.GetName())
		fmt.Println(artifactData.GetUri())
		fmt.Println(artifactData.GetVersion())
		fmt.Println(artifactData.GetExecutionId())
		fmt.Println(artifactData.ArtifactType.Enum())
	}
	// Output:
	// MNIST
	// gcs://my-bucket/mnist
	// model_version_69389a49-b841-41a3-b1b2-15b3cb8c629e
	// run-2021-03-30T16:50:45.608098
	// MODEL
}

// Example usage to find Workspace by workspace name
func ExampleMLArtifactStore_GetWorkspace() {
	artifactStore := registry.ArtifactStore("run-uuid")

	workspace := &pb.Workspace{
		Name: "workspace_1",
	}

	response, _ := artifactStore.GetWorkspace(workspace)

	fmt.Println(response.Name)
	// Output:
	// workspace_1
}

// Example to fetch artifacts in a workspace
func ExampleWorkspace_GetArtifactsByWorkspace() {
	artifactStore := registry.ArtifactStore("run-uuid")

	workspaceInfo := &pb.Workspace{
		Name: "workspace_1",
	}

	workspace, _ := artifactStore.GetWorkspace(workspaceInfo)
	artifactList, _ := workspace.GetArtifactsByWorkspace()

	for range artifactList.GetArtifacts() {
		// do something
	}
	// Output:
	//
}

// Example to fetch artifacts by type in a workspace
func ExampleWorkspace_GetArtifactsByTypeWorkspace() {
	artifactStore := registry.ArtifactStore("run-uuid")

	workspaceInfo := &pb.Workspace{
		Name: "workspace_1",
	}

	workspace, _ := artifactStore.GetWorkspace(workspaceInfo)

    artifactTypeRequest := &pb.ArtifactByTypeRequest{
		ArtifactType: pb.ArtifactByTypeRequest_MODEL,
    }
	artifactList, _ := workspace.GetArtifactsByTypeWorkspace(artifactTypeRequest)

	for _, artifactData := range artifactList.GetArtifacts() {
		fmt.Println(artifactData.GetName())
	}
	// Output:
	// MNIST
	// MNIST
	// MNIST
    // FunctionComponent
}
