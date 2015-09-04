package steps

import (
	"github.com/bitrise-io/bitrise-yml-converter/utils"
	bitriseModels "github.com/bitrise-io/bitrise/models"
	stepmanModels "github.com/bitrise-io/stepman/models"
)

const (
	// OldGitCloneFlavorBitriseStepID ...
	OldGitCloneFlavorBitriseStepID = "git-clone_flavor_bitrise"
	// NewGitCloneExtendedStepID ...
	NewGitCloneExtendedStepID = "git-clone-extended"
)

//----------------------
// old name: git-clone_flavor_bitrise
// new name: git-clone-extended

/*
old version source: https://github.com/bitrise-io/steps-git-clone.git

inputs:
  - GIT_REPOSITORY_URL
  - BITRISE_GIT_COMMIT
  - BITRISE_GIT_TAG
  - BITRISE_GIT_BRANCH
  - BITRISE_PULL_REQUEST
  - BITRISE_SOURCE_DIR
  - AUTH_USER
  - AUTH_PASSWORD
  - AUTH_SSH_PRIVATE_KEY
	- BITRISE_STEP_FORMATTED_OUTPUT_FILE_PATH
  - GIT_CLONE_IS_EXPORT_OUTPUTS
*/

/*
new version source: https://github.com/bitrise-io/steps-git-clone.git

inputs:
- repository_url
- commit
- tag
- branch
- pull_request_id
- clone_into_dir
- auth_user
- auth_password
- auth_ssh_private_key
- is_expose_outputs
*/

// ConvertGitCloneFlavorBitrise ...
func ConvertGitCloneFlavorBitrise(convertedWorkflowStep stepmanModels.StepModel) ([]bitriseModels.StepListItemModel, error) {
	newStepID := NewGitCloneExtendedStepID
	inputConversionMap := map[string]string{
		"repository_url":       "GIT_REPOSITORY_URL",
		"commit":               "BITRISE_GIT_COMMIT",
		"tag":                  "BITRISE_GIT_TAG",
		"branch":               "BITRISE_GIT_BRANCH",
		"pull_request_id":      "BITRISE_PULL_REQUEST",
		"clone_into_dir":       "BITRISE_SOURCE_DIR",
		"auth_user":            "AUTH_USER",
		"auth_password":        "AUTH_PASSWORD",
		"auth_ssh_private_key": "AUTH_SSH_PRIVATE_KEY",
		"is_expose_outputs":    "GIT_CLONE_IS_EXPORT_OUTPUTS",
	}

	return utils.ConvertStepAndCreateStepListItem(convertedWorkflowStep, newStepID, inputConversionMap)
}
