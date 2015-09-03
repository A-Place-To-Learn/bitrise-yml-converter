package converter

import (
	bitriseModels "github.com/bitrise-io/bitrise/models"
	stepmanModels "github.com/bitrise-io/stepman/models"
)

//----------------------
// old name: xcode-builder_flavor_bitrise_create-archive
// new name: xcode-archive

/*
old version source: https://github.com/bitrise-io/steps-xcode-builder.git

inputs:
  - XCODE_BUILDER_PROJECT_ROOT_DIR_PATH
  - XCODE_BUILDER_PROJECT_PATH
  - XCODE_BUILDER_SCHEME
  - XCODE_BUILDER_ACTION
  - XCODE_BUILDER_CERTIFICATE_URL
  - XCODE_BUILDER_CERTIFICATE_PASSPHRASE
  - XCODE_BUILDER_PROVISION_URL
  - XCODE_BUILDER_BUILD_TOOL
  - XCODE_BUILDER_CERTIFICATES_DIR
# Archive specific inputs
  - XCODE_BUILDER_DEPLOY_DIR
*/

/*
new version source: https://github.com/bitrise-io/steps-xcode-archive.git

inputs:
- project_path
- scheme
- workdir
- output_dir
*/

func convertXcodeBuilderFlavorBitriseCreateArchive(convertedWorkflowStep stepmanModels.StepModel) ([]bitriseModels.StepListItemModel, error) {
	stepListItems, err := certificateStep()
	if err != nil {
		return []bitriseModels.StepListItemModel{}, err
	}

	newStepID := newXcodeArchiveStepID
	inputConversionMap := map[string]string{
		"project_path": "XCODE_BUILDER_PROJECT_PATH",
		"scheme":       "XCODE_BUILDER_SCHEME",
		"output_dir":   "XCODE_BUILDER_DEPLOY_DIR",
	}

	newStep, err := convertStep(convertedWorkflowStep, newStepID, inputConversionMap)
	if err != nil {
		return []bitriseModels.StepListItemModel{}, err
	}

	stepIDDataString := BitriseVerifiedStepLibGitURI + "::" + newStepID
	stepListItems = append(stepListItems, bitriseModels.StepListItemModel{stepIDDataString: newStep})

	return stepListItems, nil
}
